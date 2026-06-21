package todo

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	EndPoint struct {
		CreateTodo   gin.HandlerFunc
		ReadTodo     gin.HandlerFunc
		ReadAllTodos gin.HandlerFunc
		UpdateTodo   gin.HandlerFunc
		DeleteTodo   gin.HandlerFunc
	}
	CreateTodoRequest struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		UserId      string `json:"user_id"`
	}
	ErrorResponse struct {
		Message string `json:"message"`
	}
)

func MakeEndPoint(s *TodoService) EndPoint {
	return EndPoint{
		CreateTodo:   createTodo(s),
		ReadTodo:     readTodo(s),
		ReadAllTodos: readAllTodos(s),
		UpdateTodo:   updateTodo(s),
		DeleteTodo:   deleteTodo(s),
	}
}

func createTodo(s *TodoService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestBody CreateTodoRequest
		userID := c.GetString("userId")
		if userID == "" {
			c.JSON(http.StatusUnauthorized, ErrorResponse{Message: "user id not found in token"})
			return
		}
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
			return
		}
		if requestBody.Title == "" || requestBody.Description == "" {
			c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Title and Description are required"})
			return
		}
		if err := s.CreateTodo(requestBody.Title, requestBody.Description, userID); err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "Todo created successfully"})
	}
}

func readTodo(s *TodoService) gin.HandlerFunc {
	return func(c *gin.Context) {
		todoId := c.Param("id")
		userId := c.GetString("userId")
		todo, err := s.ReadTodo(todoId, userId)
		if err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
			return
		}
		c.JSON(http.StatusOK, todo)
	}
}

func readAllTodos(s *TodoService) gin.HandlerFunc {
	return func(c *gin.Context) {
		todos, err := s.ReadAllTodos()
		if err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
			return
		}
		c.JSON(http.StatusOK, todos)
	}
}

func updateTodo(s *TodoService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestBody CreateTodoRequest
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
			return
		}
		if requestBody.Title == "" || requestBody.Description == "" {
			c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Title and Description are required"})
			return
		}
		if err := s.CreateTodo(requestBody.Title, requestBody.Description, requestBody.UserId); err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "Todo created successfully"})
	}
}

func deleteTodo(s *TodoService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestBody CreateTodoRequest
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
			return
		}
		if requestBody.Title == "" || requestBody.Description == "" {
			c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Title and Description are required"})
			return
		}
		if err := s.CreateTodo(requestBody.Title, requestBody.Description, requestBody.UserId); err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "Todo created successfully"})
	}
}
