package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/gocarina/gocsv"
	"github.com/spkg/bom"
)

type Country struct {
	Name       string `csv:"国名"`
	ISOCode    string `csv:"ISOコード"`
	Population int    `csv:"人工"`
}

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

	// encoding/csvの設定をgocarina/gocsvに引き継ぐ
	fn := func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(bom.NewReader(in))
		r.Comma = '\t'
		r.Comment = '#'
		return r
	}
	gocsv.SetCSVReader(fn)

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

	lines := []Country{
		{Name: "アメリカ合衆国", ISOCode: "US/USA", Population: 310232863},
		{Name: "日本", ISOCode: "JP/JPN", Population: 126288000},
		{Name: "中国", ISOCode: "CN/CHN", Population: 1330044000},
	}

	fileA, err := os.Create("country2.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer fileA.Close()

	// CSVヘッダーが不要な場合は、gocsv.MarshalWithoutHeaders(&lines, fileA)を使用する。
	if err := gocsv.MarshalFile(&lines, fileA); err != nil {
		log.Fatal(err)
	}

	fileB, err := os.Open("country2.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var linesB []Country
	if err := gocsv.UnmarshalFile(fileB, &linesB); err != nil {
		log.Fatal(err)
	}

	for _, v := range linesB {
		fmt.Printf("%#v\n", v)
	}

}
