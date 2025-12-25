package repository

import (
	"errors"
	"sync"

	"github.com/RoGogDBD/ecom/internal/models"
)

var (
	errDuplicateID = errors.New("todo с данным ID уже существует")
)

type (
	ToDoStorage struct {
		mu    sync.RWMutex
		items map[int]models.ToDo
	}
)

func NewToDoStorage() *ToDoStorage {
	return &ToDoStorage{
		items: make(map[int]models.ToDo),
	}
}

func (s *ToDoStorage) Create(todo models.ToDo) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.items[todo.ID]; exists {
		return errDuplicateID
	}

	s.items[todo.ID] = todo
	return nil
}

// TODO: func (s *ToDoStorage) GetAll() []models.ToDo {}

// TODO: func (s *ToDoStorage) GetByID(id int) (models.ToDo, error) {}

// TODO: func (s *TodoStorage) Update(todo models.ToDo) error {}

// TODO: func (s *TodoStorage) Delete(id int) error {}
