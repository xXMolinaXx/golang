package user

import (
	"time"
)

type User struct {
	Id        string     `json:"id" gorm:"type:char(50);not null;primary_key;unique"`
	Fullname  string     `json:"fullname" gorm:"type:varchar(255);not null"`
	Email     string     `json:"email" gorm:"type:varchar(255);not null;unique"`
	Password  string     `json:"password" gorm:"type:varchar(255);not null"`
	CreatedAt *time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
