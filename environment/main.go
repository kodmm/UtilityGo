package main

import (
	"log"
	"os"
	"strings"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port      uint16 `envconfig:"PORT" default:"3000"`
	Host      string `envconfig:"HOST" required:"true"`
	AdminPort uint16 `envconfig:"ADMIN_PORT" default:"3001"`
}

func main() {
	var c Config
	err := envconfig.Process("", &c)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	log.Println(c)

	// 文字列結合
	src := []string{"A", "B", "C", "D"}
	var builder strings.Builder
	builder.Grow(4)
	for i, word := range src {
		if i != 0 {
			builder.WriteByte(' ')
		}
		builder.WriteString(word)
	}
	log.Println(builder.String())

}
