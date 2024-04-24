package main

import (
	"log"

	"github.com/vladsendrix/go-movies/db"
)

func main() {
	db, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

}
