package main

import (
	"github.com/maurofuentes/go-skill-factory-final-project/db"
	"github.com/maurofuentes/go-skill-factory-final-project/handlers"
	"log"
)

func main() {
	if db.ConnectionOn() == false {
		log.Fatal("db connection failed")
		return
	}

	handlers.Handlers()
}
