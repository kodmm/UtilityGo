package sample

import (
	"fmt"
	"io"
)

type User struct {
	Name    string
	Address string
}

func DumpUserTo(w io.Writer, u *User) {
	if u.Address == "" {
		fmt.Fprintf(w, "%s(住所不定)", u.Name)
	} else {
		fmt.Fprintf(w, "%s@%s", u.Name, u.Address)
	}
}
