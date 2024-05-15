package repository

import (
	"database/sql"
	"fmt"

	"github.com/vladsendrix/go-movies/entities"
)

type Movie = entities.Movie

type MovieRepository struct {
	db *sql.DB
}

func NewMovieRepository(db *sql.DB) *MovieRepository {
	return &MovieRepository{db: db}
}

func (r *MovieRepository) GetByID(id interface{}) (interface{}, error) {
	var movie Movie
	err := r.db.QueryRow("SELECT * FROM movies WHERE id = ?", id).Scan(&movie.ID, &movie.Title, &movie.ReleaseYear, &movie.DirectorID)
	if err != nil {
		return nil, fmt.Errorf("could not get movie: %v", err)
	}
	return movie, nil
}

func (r *MovieRepository) GetAll() ([]interface{}, error) {
	rows, err := r.db.Query("SELECT * FROM movies")
	if err != nil {
		return nil, fmt.Errorf("could not get movies: %v", err)
	}
	defer rows.Close()

	var movies []interface{}
	for rows.Next() {
		var movie Movie
		if err := rows.Scan(&movie.ID, &movie.Title, &movie.ReleaseYear, &movie.DirectorID); err != nil {
			return nil, fmt.Errorf("could not scan movie: %v", err)
		}
		movies = append(movies, movie)
	}
	return movies, nil
}

func (r *MovieRepository) Create(obj interface{}) error {
	movie, ok := obj.(Movie)
	if !ok {
		return fmt.Errorf("invalid type, expected Movie")
	}
	_, err := r.db.Exec("INSERT INTO movies (title, release_year, director_id) VALUES ($1, $2, $3)", movie.Title, movie.ReleaseYear, movie.DirectorID)
	if err != nil {
		return fmt.Errorf("could not insert movie: %v", err)
	}
	return nil
}

func (r *MovieRepository) Update(id interface{}, obj interface{}) error {
	movie, ok := obj.(Movie)
	if !ok {
		return fmt.Errorf("invalid type, expected Movie")
	}
	_, err := r.db.Exec("UPDATE movies SET title = $1, release_year = $2, director_id = $3 WHERE id = $4", movie.Title, movie.ReleaseYear, movie.DirectorID, id)
	if err != nil {
		return fmt.Errorf("could not update movie: %v", err)
	}
	return nil
}

func (r *MovieRepository) Delete(id interface{}) error {
	_, err := r.db.Exec("DELETE FROM movies WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("could not delete movie: %v", err)
	}
	return nil
}

func (r *MovieRepository) DeleteAll() error {
	_, err := r.db.Exec("DELETE FROM movies")
	if err != nil {
		return fmt.Errorf("could not delete all movies: %v", err)
	}
	return nil
}
