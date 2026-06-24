package todo

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xXMolinaXx/golang/internal/domain"
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
		var todo domain.Todo
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
		todo, err := s.CreateTodo(requestBody.Title, requestBody.Description, userID)
		if err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
			return
		}
		c.JSON(http.StatusCreated, todo)
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
		pageStr := c.Query("page")
		limitStr := c.Query("limit")
		if pageStr == "" && limitStr == "" {
			c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Pagination is not implemented yet"})
			return
		}
		page, err := strconv.Atoi(pageStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid page value"})
			return
		}
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid limit value"})
			return
		}
		todos, total, err := s.ReadAllTodos(TodoFilter{
			Page:  int(page),
			Limit: int(limit),
		})
		if err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"todos": todos,
			"meta:": gin.H{
				"page":  page,
				"limit": limit,
				"total": total,
			},
		})
	}
}

func updateTodo(s *TodoService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestBody CreateTodoRequest
		todoId := c.Param("id")
		userId := c.GetString("userId")
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
			return
		}
		if requestBody.Title == "" || requestBody.Description == "" {
			c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Title and Description are required"})
			return
		}
		payload := domain.Todo{
			Id:          todoId,
			Title:       requestBody.Title,
			Description: requestBody.Description,
			UserId:      userId,
		}
		if err := s.UpdateTodo(payload); err != nil {
			c.JSON(http.StatusForbidden, ErrorResponse{Message: "Forbidden"})
			return
		}
		c.JSON(http.StatusOK, payload)
	}
}

func deleteTodo(s *TodoService) gin.HandlerFunc {
	return func(c *gin.Context) {
		todoId := c.Param("id")
		userId := c.GetString("userId")
		if err := s.DeleteTodo(todoId, userId); err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
			return
		}
		c.Status(http.StatusNoContent)
	}
}
