package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

func fetchCapacity(ctx context.Context, key string) (int, error) {
	var capacity int
	query := `SELECT value FROM parameter_master WHERE key = $1;`
	err := db.QueryRowContext(ctx, query, key).Scan(&capacity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// レコードが存在しなかった場合はデフォルト値を設定し、処理を継続。ログも出力する。
			log.Printf("fetch capacity not found, using default capacity, key = %s", key)
			return 10, nil
		}
		return -1, err
	}
	return capacity, nil
}

func main() {
	user, err := getInvitedUserWithEmail(ctx, email)
	if err != nil {
		// 呼び出し先で発生したエラーをラップし、付与情報を付与して呼び出し元に返却
		//TODO: どのような処理で、どのような引数をもとに、どんなエラーが出たのかを書く
		fmt.Errorf("fail to get invited user with email(%s): %w", email, err)
	}
}
