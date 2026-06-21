package todo

import (
	"fmt"

	"github.com/xXMolinaXx/golang/internal/domain"
)

type Service interface {
	CreateTodo(title, description, userId string) error
	ReadTodo(id, userId string) (*domain.Todo, error)
	ReadAllTodos() ([]domain.Todo, error)
	UpdateTodo(todo *domain.Todo) error
	DeleteTodo(id, userId string) error
}

type TodoService struct {
	db Repository
}

func NewTodoService(db Repository) *TodoService {
	return &TodoService{db: db} // crea la variable y retorna un puntero
}

func (s *TodoService) CreateTodo(title, description, userId string) error {
	if title == "" || description == "" || userId == "" {
		return fmt.Errorf("title, description and user_id are required")
	}

	todo := domain.Todo{
		Title:       title,
		Description: description,
		UserId:      userId,
	}
	if err := s.db.CreateTodo(&todo); err != nil {
		return err
	}
	return nil
}

func (s *TodoService) ReadTodo(id, userId string) (*domain.Todo, error) {

	return s.db.ReadTodo(id, userId)
}
func (s *TodoService) ReadAllTodos() ([]domain.Todo, error) {
	return s.db.ReadAllTodos()
}
func (s *TodoService) UpdateTodo(todo *domain.Todo) error {
	return s.db.UpdateTodo(todo)
}
func (s *TodoService) DeleteTodo(id, userId string) error {
	return s.db.DeleteTodo(id, userId)
}
