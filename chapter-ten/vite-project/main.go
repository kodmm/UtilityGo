package main

import (
	"context"
	"embed"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5"
)

//go:embed vite-project/dist/*
var assets embed.FS

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	// まずはリクエストされた通りにファイルを探索
	err := tryRead(r.URL.Path, w)
	if err == nil {
		return
	}
	// 見つからなければindex.htmlを返す
	err = tryRead("index.html", w)
	if err != nil {
		panic(err)
	}
}

func newHandler() http.Handler {
	router := chi.NewRouter()

	router.Route("/api", func(r chi.Router) {
		// 何かAPIを足したい場合はここに足す
		r.Get("/test", func(w http.ResponseWriter, r *http.Request) {})
	})
	// シングルアプリケーションを配布するハンドラー
	router.NotFound(notFoundHandler)

	return router
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	server := &http.Server{
		Addr:    ":8000",
		Handler: newHandler(),
	}

	go func() {
		<-ctx.Done()
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		server.Shutdown(ctx)
	}()
	fmt.Printf("start receiving at :8000")
	fmt.Fprintln(os.Stderr, server.ListenAndServe())

}
