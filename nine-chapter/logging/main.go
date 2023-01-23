package main

import(
	"context"
	"log"

	_ "github.com/jackc/pgx/v4/stdlib"
)

var _ pgx.Logger = (*logger)(nil)

type logger struct {}

func (l *logger) Log(ctx context.Context, level pgx.LogLevel, msg string, data map[string]interface{}) {
	if msg == "Query" {
		log.Printf("SQL:\n%v\nARGS:%v\n", data["sql"], data["args"])
	}
}

func main() {
	ctx := context.Background()

	config, err := pgx.ParseConfig("user=testuser password=pass host=localhost port=5432 dbname=testdb sslmode=disable")
	if err != nil {
		log.Fatalf("parse config: %v\n", err)
	}
	config.Logger = &logger{}

	conn, err := pgx.ConnectConfig(ctx, config)
	if err != nil {
		log.Fatalf("connect: %v\n", err)
	}

	sql := `SELECT schemaname, tablename FROM pg_tables WHERE schemaname = $1;`
	args := `information_schema`

	rows, err := conn.Query(ctx, sql, args)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var pgtables []PgTable
	for rows.Next() {
		var s, t string
		if err := rows.Scan(&s, &t); err != nil {
			//...
		}
		pgtables = append(pgtables, Pgtable{SchemaName: s, TableName: t})
	}
	if err := rows.Error(); err != nil {
		log.Fatal(err)
	}

	users := []User{
		{"0001", "Gopher"},
		{"0002", "Ferris"},
		{"0003", "Duke"},
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	// プリペアードステートメントの構築
	stmt, err := tx.PrepareContext(ctx, "INSERT INTO users(user_id, user_name, created_at)VALUES($1, $2, $3);")
	if err != nil {
		// error handling
	}
	defer stmt.Close()

	for _, u := range users {
		// 構築したプリペアードステートメントにパラメータをセットし、実行
		if _, err := stmt.ExecContext(ctx, u.UserID, u.UserName); err != nil {
			// error handling
		}
	}
	if err := tx.Commit(); err != nil {
		// errorhandling
	}
}