package tag

type Book struct {
	ISBN        string `fixed:"1,18"`
	PublishDate string `fixed:"18,25"`
	Price       string `fixed:"26,29"`
	PDF         string `fixed:"30,34,left"`
	EPUB        string `fixed:"35,39,left"`
	EbookPrice  string `fixed:"40,44"`
}

s := `978-4-87311-865-9201909174620true true 3696
978-4-87311-865-9201909174620falsefalse0000
978-4-87311-865-9201909170000true true 0000
`

for _, line := range strings.Split(s, "\n") {
	var b Book
	if err := fixedwidth.UnmarshalCSV([]byte(line), &b); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", b)
}