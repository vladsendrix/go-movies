package main

import (
	"flag"
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

	movieRepo := repository.NewMovieRepository(db)

	movieController := controller.NewMovieController(movieRepo)

	guiFlag := flag.Bool("gui", false, "Enable GUI")
	concurrencyFlag := flag.Bool("concurrency", false, "Test concurrency")
	flag.Parse()

	if *concurrencyFlag {
		concurrency.TestConcurrency(db)
	}

	if *guiFlag {
		gui.StartGUI(movieController)
	}

}
