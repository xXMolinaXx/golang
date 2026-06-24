package todo

import (
	"github.com/xXMolinaXx/golang/internal/domain"

	"gorm.io/gorm"
)

type Repository interface {
	CreateTodo(todo *domain.Todo) error
	ReadTodo(id, userId string) (*domain.Todo, error)
	ReadAllTodos(filter TodoFilter) ([]domain.Todo, int, error)
	UpdateTodo(todo domain.Todo) error
	DeleteTodo(id, userId string) error
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
func (r *repo) ReadAllTodos(filter TodoFilter) ([]domain.Todo, int, error) {
	var todos []domain.Todo
	offset := (filter.Page - 1) * filter.Limit
	err := r.db.Select("id", "title", "description", "user_id", "created_at", "updated_at").
		Order("created_at DESC").
		Limit(filter.Limit).
		Offset(offset).Find(&todos).Error
	if err != nil {
		return nil, 0, err
	}
	var total int64
	r.db.Model(&domain.Todo{}).Count(&total)
	return todos, int(total), nil
}

func (r *repo) UpdateTodo(todo domain.Todo) error {
	var todoFound domain.Todo
	err := r.db.Model(&todoFound).Where("id = ? AND user_id = ?", todo.Id, todo.UserId).Updates(map[string]interface{}{
		"title":       todo.Title,
		"description": todo.Description,
	}).Error
	return err
}

func (r *repo) DeleteTodo(id, userId string) error {
	tx := r.db.Delete(&domain.Todo{}, "id = ? AND user_id = ?", id, userId)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
