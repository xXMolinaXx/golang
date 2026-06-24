package main

import (
	"log"
	"time"

	ratelimit "github.com/JGLTechnologies/gin-rate-limit"
	"github.com/xXMolinaXx/golang/internal/todo"

	"github.com/xXMolinaXx/golang/internal/user"
	"github.com/xXMolinaXx/golang/pkg/bootstrap"
	"github.com/xXMolinaXx/golang/pkg/security"

	"github.com/cnjack/throttle"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func keyFunc(c *gin.Context) string {
	return c.ClientIP()
}

func errorHandler(c *gin.Context, info ratelimit.Info) {
	c.String(429, "Too many requests. Try again in "+time.Until(info.ResetTime).String())
}
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
	// Rate limiting middleware
	store := ratelimit.InMemoryStore(&ratelimit.InMemoryOptions{
		Rate:  time.Minute,
		Limit: 5,
	})
	mw := ratelimit.RateLimiter(store, &ratelimit.Options{
		ErrorHandler: errorHandler,
		KeyFunc:      keyFunc,
	})
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
	protected := r.Group("/", mw)
	protected.Use(security.AuthMiddleware())
	// throttle policy: 1 request per hour per user
	protected.Use(throttle.Policy(&throttle.Quota{
		Limit:  1,
		Within: time.Hour,
	}))
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
