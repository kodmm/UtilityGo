package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// recoverを使ってハンドラーで発生した panicから復帰
		defer func() {
			if err := recover(); err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func NewLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func (lrw *loggingResponseWriter) Write(b []byte) (int, error) {
	if lrw.statusCode >= 400 {
		log.Printf("Response Body: %s", b)
	}
	return lrw.ResponseWriter.Write(b)
}

func wrapHandlerWithLogging(wrappedHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		lrw := NewLoggingResponseWriter(w)
		wrappedHandler.ServeHTTP(lrw, req)

		statusCode := lrw.statusCode

		log.Printf("%d %s", statusCode, http.StatusText(statusCode))
	})
}

func MiddlewareLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("start %s\n", r.URL)
		next.ServeHTTP(w, r)
		log.Printf("end %s\n", r.URL)
	})
}

func Healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func NewMiddlewareTx(db *sql.DB) func(http.Handler) http.Handler {
	return func(wrappedHandler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tx, _ := db.Begin()
			lrw := NewLoggingResponseWriter(w)
			r = r.WithContext(context.WithValue(r.Context(), "tx", tx))

			wrappedHandler.ServeHTTP(lrw, r)

			statusCode := lrw.statusCode
			if 200 <= statusCode && statusCode < 400 {
				log.Println("transaction commited")
				tx.Commit()
			} else {
				log.Print("transaction rolling back due to status code: ", statusCode)
				tx.Rollback()
			}
		})
	}
}

func extractTx(r *http.Request) *sql.Tx {
	tx, ok := r.Context().Value("tx").(*sql.Tx)
	if !ok {
		panic("transaction middleware is not supported")
	}
	return tx
}

func main() {
	db := openDB() // *sql.DBを取得
	tx := NewMiddlewareTx(db)
	http.Handle("/comments", tx(Recovery(http.HandlerFunc(Comments))))
	http.Handle("/healthz", MiddlewareLogging(http.HandlerFunc(Healthz)))
	http.Handle("/health", Recovery(MiddlewareLogging(http.HandlerFunc(Healthz))))
	http.ListenAndServe(":8888", nil)

}

func Comments(w http.ResponseWriter, r *http.Request) {
	tx := extractTx(r)
	//DBアクセス処理
}
