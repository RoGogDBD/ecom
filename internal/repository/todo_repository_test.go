package repository

import (
	"context"
	"errors"
	"testing"

	"github.com/RoGogDBD/ecom/internal/models"
)

func TestTodoStorageCreate(t *testing.T) {
	cases := []struct {
		name    string
		initial []models.Todo
		todo    models.Todo
		wantErr error
	}{
		{
			name: "успешное создание",
			todo: models.Todo{ID: 1, Title: "новая", Description: "создана", Completed: false},
		},
		{
			name:    "дубликат id",
			initial: []models.Todo{{ID: 1, Title: "существующая"}},
			todo:    models.Todo{ID: 1, Title: "дубликат"},
			wantErr: models.ErrDuplicateID,
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			storage := NewTodoStorage()
			ctx := context.Background()
			for _, item := range tc.initial {
				if err := storage.Create(ctx, item); err != nil {
					t.Fatalf("ошибка подготовки данных: %v", err)
				}
			}

			err := storage.Create(ctx, tc.todo)
			if !errors.Is(err, tc.wantErr) {
				t.Fatalf("ожидалась ошибка %v, получено %v", tc.wantErr, err)
			}

			if tc.wantErr == nil {
				got, err := storage.GetByID(ctx, tc.todo.ID)
				if err != nil {
					t.Fatalf("ожидалась сохраненная задача, получена ошибка: %v", err)
				}
				if got != tc.todo {
					t.Fatalf("ожидалась задача %+v, получено %+v", tc.todo, got)
				}
			}
		})
	}
}
