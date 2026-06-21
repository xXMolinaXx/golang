package domain

import (
	"fmt"
	"time"

	"github.com/xXMolinaXx/golang/pkg/security"

	"github.com/google/uuid"

	"gorm.io/gorm"
)

type User struct {
	Id        string     `json:"id" gorm:"type:char(50);not null;primary_key;unique"`
	Fullname  string     `json:"fullname" gorm:"type:varchar(255);not null"`
	Email     string     `json:"email" gorm:"type:varchar(255);not null;unique"`
	Password  string     `json:"-" gorm:"type:varchar(255);not null"`
	Todos     []Todo     `json:"todos,omitempty" gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt *time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	hashService := security.NewHashService()
	if u.Password == "" {
		return fmt.Errorf("password is required")
	}
	hashedPassword, err := hashService.HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = hashedPassword
	if u.Id == "" {
		u.Id = uuid.New().String()
	}
	return
}
