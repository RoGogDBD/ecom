package service

import (
	"context"
	"strings"

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

func (s *TodoService) Create(ctx context.Context, todo models.Todo) error {
	if err := validateTodo(todo); err != nil {
		return err
	}

	return s.storage.Create(ctx, todo)
}

func (s *TodoService) Update(ctx context.Context, todo models.Todo) error {
	if err := validateTodo(todo); err != nil {
		return err
	}

	return s.storage.Update(ctx, todo)
}

func (s *TodoService) Delete(ctx context.Context, id int) error {
	if id <= 0 {
		return models.ErrInvalidID
	}

	return s.storage.Delete(ctx, id)
}

func (s *TodoService) GetAll(ctx context.Context) ([]models.Todo, error) {
	return s.storage.GetAll(ctx)
}

func (s *TodoService) GetByID(ctx context.Context, id int) (models.Todo, error) {
	if id <= 0 {
		return models.Todo{}, models.ErrInvalidID
	}

	return s.storage.GetByID(ctx, id)
}

// ******************
// Хелпующие функции.
// ******************

// validateTodo проверяет корректность данных.
func validateTodo(todo models.Todo) error {
	if todo.ID <= 0 {
		return models.ErrInvalidID
	}

	if strings.TrimSpace(todo.Title) == "" {
		return models.ErrEmptyTitle
	}

	return nil
}
