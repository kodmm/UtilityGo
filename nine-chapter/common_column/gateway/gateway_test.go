package gateway

import (
	"database/sql"
	"io/ioutil"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestFetchUser(t *testing.T) {
	connStr := "host=localhost port=5432 user=testuser dbname=testdb password=pass sslmode=disable"
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	sqlBytes, err := ioutil.ReadFile("./schema.sql")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := db.ExecContext(context.TODO(), string(sqlBytes)); err != nil {
		t.Fatal(err)
	}
	tests := []struct {
		name string
		userID string
		inputTestSQL string
		want *User
		hasErr bool
	}{
		{
			name: "1件取得",
			userID: "001",
			inputTestSQL: "./testdata/input_user_1.sql",
			want: &User{UserID: "001", UserName: "gopher1"},
			hasErr: false, 
		},
		{
			name: "0件取得",
			userID: "9999",
			inputTestSQL: "./testdata/input_user_2.sql",
			want: nil,
			hasErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sqlBytes, err := ioutil.ReadFile(tt.inputTestSQL)
			if err != nil {
				t.Fatal(err)
			}
			if _, err := db.ExecContext(context.TODO(), string(sqlBytes)); err != nil {
				t.Fatal(err)
			}
			t.Cleanup(func() {
				if _, err := db.ExecContext(context.TODO(), `TRUNCATE users;`); err != nil {
					t.Fatal(err)
				}
			})
		})
		got, err := FetchUser(context.TODO(), tt.userID)
		if (err != nil) != tt.hasErr {
			t.Fatalf("FetchUser() error = %v, hasError %v", err, tt.hasErr)
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("FetchUser() got = %v, want %v", got, tt.want)
		}
	}

}

func TestMockFetchUser(t *testing.T) {
	type mock struct {
		db *sql.DB
		sqlmock sqlmock.Sqlmock
	}
	tests := []struct {
		name string
		userID string
		mock mock
		want *User
		hasErr bool
	}{
		{
			name: "1件取得",
			userID: "0001",
			// mockの実装
			mock: func() mock {
				db, m, err := sqlmock.New()
				if err != nil {
					t.Fatal(err)
				}
				m.ExpectQuery(regexp.QuoteMeta(`SELECT user_id, user_name FROM users WHERE user_id $1;`)).
					WithArgs("0001").
					WillReturnRows(sqlmock.NewRows([]string{"user_id", "user_name"}).
					AddRow("0001", "gopher1"),
					)
				return mock{db, m}
			}(),
			want: &User{UserID: "0001", UserName: "gopher1"},
			hasErr: false,
		},
		{
			name: "0件取得",
			userID: "9999",
			mock: func() mock {
				db, m, err := sqlmock.New()
				if err != nil {
					t.Fatal(err)
				}
				m.ExpectQuery(regexp.QuoteMeta(`SELECT user_id, user_name FROM users WHERE user_id = $1;`)).
				WithArgs("9999").
				WillReturnError(sql.ErrNoRows)
				return mock{db, m}
			}(),
			want: nil,
			hasErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db = tt.mock.db
			got, err := FetchUser(context.Backgrouund(), tt.userID)
			if (err != nil) != tt.hasErr {
				t.Fatalf("FetchUser() error = %v, hasErr %v", err, tt.hasErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FetchUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}