package repository

import (
	"context"
)

type Repository interface {
	GetByID(ctx context.Context, id interface{}) (interface{}, error)
	GetAll(ctx context.Context) ([]interface{}, error)
	Create(ctx context.Context, obj interface{}) (interface{}, error)
	Update(ctx context.Context, id interface{}, obj interface{}) error
	Delete(ctx context.Context, id interface{}) error
}
