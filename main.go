package main

import (
	"context"
	"log"
	"sync"

	"github.com/vladsendrix/go-movies/database"
	"github.com/vladsendrix/go-movies/repository"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	directorRepo := repository.NewDirectorRepository(db)

	director := repository.Director{
		Name: "Christopher Nolan",
	}

	_, err = directorRepo.Create(context.Background(), director)
	if err != nil {
		log.Fatal(err)
	}

	movieRepo := repository.NewMovieRepository(db)

	err = movieRepo.DeleteAll(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	movie := repository.Movie{
		Title:       "Inception",
		ReleaseYear: 2010,
		DirectorID:  1,
	}

	_, err = movieRepo.Create(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Director and movie added successfully")

	movies, err := movieRepo.GetAll(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	for _, m := range movies {
		movie := m.(repository.Movie)
		log.Printf("Movie: %s, Release Year: %d, Director ID: %d\n", movie.Title, movie.ReleaseYear, movie.DirectorID)
	}

	var wg sync.WaitGroup
	wg.Add(2)

	// Transaction 1: Update a movie's release year
	go func() {
		defer wg.Done()

		tx, err := db.Begin()
		if err != nil {
			log.Fatal(err)
		}

		_, err = tx.Exec("UPDATE movies SET release_year = 2022 WHERE director_id = 1")
		if err != nil {
			log.Fatal(err)
		}

		// Don't commit the transaction yet
	}()

	// Transaction 2: Read the movie's release year
	go func() {
		defer wg.Done()

		tx, err := db.Begin()
		if err != nil {
			log.Fatal(err)
		}

		var releaseYear int
		err = tx.QueryRow("SELECT release_year FROM movies WHERE director_id = 1").Scan(&releaseYear)
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Release Year:", releaseYear)

		if err := tx.Commit(); err != nil {
			log.Fatal(err)
		}
	}()

	wg.Wait()
}
