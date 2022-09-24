package main

import (
	"log"
	"os"

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
}
