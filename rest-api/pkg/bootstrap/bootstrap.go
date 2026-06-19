package bootstrap

import (
	"os"

	"github.com/xXMolinaXx/golang/internal/domain"

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
		if err := db.AutoMigrate(&domain.User{}); err != nil {
			return nil, err
		}
	}

	return db, nil
}
