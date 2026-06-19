package main

import (
	"log"

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

	// Public endpoints
	r.POST("/users", userEndpoints.CreateUser)
	r.POST("/login", userEndpoints.Login)

	// Protected endpoints
	protected := r.Group("/")
	protected.Use(security.AuthMiddleware())
	protected.GET("/users/:id", userEndpoints.GetUser)
	protected.GET("/users/", userEndpoints.GetAllUsers)
	protected.PUT("/users/:id", userEndpoints.UpdateUser)
	protected.DELETE("/users/:id", userEndpoints.DeleteUser)

	r.Run(":3000")
}
