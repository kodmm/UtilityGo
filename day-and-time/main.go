package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	tz, _ := time.LoadLocation("America/Los_Angeles")
	future := time.Date(2015, time.October, 21, 7, 28, 0, 0, tz)

	fmt.Println(now.String())
	fmt.Println(future.Format(time.RFC3339Nano))

	now2 := time.Date(1995, time.October, 26, 9, 0, 0, 0, time.Local)
	past := time.Date(1995, time.November, 12, 6, 38, 0, 0, time.UTC)

	fmt.Println(now2)
	fmt.Println(past)

	fiveMinute := 5 * time.Minute
	fmt.Println(now.Add(fiveMinute))
	fmt.Println("*************")
	// intとは型違いで直接演算できないので、即値との計算以外は
	// time.Durationへの明示的なキャストが必要
	// キャストがないとエラーが発生する。
	// invalid operation: seconds * time.Second(mismatched types int and time.Duration)
	const seconds int = 10
	timeSeconds := time.Duration(seconds) * time.Second

	// Timeの演算でduration作成
	dur := now.Sub(past)

	fmt.Println((timeSeconds))
	fmt.Println(dur)

	// 1時間にまとめてパッチで読み込むファイル名を取得
	filepath := time.Now().Truncate(time.Minute).Format("20060102150405.json")

	// 5分ごと5分前の時刻
	fiveMinuteAfter := time.Now().Add(fiveMinute)
	fiveMinuteBefore := time.Now().Add(-fiveMinute)

	fmt.Println("filepath", filepath)
	fmt.Println("fiveMinuteBefore", fiveMinuteBefore)
	fmt.Println("fiveMinuteAfter", fiveMinuteAfter)
}
