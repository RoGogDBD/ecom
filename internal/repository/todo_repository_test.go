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
			name: "success",
			todo: models.Todo{ID: 1, Title: "new", Description: "created", Completed: false},
		},
		{
			name:    "duplicate id",
			initial: []models.Todo{{ID: 1, Title: "existing"}},
			todo:    models.Todo{ID: 1, Title: "duplicate"},
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
					t.Fatalf("setup create failed: %v", err)
				}
			}

			err := storage.Create(ctx, tc.todo)
			if !errors.Is(err, tc.wantErr) {
				t.Fatalf("expected error %v, got %v", tc.wantErr, err)
			}

			if tc.wantErr == nil {
				got, err := storage.GetByID(ctx, tc.todo.ID)
				if err != nil {
					t.Fatalf("expected stored todo, got error: %v", err)
				}
				if got != tc.todo {
					t.Fatalf("expected todo %+v, got %+v", tc.todo, got)
				}
			}
		})
	}
}
