package main

import (
	"log"

	"github.com/vladsendrix/go-movies/concurrency"
	"github.com/vladsendrix/go-movies/controller"
	"github.com/vladsendrix/go-movies/database"
	"github.com/vladsendrix/go-movies/gui"
	"github.com/vladsendrix/go-movies/repository"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	concurrency.TestConcurrency(db)

	movieRepo := repository.NewMovieRepository(db)

	movieController := controller.NewMovieController(movieRepo)

	gui.StartGUI(movieController)

}
