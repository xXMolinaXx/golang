package todo

import (
	"fmt"

	"github.com/xXMolinaXx/golang/internal/domain"
)

type Service interface {
	CreateTodo(title, description, userId string) error
	ReadTodo(id, userId string) (*domain.Todo, error)
	ReadAllTodos(filter TodoFilter) ([]domain.Todo, int, error)
	UpdateTodo(todo domain.Todo) error
	DeleteTodo(id, userId string) error
}

type TodoService struct {
	db Repository
}
type TodoFilter struct {
	Page  int
	Limit int
}

func NewTodoService(db Repository) *TodoService {
	return &TodoService{db: db} // crea la variable y retorna un puntero
}

func (s *TodoService) CreateTodo(title, description, userId string) (domain.Todo, error) {
	if title == "" || description == "" || userId == "" {
		return domain.Todo{}, fmt.Errorf("title, description and user_id are required")
	}

	todo := domain.Todo{
		Title:       title,
		Description: description,
		UserId:      userId,
	}
	if err := s.db.CreateTodo(&todo); err != nil {
		return domain.Todo{}, err
	}
	return todo, nil
}

func (s *TodoService) ReadTodo(id, userId string) (*domain.Todo, error) {

	return s.db.ReadTodo(id, userId)
}
func (s *TodoService) ReadAllTodos(filter TodoFilter) ([]domain.Todo, int, error) {
	return s.db.ReadAllTodos(filter)
}
func (s *TodoService) UpdateTodo(todo domain.Todo) error {
	return s.db.UpdateTodo(todo)
}
func (s *TodoService) DeleteTodo(id, userId string) error {
	return s.db.DeleteTodo(id, userId)
}
