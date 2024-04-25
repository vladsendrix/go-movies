package main

import (
	"context"
	"log"

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
}
