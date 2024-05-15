package repository

import (
	"database/sql"
	"fmt"

	"github.com/vladsendrix/go-movies/entities"
)

type Director = entities.Director

type DirectorRepository struct {
	db *sql.DB
}

func NewDirectorRepository(db *sql.DB) *DirectorRepository {
	return &DirectorRepository{db: db}
}

func (r *DirectorRepository) GetByID(id interface{}) (interface{}, error) {
	var director Director
	err := r.db.QueryRow("SELECT * FROM directors WHERE id = ?", id).Scan(&director.Name)
	if err != nil {
		return nil, fmt.Errorf("could not get director: %v", err)
	}
	return director, nil
}

func (r *DirectorRepository) GetAll() ([]interface{}, error) {
	rows, err := r.db.Query("SELECT * FROM directors")
	if err != nil {
		return nil, fmt.Errorf("could not get directors: %v", err)
	}
	defer rows.Close()

	var directors []interface{}
	for rows.Next() {
		var director Director
		if err := rows.Scan(&director.ID, &director.Name); err != nil {
			return nil, fmt.Errorf("could not scan director: %v", err)
		}
		directors = append(directors, director)
	}
	return directors, nil
}

func (r *DirectorRepository) Create(obj interface{}) error {
	director, ok := obj.(Director)
	if !ok {
		return fmt.Errorf("invalid type, expected Director")
	}
	_, err := r.db.Exec("INSERT INTO directors (name) VALUES ($1)", director.Name)
	if err != nil {
		return fmt.Errorf("could not insert director: %v", err)
	}
	return nil
}

func (r *DirectorRepository) Update(id interface{}, obj interface{}) error {
	director, ok := obj.(Director)
	if !ok {
		return fmt.Errorf("invalid type, expected Director")
	}
	_, err := r.db.Exec("UPDATE directors SET name = $1 WHERE id = $2", director.Name, id)
	if err != nil {
		return fmt.Errorf("could not update director: %v", err)
	}
	return nil
}

func (r *DirectorRepository) Delete(id interface{}) error {
	_, err := r.db.Exec("DELETE FROM directors WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("could not delete director: %v", err)
	}
	return nil
}

func (r *DirectorRepository) DeleteAll() error {
	_, err := r.db.Exec("DELETE FROM directors")
	if err != nil {
		return fmt.Errorf("could not delete all directors: %v", err)
	}
	return nil
}
