package main

import (
	"log"

	"github.com/vladsendrix/go-movies/database"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

}
