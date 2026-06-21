package domain

import (
	"time"

	"github.com/google/uuid"

	"gorm.io/gorm"
)

type Todo struct {
	Id          string     `json:"id" gorm:"type:char(50);not null;primary_key;unique"`
	Title       string     `json:"title" gorm:"type:char(50);not null;"`
	Description string     `json:"description" gorm:"type:char(100);not null;"`
	UserId      string     `json:"user_id" gorm:"type:char(50);not null;index"`
	User        *User      `json:"user,omitempty" gorm:"foreignKey:UserId;references:Id"`
	CreatedAt   *time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   *time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (u *Todo) BeforeCreate(tx *gorm.DB) (err error) {
	if u.Id == "" {
		u.Id = uuid.New().String()
	}
	return
}
