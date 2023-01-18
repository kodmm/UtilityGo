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
}