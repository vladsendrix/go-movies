package repository

import (
	"context"
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

func (r *DirectorRepository) GetByID(ctx context.Context, id interface{}) (interface{}, error) {
	var director Director
	err := r.db.QueryRowContext(ctx, "SELECT * FROM directors WHERE id = $1", id).Scan(&director.ID, &director.Name, &director.BirthYear)
	if err != nil {
		return nil, fmt.Errorf("could not get director: %v", err)
	}
	return director, nil
}

func (r *DirectorRepository) GetAll(ctx context.Context) ([]interface{}, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT * FROM directors")
	if err != nil {
		return nil, fmt.Errorf("could not get directors: %v", err)
	}
	defer rows.Close()

	var directors []interface{}
	for rows.Next() {
		var director Director
		if err := rows.Scan(&director.ID, &director.Name, &director.BirthYear); err != nil {
			return nil, fmt.Errorf("could not scan director: %v", err)
		}
		directors = append(directors, director)
	}
	return directors, nil
}

func (r *DirectorRepository) Create(ctx context.Context, obj interface{}) (interface{}, error) {
	director, ok := obj.(Director)
	if !ok {
		return nil, fmt.Errorf("invalid type, expected Director")
	}
	_, err := r.db.ExecContext(ctx, "INSERT INTO directors (name, birth_year) VALUES ($1, $2)", director.Name, director.BirthYear)
	if err != nil {
		return nil, fmt.Errorf("could not insert director: %v", err)
	}
	return director, nil
}

func (r *DirectorRepository) Update(ctx context.Context, id interface{}, obj interface{}) error {
	director, ok := obj.(Director)
	if !ok {
		return fmt.Errorf("invalid type, expected Director")
	}
	_, err := r.db.ExecContext(ctx, "UPDATE directors SET name = $1, birth_year = $2 WHERE id = $3", director.Name, director.BirthYear, id)
	if err != nil {
		return fmt.Errorf("could not update director: %v", err)
	}
	return nil
}

func (r *DirectorRepository) Delete(ctx context.Context, id interface{}) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM directors WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("could not delete director: %v", err)
	}
	return nil
}
