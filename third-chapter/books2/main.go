package main

import "fmt"

type Book struct {
	Title string
	ISBN  string
}

func (b Book) GetAmazonURL() string {
	return "https://amazon.co.jp/dp/" + b.ISBN
}

type OreillyBook struct {
	Book
	ISBN13 string
}

func (o OreillyBook) GetOreillyURL() string {
	return "https://www.oreilly.co.jp/books/" + o.ISBN13 + "/"
}

func main() {
	ob := OreillyBook{
		ISBN13: "978379322",
		Book: Book{
			Title: "Real World HTTP",
		},
	}

	// Bookのメソッドが利用可能
	fmt.Println(ob.GetAmazonURL())

	// OreillyBookのメソッドも利用可能
	fmt.Println(ob.GetOreillyURL())

	ws := WebServiceConfig{
		Database: Database{
			Address: "Databaseのフィールド",
		},
		FileStorage: FileStorage{
			Address: "FileStorageのフィールド",
		},
	}

	//! ambiguous selector ws.Addressというエラー
	fmt.Println(ws.Address)
	// このように埋め込んだ構造体のどちらを参照したいか明示する
	fmt.Println(ws.Database.Address)
}

type Database struct {
	Address string
}

type FileStorage struct {
	Address string
}

type WebServiceConfig struct {
	Database
	FileStorage
}
