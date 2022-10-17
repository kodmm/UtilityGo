package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
)

var ErrNotFound = errors.New("not found")

func findBook(isbn string) (*Book, error) {
	// ...
	// 値が取得できなかったため、ErrNotFoundを返す
	return nil, ErrNotFound
}

func validate(length int) error {
	if length <= 0 {
		return fmt.Errorf("length must be greater than0, length = %d", length)
	}
	// ...
	return nil
}

// HTTPError はステータスコードが200以外の場合のエラーを扱う構造体
type HTTPError struct {
	StatusCode int
	URL        string
}

// Errorメソッドを実装することでHTTPError 構造体のポインターは Errorインタフェースを満たしている
func (he *HTTPError) Error() string {
	return fmt.Sprintf("http status code = %d, url = %s", he.StatusCode, he.URL)
}

func ReadContents(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// レスポンスのステータスコードが200ではない場合は、*HTTPErrorとしてエラーを返却
	if resp.StatusCode != http.StatusOK {
		return nil, &HTTPError{
			StatusCode: resp.StatusCode,
			URL:        url,
		}
	}
	return io.ReadAll(resp.Body)
}

// loadConfigError は設定ファイル読み込みを表すエラー
type loadConfigError struct {
	msg string
	err error
}

func (e *loadConfigError) Error() string {
	return fmt.Sprintf("cannot load config: %s (%s)", e.msg, e.err.Error())
}

func (e *loadConfigError) Unwrap() error {
	return e.err
}

type Config struct {
	// ...
}

func LoadConfig(configFilePath string) (*Config, error) {
	var cfg *Config
	data, err := os.ReadFile(configFilePath)
	if err != nil {
		return nil, &loadConfigError{msg: fmt.Sprintf("read file `%s`", configFilePath), err: err}
	}
	if err = json.Unmarshal(data, cfg); err != nil {
		return nil, &loadConfigError{msg: fmt.Sprintf("parse config file `%s`", configFilePath), err: err}
	}
	return cfg, nil
}

func fetchSession(ctx context.Context, sessionID string) (*Session, error) {
	var s *Session

	query := `SELECT * FROM sessions WHERE session_id = $1;`
	err := db.QueryRowContext(ctx, query, sessionID).Scan(&s)
	if err != nil {
		return nil, fmt.Errorf("fetch session, sessionID = %s: %w", sessionID, err)
	}
	return s, nil
}

func main() {
	isNew := false
	u, err := fetchSession(ctx, sessionID)
	if err != nil {
		// DBレイヤーから返却されたエラーの詳細を調べてハンドリングできる
		// 比較できないエラーや==で合致しなかった場合にエラーがIs()メソッドを実装している場合は、
		// 実装されているIs()メソッドを使って評価し、合致した場合は結果を返す
		// Is()を実装していない場合やIs()で合致しなかった場合はエラーをUnwrap()し、次のアンラップされたエラーと評価する。
		// 最終的にUnwrap()して得られるエラーがなくなった場合は合致しなかった結果を返す
		if !errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("fail to fetch session: %w", err)
		}
	}

	if _, err := LoadConfig("non-existing"); err != nil {
		// errors.As の中でエラーをアンラップして詳細なエラーを調べられる
		var pathError *fs.PathError
		if errors.As(err, &pathError) {
			// ...
		} else {
			// ...
		}
	}

}
