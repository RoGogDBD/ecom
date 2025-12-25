package service

type (
	Storage     interface{}
	TodoService struct {
		storage Storage
	}
)

func NewTodoService(storage Storage) *TodoService {
	return &TodoService{storage: storage}
}
