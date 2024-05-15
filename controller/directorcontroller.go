package controller

import (
	"fmt"

	"github.com/vladsendrix/go-movies/entities"
	"github.com/vladsendrix/go-movies/repository"
)

type DirectorController struct {
	DirectorRepo *repository.DirectorRepository
}

func NewDirectorController(directorRepo *repository.DirectorRepository) *DirectorController {
	return &DirectorController{DirectorRepo: directorRepo}
}

func (c *DirectorController) GetByID(id int) (*entities.Director, error) {
	director, err := c.DirectorRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	directorEntity, ok := director.(*entities.Director)
	if !ok {
		return nil, fmt.Errorf("could not assert type: expected *entities.Director")
	}

	return directorEntity, nil
}

func (c *DirectorController) GetAll() ([]*entities.Director, error) {
	directors, err := c.DirectorRepo.GetAll()
	if err != nil {
		return nil, err
	}

	directorEntities := make([]*entities.Director, len(directors))
	for i, director := range directors {
		directorEntity, ok := director.(*entities.Director)
		if !ok {
			return nil, fmt.Errorf("could not assert type: expected *entities.Director")
		}
		directorEntities[i] = directorEntity
	}

	return directorEntities, nil
}

func (c *DirectorController) Create(director *entities.Director) error {
	err := c.DirectorRepo.Create(director)
	if err != nil {
		return err
	}
	return nil
}

func (c *DirectorController) Delete(id interface{}) error {
	err := c.DirectorRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (c *DirectorController) Update(id int, director *entities.Director) error {
	err := c.DirectorRepo.Update(id, director)
	if err != nil {
		return err
	}
	return nil
}
