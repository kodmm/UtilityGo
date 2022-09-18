package main

import "fmt"

//go:generate stringer -type=CarType
type CarType int

const (
	Sedan CarType = iota + 1
	Hatchback
	MPV
	SUV
	CrossOver
	Coupe
	Convertible
)

//go:generate stringer -type=CarOption
type CarOption uint64

const (
	GPS CarOption = 1 << iota
	AWD
	SunRoof
	HeatedSeat
	DriveAssist
)

//go:generate enumer -type=SMAP -json
type SMAP int

const (
	NAKAI SMAP = iota + 1
	KIMURA
	INAGAKI
	KUSANAGI
	KATORI
)

func main() {
	c := AWD
	fmt.Printf("握力王の愛車は%sです\n", c)

	reader := NAKAI
	fmt.Printf("SMAPのリーダーは%sです\n", reader)
}
