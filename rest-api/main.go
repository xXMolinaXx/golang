package main

import (
	"log"
	"rest/api/internal/user"
	"rest/api/pkg/bootstrap"

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
	userRepo := user.NewUserRepository(db)
	userService := user.NewUserService(userRepo)
	userEndpoints := user.MakeEndPoint(userService)
	r.POST("/users", userEndpoints.CreateUser)
	r.GET("/users/:id", userEndpoints.GetUser)
	r.GET("/users", userEndpoints.GetAllUsers)
	r.PUT("/users/:id", userEndpoints.UpdateUser)
	r.DELETE("/users/:id", userEndpoints.DeleteUser)
	r.Run(":3000")
}
