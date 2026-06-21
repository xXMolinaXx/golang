package main

import (
	"log"

	"github.com/xXMolinaXx/golang/internal/todo"

	"github.com/xXMolinaXx/golang/internal/user"
	"github.com/xXMolinaXx/golang/pkg/bootstrap"
	"github.com/xXMolinaXx/golang/pkg/security"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db, errDb := bootstrap.DBConnection()
	if errDb != nil {
		log.Fatal("Error connecting to the database")
	}
	r := gin.Default()
	hashService := security.NewHashService()
	jwtImpl := security.NewJwtImpl()
	userRepo := user.NewUserRepository(db)
	userService := user.NewUserService(userRepo, hashService, jwtImpl)
	userEndpoints := user.MakeEndPoint(userService)
	todoRepo := todo.NewTodoRepository(db)
	todoService := todo.NewTodoService(todoRepo)
	todoEndpoints := todo.MakeEndPoint(todoService)
	// Public endpoints
	r.POST("/register", userEndpoints.CreateUser)
	r.POST("/login", userEndpoints.Login)

	// Protected endpoints
	protected := r.Group("/")
	protected.Use(security.AuthMiddleware())
	// User endpoints
	protected.GET("/users/:id", userEndpoints.GetUser)
	protected.GET("/users/", userEndpoints.GetAllUsers)
	protected.PUT("/users/:id", userEndpoints.UpdateUser)
	protected.DELETE("/users/:id", userEndpoints.DeleteUser)

	// Todo endpoints
	protected.POST("/todos", todoEndpoints.CreateTodo)
	protected.GET("/todos/:id", todoEndpoints.ReadTodo)
	protected.GET("/todos/", todoEndpoints.ReadAllTodos)
	protected.PUT("/todos/:id", todoEndpoints.UpdateTodo)
	protected.DELETE("/todos/:id", todoEndpoints.DeleteTodo)
	r.Run(":3000")
}
