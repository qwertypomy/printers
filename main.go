package main

import (
	"github.com/qwertypomy/printers/config"
	"github.com/qwertypomy/printers/db"
	"github.com/qwertypomy/printers/models"
	"log"
)

var Config models.Config

func main() {
	Config, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
		return
	}

	db.Populate(Config)
}
