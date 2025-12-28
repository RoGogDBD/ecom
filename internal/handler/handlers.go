package handler

import (
	"net/http"

	"github.com/RoGogDBD/ecom/internal/models"
)

const notFoundMessage = "todo not found"

func (r *Router) handleTodos(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		r.handleGetAll(w, req)
	case http.MethodPost:
		r.handleCreate(w, req)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (r *Router) handleTodoByID(w http.ResponseWriter, req *http.Request) {
	id, ok := parseID(req.URL.Path)
	if !ok {
		writeError(w, http.StatusNotFound, notFoundMessage)
		return
	}

	switch req.Method {
	case http.MethodGet:
		r.handleGetByID(w, req, id)
	case http.MethodPut:
		r.handleUpdate(w, req, id)
	case http.MethodDelete:
		r.handleDelete(w, req, id)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (r *Router) handleCreate(w http.ResponseWriter, req *http.Request) {
	todo, err := decodeTodo(req)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := r.service.Create(req.Context(), todo); err != nil {
		writeServiceError(w, err)
		return
	}

	writeJSON(w, http.StatusCreated, todo)
}

func (r *Router) handleGetAll(w http.ResponseWriter, req *http.Request) {
	items, err := r.service.GetAll(req.Context())
	if err != nil {
		writeServiceError(w, err)
		return
	}

	if items == nil {
		items = []models.Todo{}
	}

	writeJSON(w, http.StatusOK, items)
}

func (r *Router) handleGetByID(w http.ResponseWriter, req *http.Request, id int) {
	item, err := r.service.GetByID(req.Context(), id)
	if err != nil {
		writeServiceError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, item)
}

func (r *Router) handleUpdate(w http.ResponseWriter, req *http.Request, id int) {
	todo, err := decodeTodo(req)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	todo.ID = id

	if err := r.service.Update(req.Context(), todo); err != nil {
		writeServiceError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, todo)
}

func (r *Router) handleDelete(w http.ResponseWriter, req *http.Request, id int) {
	if err := r.service.Delete(req.Context(), id); err != nil {
		writeServiceError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// func swaggerHandler(w http.ResponseWriter, req *http.Request) {
// 	if req.Method != http.MethodGet {
// 		w.WriteHeader(http.StatusMethodNotAllowed)
// 		return
// 	}

// }
