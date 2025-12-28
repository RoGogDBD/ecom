package models

import "errors"

var (
	// Ошибки валидации данных.
	ErrInvalidID  = errors.New("id должен быть положительным числом")
	ErrEmptyTitle = errors.New("title не может быть пустым")
	// Ошибки операций.
	ErrDuplicateID = errors.New("todo с данным ID уже существует")
	ErrNotFound    = errors.New("todo не найден")
)

type (
	Todo struct {
		ID          int    `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Completed   bool   `json:"completed"`
	}
)
