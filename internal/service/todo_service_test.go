package service

import (
	"context"
	"errors"
	"testing"

	"github.com/RoGogDBD/ecom/internal/models"
)

type stubStorage struct {
	createErr   error
	updateErr   error
	createCalls int
	updateCalls int
}

func (s *stubStorage) Create(_ context.Context, _ models.Todo) error {
	s.createCalls++
	return s.createErr
}

func (s *stubStorage) Update(_ context.Context, _ models.Todo) error {
	s.updateCalls++
	return s.updateErr
}

func (s *stubStorage) Delete(_ context.Context, _ int) error {
	return nil
}

func (s *stubStorage) GetAll(_ context.Context) ([]models.Todo, error) {
	return nil, nil
}

func (s *stubStorage) GetByID(_ context.Context, _ int) (models.Todo, error) {
	return models.Todo{}, nil
}

func TestTodoServiceCreate(t *testing.T) {
	cases := []struct {
		name            string
		todo            models.Todo
		createErr       error
		wantErr         error
		wantCreateCalls int
	}{
		{
			name:            "успешное создание",
			todo:            models.Todo{ID: 1, Title: "отправить заказ", Description: "упаковать товары", Completed: false},
			wantCreateCalls: 1,
		},
		{
			name:    "ошибка валидации: пустой заголовок",
			todo:    models.Todo{ID: 2, Title: "  ", Description: "нет заголовка"},
			wantErr: models.ErrEmptyTitle,
		},
		{
			name:    "ошибка валидации: неверный id",
			todo:    models.Todo{ID: 0, Title: "заголовок"},
			wantErr: models.ErrInvalidID,
		},
		{
			name:            "дубликат id",
			todo:            models.Todo{ID: 3, Title: "дубликат"},
			createErr:       models.ErrDuplicateID,
			wantErr:         models.ErrDuplicateID,
			wantCreateCalls: 1,
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			storage := &stubStorage{createErr: tc.createErr}
			service := NewTodoService(storage)

			err := service.Create(context.Background(), tc.todo)
			if !errors.Is(err, tc.wantErr) {
				t.Fatalf("ожидалась ошибка %v, получено %v", tc.wantErr, err)
			}
			if storage.createCalls != tc.wantCreateCalls {
				t.Fatalf("ожидалось вызовов создания %d, получено %d", tc.wantCreateCalls, storage.createCalls)
			}
		})
	}
}

func TestTodoServiceUpdate(t *testing.T) {
	cases := []struct {
		name            string
		todo            models.Todo
		updateErr       error
		wantErr         error
		wantUpdateCalls int
	}{
		{
			name:            "успешное обновление",
			todo:            models.Todo{ID: 1, Title: "переименовать", Description: "новое", Completed: true},
			wantUpdateCalls: 1,
		},
		{
			name:    "ошибка валидации: пустой заголовок",
			todo:    models.Todo{ID: 2, Title: ""},
			wantErr: models.ErrEmptyTitle,
		},
		{
			name:    "ошибка валидации: неверный id",
			todo:    models.Todo{ID: -1, Title: "заголовок"},
			wantErr: models.ErrInvalidID,
		},
		{
			name:            "не найдено",
			todo:            models.Todo{ID: 3, Title: "нет записи"},
			updateErr:       models.ErrNotFound,
			wantErr:         models.ErrNotFound,
			wantUpdateCalls: 1,
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			storage := &stubStorage{updateErr: tc.updateErr}
			service := NewTodoService(storage)

			err := service.Update(context.Background(), tc.todo)
			if !errors.Is(err, tc.wantErr) {
				t.Fatalf("ожидалась ошибка %v, получено %v", tc.wantErr, err)
			}
			if storage.updateCalls != tc.wantUpdateCalls {
				t.Fatalf("ожидалось вызовов обновления %d, получено %d", tc.wantUpdateCalls, storage.updateCalls)
			}
		})
	}
}
