package todo

import (
	"github.com/xXMolinaXx/golang/internal/domain"

	"gorm.io/gorm"
)

type Repository interface {
	CreateTodo(todo *domain.Todo) error
	ReadTodo(id, userId string) (*domain.Todo, error)
	ReadAllTodos() ([]domain.Todo, error)
	UpdateTodo(todo *domain.Todo) error
	DeleteTodo(id string) error
}

type repo struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) Repository {
	return &repo{db: db}
}

func (r *repo) CreateTodo(todo *domain.Todo) error {
	err := r.db.Create(todo).Error
	return err
}
func (r *repo) ReadTodo(id, userId string) (*domain.Todo, error) {
	var todo domain.Todo
	err := r.db.Select("id", "title", "description", "created_at").First(&todo, "id = ? AND user_id = ?", id, userId).Error
	if err != nil {
		return nil, err
	}
	return &todo, nil

}
func (r *repo) ReadAllTodos() ([]domain.Todo, error) {
	var todos []domain.Todo
	err := r.db.Select("id", "title", "description", "user_id", "created_at", "updated_at").Find(&todos).Error
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (r *repo) UpdateTodo(todo *domain.Todo) error {
	err := r.db.Save(todo).Error
	return err
}

func (r *repo) DeleteTodo(id string) error {
	err := r.db.Delete(&domain.Todo{}, "id = ?", id).Error
	return err
}
