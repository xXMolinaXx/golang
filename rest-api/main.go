package main

import (
	"log"
	"os"
	"rest/api/src/user"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	databaseName := os.Getenv("DATABASE_NAME")
	db, err := gorm.Open(sqlite.Open(databaseName), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db = db.Debug()
	// Migrate the schema
	db.AutoMigrate(&user.User{})
	r := gin.Default()
	userRepo := user.NewUserRepository(db)
	userService := user.NewUserService(userRepo)
	userEndpoints := user.MakeEndPoint(userService)
	r.POST("/users", userEndpoints.CreateUser)
	r.GET("/users/:id", userEndpoints.GetUser)
	r.GET("/users", userEndpoints.GetAllUsers)
	r.PUT("/users/:id", userEndpoints.UpdateUser)
	r.DELETE("/users/:id", userEndpoints.DeleteUser)
	r.Run()
}
