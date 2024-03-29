package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/vektah/gqlparser/v2/validator"
)

type Comment struct {
	Message  string `validate:"required,min=1,max=140"`
	UserName string `validate:"required,min=1,max=15"`
}

type Book struct {
	Title string `validate:"required"`
	Price *int   `validate:"required"`
}

func main() {
	var mutex = &sync.RWMutex{}
	comments := make([]Comment, 0, 100)

	http.HandleFunc("/comments", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		switch r.Method {
		case http.MethodGet:
			mutex.RLock() // 読み取り時に書き込みがあることを考慮しロックする。本来はDBから読みとる処理を代替

			if err := json.NewEncoder(w).Encode(comments); err != nil {
				http.Error(w, fmt.Sprintf(`{"status": "%s"}`, err), http.StatusInternalServerError)
				return
			}
			mutex.RUnlock()
		case http.MethodPost: //POSTメソッドの処理
			var c Comment
			if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
				http.Error(w, fmt.Sprintf(`{"status": "%s"}`, err), http.StatusInternalServerError)
				return
			}

			validate := validator.New()
			if err := validate.Struct(c); err != nil {
				var out []string
				var ve validator.ValidationErrors
				if errors.As(err, &ve) {
					for _, fe := range ve {
						switch fe.Field() {
						case "Message":
							out = append(out, fmt.Sprintf("Messageは1 ~ 140文字です"))
						case "UserName":
							out = append(out, fmt.Sprintf("Messageは1 ~ 15文字です"))
						}
					}
				}
				http.Error(w, fmt.Sprintf(`{"status": "%s"}`, strings.Join(out, ",")), http.StatusBadRequest)
				return
			}

			mutex.Lock() //同時にアクセスを防ぐためのロック。本来はDBに保存する処理を代替
			comments = append(comments, c)
			mutex.Unlock()

			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(`{"status": "created"}`))

		default:
			http.Error(w, `{"status": "permit only GET or POST"}`, http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/params", func(w http.ResponseWriter, r *http.Request) {
		// FormValue()を呼ぶ場合は
		// ParseForm()メソッドの呼び出しは省略可
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		word := r.FormValue("searchword")
		log.Printf("searchword = %s\n", word)

		words, ok := r.Form["searchword"]
		log.Printf("search words = %v has values %v", words, ok)

		log.Printf("all queries")
		for key, values := range r.Form {
			log.Printf(" %s : %v\n", key, values)
		}
	})

	http.HandleFunc("/file", func(w http.ResponseWriter, r *http.Request) {
		// ParseMultipartForm()メソッドの
		// 呼び出しは省略可だが、省略時は32MBになる
		err := r.ParseMultipartForm(32 * 1024 * 1024)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// ファイルを取り出してストレージに取り出す
		f, h, err := r.FormFile("file")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		log.Println(h.Filename)
		o, err := os.Create(h.Filename)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer o.Close()
		_, err = io.Copy(o, f)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// ファイルと一緒に送信されたデータの取得
		value := r.PostFormValue("data")
		log.Printf(" value = %s\n", value)

	})

	http.ListenAndServe(":8888", nil)

	s := `{"Title": "Real World HTTP", "Price": 0}`
	var b Book
	if err := json.Unmarshal([]byte(s), &b); err != nil {
		log.Fatal(err)
	}

	if err := validator.New().Struct(b); err != nil {
		var ve validator.ValidationErrors // validatorの独自型に変換
		if errors.As(err, &ve) {
			for _, fe := range ve {
				fmt.Printf("フィールド %s が %s 違反です(値: %v)\n", fe.Field(), fe.Tag(), fe.Value())
				// フィールドPriceがrequired違反です(値: 0)
			}
		}
	}
}
