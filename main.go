package main

import (
	"log"

	"github.com/Marlos-Rodriguez/Twitter-Clon-Back/db"
	"github.com/Marlos-Rodriguez/Twitter-Clon-Back/handlers"
)

func main() {
	//Check conection to database
	if !db.CheckConnection() {
		log.Fatal("Sin conexion a la DB")
		return
	}

	//Run the handlers
	handlers.Handlers()
}
