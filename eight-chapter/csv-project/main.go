package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/spkg/bom"
)

func main() {
	f, err := os.Open("country.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	r := csv.NewReader(bom.NewReader(f)) // BOMの回避

	// Shift-JIS を扱う
	// r := csv.NewReader(transform.NewReader(f, japanese.ShiftJIS.NewDecoder()))

	// 区切り文字をカンマ(,)から変えたい場合はCommaフィールドを書き換える。
	// r.Comma = '\t'

	r.Comment = '#' // # で始まる行をコメントとみなし、取り込みをスキップ

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(record)
	}

	records := [][]string{
		{"書籍名", "出版年", "ページ数"},
		{"Golang", "2016", "280"},
		{"Python", "2018", "256"},
		{"Node.js", "2018", "316"},
	}

	file, err := os.OpenFile("oreilly.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	w := csv.NewWriter(file)

	// 区切り文字
	w.Comma = '\t'
	defer w.Flush()

	for _, record := range records {
		if err := w.Write(record); err != nil {
			log.Fatal(err)
		}
	}

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}
