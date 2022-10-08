package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type Handler struct {
	db *sql.DB
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	var v Value
	err := h.db.Get(&v, "SELECT * FROM person WHERE id=$1", r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(v)
}

func register(h *Handler) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/value", h.Get)
	return mux
}

func register(db *sql.DB) http.Handler {
	mux := http.NewServeMux()

	// 関数定義は関数ない関数で実装
	// 以前レシーバ経由で使っていたDB Connectionは
	// クロージャがキャプチャした関数定義外のregister()の引数を利用
	mux.HandleFunc("/value", func(w http.ResponseWriter, r *http.Request) {
		var v Value
		err := db.Get(&v, "SELECT * FROM person WHERE id=$1", r.URL.Query().Get("id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(v)
	})
	return mux
}

func main() {

}
