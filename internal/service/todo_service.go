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

// TODO: func (s *TodoService) Create(todo models.Todo) error {}.
// TODO: func (s *TodoService) Update(todo models.Todo) error {}.
// TODO: func (s *TodoService) Delete(id int) error {}.
// TODO: func (s *TodoService) GetAll() []models.Todo {}.
// TODO: func (s *TodoService) GetByID(id int) (models.Todo, error) {}.
// TODO: func validateTodo(todo models.Todo) error {}.
