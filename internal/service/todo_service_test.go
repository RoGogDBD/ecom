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
			name:            "success",
			todo:            models.Todo{ID: 1, Title: "ship order", Description: "pack items", Completed: false},
			wantCreateCalls: 1,
		},
		{
			name:    "validation error empty title",
			todo:    models.Todo{ID: 2, Title: "  ", Description: "missing title"},
			wantErr: models.ErrEmptyTitle,
		},
		{
			name:    "validation error invalid id",
			todo:    models.Todo{ID: 0, Title: "title"},
			wantErr: models.ErrInvalidID,
		},
		{
			name:            "duplicate id",
			todo:            models.Todo{ID: 3, Title: "duplicate"},
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
				t.Fatalf("expected error %v, got %v", tc.wantErr, err)
			}
			if storage.createCalls != tc.wantCreateCalls {
				t.Fatalf("expected create calls %d, got %d", tc.wantCreateCalls, storage.createCalls)
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
			name:            "success",
			todo:            models.Todo{ID: 1, Title: "rename", Description: "new", Completed: true},
			wantUpdateCalls: 1,
		},
		{
			name:    "validation error empty title",
			todo:    models.Todo{ID: 2, Title: ""},
			wantErr: models.ErrEmptyTitle,
		},
		{
			name:    "validation error invalid id",
			todo:    models.Todo{ID: -1, Title: "title"},
			wantErr: models.ErrInvalidID,
		},
		{
			name:            "not found",
			todo:            models.Todo{ID: 3, Title: "missing"},
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
				t.Fatalf("expected error %v, got %v", tc.wantErr, err)
			}
			if storage.updateCalls != tc.wantUpdateCalls {
				t.Fatalf("expected update calls %d, got %d", tc.wantUpdateCalls, storage.updateCalls)
			}
		})
	}
}
