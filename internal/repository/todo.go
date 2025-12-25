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
	// ToDoStorage потокобезопасное хранилище для объектов models.Todo.
	ToDoStorage struct {
		mu    sync.RWMutex
		items map[int]models.Todo
	}
)

// NewToDoStorage создает и возвращает новый экземпляр ToDoStorage.
func NewToDoStorage() *ToDoStorage {
	return &ToDoStorage{
		items: make(map[int]models.Todo),
	}
}

// Create добавляет новый объект в хранилище.
func (s *ToDoStorage) Create(todo models.Todo) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.items[todo.ID]; exists {
		return errDuplicateID
	}

	s.items[todo.ID] = todo
	return nil
}

// GetAll возвращает все объекты из хранилища.
func (s *ToDoStorage) GetAll() []models.Todo {
	s.mu.RLock()
	defer s.mu.RUnlock()

	result := make([]models.Todo, 0, len(s.items))
	for _, todo := range s.items {
		result = append(result, todo)
	}

	return result
}

// GetByID возвращает объект по его ID.
func (s *ToDoStorage) GetByID(id int) (models.Todo, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	todo, exists := s.items[id]
	if !exists {
		return models.Todo{}, errNotFound
	}

	return todo, nil
}

// TODO: func (s *TodoStorage) Update(todo models.ToDo) error {}

// TODO: func (s *TodoStorage) Delete(id int) error {}
