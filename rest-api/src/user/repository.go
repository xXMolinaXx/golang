package user

import (
	"gorm.io/gorm"
)

type Repository interface {
	CreateUser(user *User) error
	ReadUser(id string) (*User, error)
	ReadAllUsers() ([]User, error)
	UpdateUser(user *User) error
	DeleteUser(id string) error
}

type repo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) Repository {
	return &repo{db: db}
}

func (r *repo) CreateUser(user *User) error {
	err := r.db.Create(user).Error
	return err
}
func (r *repo) ReadUser(id string) (*User, error) {
	var user User
	err := r.db.First(&user, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil

}
func (r *repo) ReadAllUsers() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *repo) UpdateUser(user *User) error {
	err := r.db.Save(user).Error
	return err
}

func (r *repo) DeleteUser(id string) error {
	err := r.db.Delete(&User{}, "id = ?", id).Error
	return err
}
