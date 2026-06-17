package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	EndPoint struct {
		CreateUser  gin.HandlerFunc
		GetUser     gin.HandlerFunc
		GetAllUsers gin.HandlerFunc
		UpdateUser  gin.HandlerFunc
		DeleteUser  gin.HandlerFunc
	}
	CreateUserRequest struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	ErrorResponse struct {
		Message string `json:"message"`
	}
)

func MakeEndPoint(s *UserService) EndPoint {
	return EndPoint{
		CreateUser:  createUser(s),
		GetUser:     getUser(s),
		GetAllUsers: getAllUsers(s),
		UpdateUser:  updateUser(s),
		DeleteUser:  deleteUser(s),
	}
}

func createUser(s *UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestBody CreateUserRequest
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Failed to parse JSON"})
			return
		}
		if requestBody.Name == "" || requestBody.Email == "" {
			c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Name and Email are required"})
			return
		}
		if err := s.CreateUser(requestBody.Name, requestBody.Email, requestBody.Password); err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Failed to create user"})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
	}
}

func getUser(s *UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			c.JSON(http.StatusBadRequest, ErrorResponse{Message: "User id is required"})
			return
		}
		user, err := s.ReadUser(id)
		if err != nil {
			c.JSON(http.StatusNotFound, ErrorResponse{Message: "User not found"})
			return
		}
		c.JSON(http.StatusOK, user)
	}
}

func getAllUsers(s *UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := s.ReadAllUsers()
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Failed to retrieve users"})
			return
		}
		c.JSON(http.StatusOK, users)
	}
}

func updateUser(s *UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			c.JSON(http.StatusBadRequest, ErrorResponse{Message: "User id is required"})
			return
		}
		var requestBody CreateUserRequest
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Failed to parse JSON"})
			return
		}
		user, err := s.ReadUser(id)
		if err != nil {
			c.JSON(http.StatusNotFound, ErrorResponse{Message: "User not found"})
			return
		}
		if err := s.UpdateUser(user, requestBody); err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Failed to update user"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
	}
}

func deleteUser(s *UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			c.JSON(http.StatusBadRequest, ErrorResponse{Message: "User id is required"})
			return
		}
		if err := s.DeleteUser(id); err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Failed to delete user"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
	}
}
