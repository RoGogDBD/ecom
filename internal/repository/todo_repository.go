package repository

import (
	"context"
	"sync"

	"github.com/RoGogDBD/ecom/internal/models"
)

type (
	// TodoStorage потокобезопасное хранилище для объектов.
	TodoStorage struct {
		mu    sync.RWMutex
		items map[int]models.Todo
	}
)

// NewTodoStorage создает и возвращает новый экземпляр ToDoStorage.
func NewTodoStorage() *TodoStorage {
	return &TodoStorage{
		items: make(map[int]models.Todo),
	}
}

// Create добавляет новый объект в хранилище.
func (s *TodoStorage) Create(_ context.Context, todo models.Todo) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.items[todo.ID]; exists {
		return models.ErrDuplicateID
	}

	s.items[todo.ID] = todo
	return nil
}

// Update обновляет существующий объект в хранилище.
func (s *TodoStorage) Update(_ context.Context, todo models.Todo) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.items[todo.ID]; !exists {
		return models.ErrNotFound
	}

	s.items[todo.ID] = todo
	return nil
}

// Delete удаляет объект из хранилища по его ID.
func (s *TodoStorage) Delete(_ context.Context, id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.items[id]; !exists {
		return models.ErrNotFound
	}

	delete(s.items, id)
	return nil
}

// GetAll возвращает все объекты из хранилища.
func (s *TodoStorage) GetAll(_ context.Context) ([]models.Todo, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	result := make([]models.Todo, 0, len(s.items))
	for _, todo := range s.items {
		result = append(result, todo)
	}

	return result, nil
}

// GetByID возвращает объект по его ID.
func (s *TodoStorage) GetByID(_ context.Context, id int) (models.Todo, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	todo, exists := s.items[id]
	if !exists {
		return models.Todo{}, models.ErrNotFound
	}

	return todo, nil
}
