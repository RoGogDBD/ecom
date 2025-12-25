package service

import (
	"context"

	"github.com/RoGogDBD/ecom/internal/models"
)

type (
	Storage interface {
		Create(ctx context.Context, todo models.Todo) error
		Update(ctx context.Context, todo models.Todo) error
		Delete(ctx context.Context, id int) error
		GetAll(ctx context.Context) ([]models.Todo, error)
		GetByID(ctx context.Context, id int) (models.Todo, error)
	}
	TodoService struct {
		storage Storage
	}
)

func NewTodoService(storage Storage) *TodoService {
	return &TodoService{storage: storage}
}

// TODO: func (s *TodoService) Create(ctx context.Context, todo models.Todo) error {}.
// TODO: func (s *TodoService) Update(ctx context.Context, todo models.Todo) error {}.
// TODO: func (s *TodoService) Delete(ctx context.Context, id int) error {}.
// TODO: func (s *TodoService) GetAll(ctx context.Context) []models.Todo {}.
// TODO: func (s *TodoService) GetByID(ctx context.Context, id int) (models.Todo, error) {}.
// TODO: func validateTodo(todo models.Todo) error {}.
