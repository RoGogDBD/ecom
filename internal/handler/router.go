package handler

import (
	"context"
	"net/http"

	"github.com/RoGogDBD/ecom/internal/models"
)

type (
	TodoService interface {
		Create(ctx context.Context, todo models.Todo) error
		Update(ctx context.Context, todo models.Todo) error
		Delete(ctx context.Context, id int) error
		GetAll(ctx context.Context) ([]models.Todo, error)
		GetByID(ctx context.Context, id int) (models.Todo, error)
	}

	Router struct {
		service TodoService
	}
)

func NewRouter(service TodoService) http.Handler {
	r := &Router{service: service}
	mux := http.NewServeMux()

	mux.HandleFunc("/todos", r.handleTodos)
	mux.HandleFunc("/todos/", r.handleTodoByID)
	// mux.HandleFunc("/swagger.json", swaggerHandler)

	return mux
}
