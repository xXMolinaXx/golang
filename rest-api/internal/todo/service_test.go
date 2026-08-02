package todo_test

import (
	"testing"

	"github.com/xXMolinaXx/golang/internal/domain"
	"github.com/xXMolinaXx/golang/internal/todo"
)

type repoMock struct{}

func createMockTodoRepository() todo.Repository {
	return &repoMock{}
}
func (r *repoMock) CreateTodo(todo *domain.Todo) error {
	todo.Id = "mocked-id"
	return nil
}
func (r *repoMock) ReadTodo(id, userId string) (*domain.Todo, error) {
	return &domain.Todo{
		Id:          id,
		Title:       "Mocked Title",
		Description: "Mocked Description",
		UserId:      userId,
	}, nil
}
func (r *repoMock) ReadAllTodos(filter todo.TodoFilter) ([]domain.Todo, int, error) {
	todos := []domain.Todo{
		{Id: "1", Title: "Todo 1", Description: "Description 1", UserId: "user123"},
		{Id: "2", Title: "Todo 2", Description: "Description 2", UserId: "user123"},
	}
	return todos, len(todos), nil
}
func (r *repoMock) UpdateTodo(todo domain.Todo) error {
	return nil
}
func (r *repoMock) DeleteTodo(id, userId string) error {
	return nil
}
func createTodoService() *todo.TodoService {
	repo := createMockTodoRepository()
	todoService := todo.NewTodoService(repo)
	return todoService
}

func TestCreateTodo(t *testing.T) {
	todoService := createTodoService()
	todo, err := todoService.CreateTodo("Test Title", "Test Description", "user123")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if todo.Title != "Test Title" || todo.Description != "Test Description" || todo.UserId != "user123" {
		t.Fatalf("Expected todo with title 'Test Title', description 'Test Description' and userId 'user123', got %+v", todo)
	}
}
func TestReadTodo(t *testing.T) {
	todoService := createTodoService()
	todo, err := todoService.ReadTodo("mocked-id", "user123")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if todo.Id != "mocked-id" || todo.UserId != "user123" {
		t.Fatalf("Expected todo with id 'mocked-id' and userId 'user123', got %+v", todo)
	}
}
func TestReadAllTodos(t *testing.T) {
	todoService := createTodoService()
	filter := todo.TodoFilter{Page: 1, Limit: 10}
	todos, total, err := todoService.ReadAllTodos(filter)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if total != 2 {
		t.Fatalf("Expected total 2, got %d", total)
	}
	if len(todos) != 2 {
		t.Fatalf("Expected 2 todos, got %d", len(todos))
	}
}

func TestUpdateTodo(t *testing.T) {
	todoService := createTodoService()
	todo := domain.Todo{Id: "mocked-id", Title: "Updated Title", Description: "Updated Description", UserId: "user123"}
	err := todoService.UpdateTodo(todo)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestDeleteTodo(t *testing.T) {
	todoService := createTodoService()
	err := todoService.DeleteTodo("mocked-id", "user123")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}
