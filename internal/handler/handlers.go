package handler

import (
	"net/http"
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

// @Summary Создать задачу
// @Tags todos
// @Accept json
// @Produce json
// @Param todo body Todo true "Todo"
// @Success 201 {object} Todo
// @Failure 400 {object} ErrorResponse
// @Failure 409 {object} ErrorResponse
// @Router /todos [post]
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

// @Summary Получить список всех задач
// @Tags todos
// @Produce json
// @Success 200 {array} Todo
// @Router /todos [get]
func (r *Router) handleGetAll(w http.ResponseWriter, req *http.Request) {
	items, err := r.service.GetAll(req.Context())
	if err != nil {
		writeServiceError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, items)
}

// @Summary Получить задачу по ID
// @Tags todos
// @Produce json
// @Param id path int true "Todo ID"
// @Success 200 {object} Todo
// @Failure 404 {object} ErrorResponse
// @Router /todos/{id} [get]
func (r *Router) handleGetByID(w http.ResponseWriter, req *http.Request, id int) {
	item, err := r.service.GetByID(req.Context(), id)
	if err != nil {
		writeServiceError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, item)
}

// @Summary Обновить задачу по ID
// @Tags todos
// @Accept json
// @Produce json
// @Param todo body Todo true "Todo"
// @Param id path int true "Todo ID"
// @Success 200 {object} Todo
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /todos/{id} [put]
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

// @Summary Удалить задачу по ID
// @Tags todos
// @Param id path int true "Todo ID"
// @Success 204
// @Failure 404 {object} ErrorResponse
// @Router /todos/{id} [delete]
func (r *Router) handleDelete(w http.ResponseWriter, req *http.Request, id int) {
	if err := r.service.Delete(req.Context(), id); err != nil {
		writeServiceError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func swaggerHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

}
