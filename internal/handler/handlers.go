package handler

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"

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

// ******************
// Хелпующие функции.
// ******************

func parseID(path string) (int, bool) {
	trimmed := strings.TrimPrefix(path, "/todos/")
	if trimmed == "" || strings.Contains(trimmed, "/") {
		return 0, false
	}

	id, err := strconv.Atoi(trimmed)
	if err != nil {
		return 0, false
	}

	return id, true
}

func decodeTodo(req *http.Request) (models.Todo, error) {
	defer req.Body.Close()

	dec := json.NewDecoder(req.Body)
	dec.DisallowUnknownFields()

	var todo models.Todo
	if err := dec.Decode(&todo); err != nil {
		return models.Todo{}, err
	}

	if err := dec.Decode(&struct{}{}); !errors.Is(err, io.EOF) {
		return models.Todo{}, errors.New("invalid JSON payload")
	}

	return todo, nil
}

func writeServiceError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, models.ErrInvalidID), errors.Is(err, models.ErrEmptyTitle):
		writeError(w, http.StatusBadRequest, err.Error())
	case errors.Is(err, models.ErrDuplicateID):
		writeError(w, http.StatusConflict, err.Error())
	case errors.Is(err, models.ErrNotFound):
		writeError(w, http.StatusNotFound, err.Error())
	default:
		writeError(w, http.StatusInternalServerError, "internal server error")
	}
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}
