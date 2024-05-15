package main

import (
	"database/sql"
	"log"

	"github.com/vladsendrix/go-movies/cli"
	"github.com/vladsendrix/go-movies/concurrency"
	"github.com/vladsendrix/go-movies/controller"
	"github.com/vladsendrix/go-movies/database"
	"github.com/vladsendrix/go-movies/gui"
	"github.com/vladsendrix/go-movies/repository"
)

func setupDatabase() *sql.DB {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func setupMovieController(db *sql.DB) *controller.MovieController {
	movieRepo := repository.NewMovieRepository(db)
	movieController := controller.NewMovieController(movieRepo)
	return movieController
}

func main() {
	db := setupDatabase()
	defer db.Close()

	concurrencyFlag := cli.ParseFlags()

	if concurrencyFlag {
		concurrency.TestConcurrency(db)
	}

	movieController := setupMovieController(db)
	gui.StartGUI(movieController)
}
