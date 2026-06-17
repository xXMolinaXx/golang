package bootstrap

import (
	"os"
	"rest/api/internal/user"

	"gorm.io/driver/sqlite"

	"gorm.io/gorm"
)

func DBConnection() (*gorm.DB, error) {
	databaseName := os.Getenv("DATABASE_NAME")
	db, err := gorm.Open(sqlite.Open(databaseName), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if os.Getenv("DATABASE_DEBUG") == "true" {
		db = db.Debug()
	}

	if os.Getenv("DATABASE_MIGRATE") == "true" {
		if err := db.AutoMigrate(&user.User{}); err != nil {
			return nil, err
		}
	}

	return db, nil
}
