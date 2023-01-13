package main

type Book struct {
	ISBN string
	PublishDate string
	Price string
	PDF string
	EPUB string
	EbookPrice string
}

s := `978-4-87311-865-9201909174620true true 3696
978-4-87311-865-9201909174620falsefalse0000
978-4-87311-865-9201909170000true true 0000
`

for _, line := range strings.Split(s, "\n") {
	r := []rune(line)

	res := Book{
		ISBN: string(r[0:17]),
		PublishDate: string(r[17:25]),
		Price: string(r[25:29]),
		PDF: string(r[29:34]),
		EPUB: string(r[34:39]),
		EbookPrice: string(r[39:43]),
	}

	fmt.Printf("%+v\n", res)

}