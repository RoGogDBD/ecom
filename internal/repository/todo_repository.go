package repository

import (
	"errors"
	"sync"

	"github.com/RoGogDBD/ecom/internal/models"
)

var (
	errDuplicateID = errors.New("todo с данным ID уже существует")
	errNotFound    = errors.New("todo не найден")
)

type (
	// TodoStorage потокобезопасное хранилище для объектов.
	TodoStorage struct {
		mu    sync.RWMutex
		items map[int]models.Todo
	}
)

// NewToDoStorage создает и возвращает новый экземпляр ToDoStorage.
func NewToDoStorage() *TodoStorage {
	return &TodoStorage{
		items: make(map[int]models.Todo),
	}
}

// Create добавляет новый объект в хранилище.
func (s *TodoStorage) Create(todo models.Todo) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.items[todo.ID]; exists {
		return errDuplicateID
	}

	s.items[todo.ID] = todo
	return nil
}

// Update обновляет существующий объект в хранилище.
func (s *TodoStorage) Update(todo models.Todo) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.items[todo.ID]; !exists {
		return errDuplicateID
	}

	s.items[todo.ID] = todo
	return nil
}

// Delete удаляет объект из хранилища по его ID.
func (s *TodoStorage) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.items[id]; !exists {
		return errNotFound
	}

	delete(s.items, id)
	return nil
}

// GetAll возвращает все объекты из хранилища.
func (s *TodoStorage) GetAll() []models.Todo {
	s.mu.RLock()
	defer s.mu.RUnlock()

	result := make([]models.Todo, 0, len(s.items))
	for _, todo := range s.items {
		result = append(result, todo)
	}

	return result
}

// GetByID возвращает объект по его ID.
func (s *TodoStorage) GetByID(id int) (models.Todo, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	todo, exists := s.items[id]
	if !exists {
		return models.Todo{}, errNotFound
	}

	return todo, nil
}
