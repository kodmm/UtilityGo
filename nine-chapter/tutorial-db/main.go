package main

import(
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type User struct {
	UserID string
	UserName string
	CreatedAt time.Time
}

func main() {
	db, err := sql.Open("pgx", "host=localhost port=5432 user=testuser dbname=testdb password=pass sslmode=disable")
	if err != nil {
		// err handling
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.QueryContext(ctx, `SELECT user_id, user_name, created_at FROM users ORDER BY user_id;`)
	if err != nil {
		log.Fatalf("query all users: %v", err)
	}
	defer rows.Close()

	var users []*User
	for rows.Next() {
		var (
			userID, userName string
			createdAt time.Time
		)
		if err := rows.Scan(&userID, &userName, &createdAt); err != nil {
			log.Fatalf("scan the user: %v", err)
		}
		users = append(users, &User{
			UserID: userID,
			UserName: userName,
			CreatedAt: createdAt,
		})
	}
	if err := rows.Close(); err != nil {
		log.Fatalf("rows close: %v", err)
	}

	if err := rows.Err(); err != nil {
		log.Fatalf("scan users: %v", err)
	}

	var(
		userName string
		createdAt time.Time
	)
	row := db.QueryRowContext(ctx, `SELECT user_name, created_at FROM users WHERE user_id = $1;`, userID)
	err = row.Scan(&userName, &createdAt)
	if err != nil {
		log.Fatalf("query row(user_id=%s): %v", userID, err)
	}
	u := User{
		UserID: userID,
		UserName: userName,
		CreatedAt: createdAt,
	}

	type Service struct {
		db *sql.DB
	}

	func (s *Service) UpdateProduct(ctx context.Context, productID, string) error {
		tx, err := s.db.Begin()
		if err != nil {
			return err
		}
		if _, err := tx.ExecContext(ctx, "UPDATE products SET price = 200 WHERE product_id = $1", productID); err != nil {
			tx.Rollback()
			return err
		}
		return tx.Commit()
	}

	func (s *Service) UpdateProducte2(ctx context.Context, productID string) (err error) {
		tx, err := s.db.Begin()
		if err != nil {
			return err
		}
		defer tx.Rollback()

		if _, err := tx.ExecContext(ctx, "UPDATE products SET price = 200 WHERE product_id = $1", productID); err != nil {
			return err
		}

		return tx.Commit()
	}

	/// txAdminはトランザクション制御するための構造体
	type txAdmin struct {
		*sql.DB
	}

	type ServiceAdmin struct {
		tx txAdmin
	}

	// Transactionはトランザクションを制御するメソッド
	// アプリケーション開発者が本メソッドを使って、DMLのクエリーを発行する
	func (t *txAdmin) Transaction(ctx context.Context, f func(ctx context.Context) (err error)) error {
		tx, err := t.Begin()
		if err != nil {
			return err
		}
		defer tx.Rollback()
		if err := f(ctx); err != nil {
			return fmt.Errorf("transaction query failed: %w", err)
		}
		return tx.Commit()
	}

	func (s *ServiceAdmin) UpdateProduct(ctx context.Context, productID string) error {
		updateFunc := func(ctx context.Context) error {
			if _, err := s.tx.ExecContext(ctx, "UPDATE products SET price = 200 WHERE product_id = $1", productID); err != nil {
				return err
			}
			if _, err := s.tx.ExecContext(ctx, "UPDATE products SET price = 200 WHERE product_id = $1", productID); err != nil {
				return err
			}
			return nil
		}
		return s.tx.Transaction(ctx, updateFunc)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	// クエリーはコンテキストによって5秒後にキャンセルされます
	if _, err := db.ExecContext(ctx, "SELECT pg_sleep(100)"); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			log.Println("canceling query")
		} else {
			// error handling
		}
	}
}