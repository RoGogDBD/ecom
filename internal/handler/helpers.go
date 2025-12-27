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

const (
	todosPathPrefix = "/todos/"

	contentTypeHeader = "Content-Type"
	contentTypeJSON   = "application/json"

	invalidJSONPayloadMsg  = "invalid JSON payload"
	internalServerErrorMsg = "internal server error"

	jsonErrorKey = "error"
)

func parseID(path string) (int, bool) {
	trimmed := strings.TrimPrefix(path, todosPathPrefix)
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
		return models.Todo{}, errors.New(invalidJSONPayloadMsg)
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
		writeError(w, http.StatusInternalServerError, internalServerErrorMsg)
	}
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{jsonErrorKey: message})
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set(contentTypeHeader, contentTypeJSON)
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}
