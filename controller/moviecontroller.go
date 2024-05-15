package controller

import (
	"github.com/vladsendrix/go-movies/entities"
	"github.com/vladsendrix/go-movies/repository"
)

type MovieController struct {
	MovieRepo *repository.MovieRepository
}

func NewMovieController(movieRepo *repository.MovieRepository) *MovieController {
	return &MovieController{MovieRepo: movieRepo}
}

func (c *MovieController) GetByID(id int) (*entities.Movie, error) {
	movie, err := c.MovieRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return movie, nil
}

func (c *MovieController) GetAll() ([]*entities.Movie, error) {
	movies, err := c.MovieRepo.GetAll()
	if err != nil {
		return nil, err
	}

	movieEntities := make([]*entities.Movie, len(movies))
	for i, movie := range movies {
		movieEntities[i] = &movie
	}

	return movieEntities, nil
}

func (c *MovieController) Create(movie *entities.Movie) error {
	err := c.MovieRepo.Create(movie)
	if err != nil {
		return err
	}
	return nil
}

func (c *MovieController) Delete(id int) error {
	err := c.MovieRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (c *MovieController) Update(id int, movie *entities.Movie) error {
	err := c.MovieRepo.Update(id, movie)
	if err != nil {
		return err
	}
	return nil
}
