package main

import (
	"io"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func initServer() http.Handler {
	r := chi.NewRouter()
	r.Get("/fortune", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		io.WriteString(w, `{"fortune": "大吉"}`)
	})
	return r
}

func main() {
	log.Println("open at 8080")
	log.Println(http.ListenAndServe(":8080", initServer()))
}
