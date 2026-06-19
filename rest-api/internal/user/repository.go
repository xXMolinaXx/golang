package user

import (
	"github.com/xXMolinaXx/golang/internal/domain"

	"gorm.io/gorm"
)

type Repository interface {
	CreateUser(user *domain.User) error
	ReadUser(id string) (*domain.User, error)
	ReadAllUsers() ([]domain.User, error)
	UpdateUser(user *domain.User) error
	DeleteUser(id string) error
	Login(email, password string) (*domain.User, error)
}

type repo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) Repository {
	return &repo{db: db}
}

func (r *repo) CreateUser(user *domain.User) error {
	err := r.db.Create(user).Error
	return err
}
func (r *repo) ReadUser(id string) (*domain.User, error) {
	var user domain.User
	err := r.db.Select("id", "fullname", "email", "created_at", "updated_at").First(&user, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil

}
func (r *repo) ReadAllUsers() ([]domain.User, error) {
	var users []domain.User
	err := r.db.Select("id", "fullname", "email", "created_at", "updated_at").Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *repo) UpdateUser(user *domain.User) error {
	err := r.db.Save(user).Error
	return err
}

func (r *repo) DeleteUser(id string) error {
	err := r.db.Delete(&domain.User{}, "id = ?", id).Error
	return err
}

func (r *repo) Login(email, password string) (*domain.User, error) {
	var user domain.User
	err := r.db.Select("id", "fullname", "email", "password").Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
