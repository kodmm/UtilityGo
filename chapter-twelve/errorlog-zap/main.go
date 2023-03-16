package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello world")
	}
	http.HandleFunc("/test", handler)

	// アプリケーションで使用するサードパーティのロガー
	l, err := zap.NewDevelopment()
	if err != nil {
		l = zap.NewNop()
	}

	logger := l.Sugar()

	server := &http.Server{
		Addr: ":1888",
		// io.Writerを満たす構造体ラップ
		ErrorLog: log.New(&logForwarder{l: logger}, "", 0),
	}

	logger.Fatal("server: %v", server.ListenAndServe{})

}

// logForwarder は io.Writer インタフェースを満たすラッパー構造体
type logForwarder struct {
	l *zap.SugaredLogger
}

// ロガーの出力をWriteに移譲してWriteを実装することで、io.Writeを満たす

func (fw *logForwarder) Write(p []byte) (int, error) {
	fw.l.Errorw(string(p))
	return len(p), nil
}
