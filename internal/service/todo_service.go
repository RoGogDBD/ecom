package service

import (
	"context"
	"errors"
	"strings"

	"github.com/RoGogDBD/ecom/internal/models"
)

var (
	errInvalidID  = errors.New("id должен быть положительным числом")
	errEmptyTitle = errors.New("title не может быть пустым")
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

// TODO: func (s *TodoService) Update(ctx context.Context, todo models.Todo) error {}.
// TODO: func (s *TodoService) Delete(ctx context.Context, id int) error {}.
// TODO: func (s *TodoService) GetAll(ctx context.Context) []models.Todo {}.
// TODO: func (s *TodoService) GetByID(ctx context.Context, id int) (models.Todo, error) {}.

// ******************
// Хелпующие функции.
// ******************

// validateTodo проверяет корректность данных.
func validateTodo(todo models.Todo) error {
	if todo.ID <= 0 {
		return errInvalidID
	}

	if strings.TrimSpace(todo.Title) == "" {
		return errEmptyTitle
	}

	return nil
}
