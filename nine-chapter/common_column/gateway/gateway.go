package gateway

import(
	"context"
	"fmt"
)

// FetchUserはuserIDからユーザの情報を取得する
func FetchUser(ctx context.Context, userID string) (*User, error) {
	row := db.Query.RowContext(ctx, `SELECT user_id, user_name FROM users WHERE user_id = $1;`,
					userID)
	user, err := scanUser(row)
	if err != nil {
		return nil, fmt.Errorf("scan user: %w", err)
	}
	return user, nil
}

func scanUser(row *sql.Row) (*User, error) {
	var u User
	err := row.Scan(&u.UserID, &u.UserName)
	if err != nil {
		return nil, fmt.Errorf("row scan: %w", err)
	}
	return &u, nil
}