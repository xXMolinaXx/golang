package user

import (
	"fmt"

	"github.com/xXMolinaXx/golang/internal/domain"
	"github.com/xXMolinaXx/golang/pkg/security"
)

type Service interface {
	CreateUser(name, email, password string) error
	ReadUser(id string) (*domain.User, error)
	ReadAllUsers() ([]domain.User, error)
	UpdateUser(user *domain.User, requestBody CreateUserRequest) error
	DeleteUser(id string) error
	Login(email, password string) (*domain.User, string, string, error)
}

type UserService struct {
	db          Repository
	hashService security.HashService
	jwtService  security.Ijwt
}

func NewUserService(db Repository, hashService security.HashService, jwtService security.Ijwt) *UserService {
	return &UserService{db: db, hashService: hashService, jwtService: jwtService} // crea la variable y retorna un puntero
}

func (s *UserService) CreateUser(name, email, password string) error {
	if password == "" {
		return fmt.Errorf("password is required")
	}

	user := domain.User{
		Fullname: name,
		Email:    email,
		Password: password,
	}
	return s.db.CreateUser(&user)
}

func (s *UserService) ReadUser(id string) (*domain.User, error) {
	return s.db.ReadUser(id)
}
func (s *UserService) ReadAllUsers() ([]domain.User, error) {
	return s.db.ReadAllUsers()
}
func (s *UserService) UpdateUser(user *domain.User, requestBody CreateUserRequest) error {
	user.Fullname = requestBody.Name
	user.Email = requestBody.Email

	return s.db.UpdateUser(user)
}
func (s *UserService) DeleteUser(id string) error {
	return s.db.DeleteUser(id)
}
func (s *UserService) Login(email, password string) (*domain.User, string, string, error) {

	user, err := s.db.Login(email, password)
	if err != nil {
		return nil, "", "", err
	}
	if !s.hashService.CheckPasswordHash(password, user.Password) {
		return nil, "", "", fmt.Errorf("invalid email or password")
	}
	token, err := s.jwtService.GenerateToken(user.Email, user.Fullname, false)
	if err != nil {
		return nil, "", "", err
	}
	refreshToken, err := s.jwtService.GenerateToken(user.Email, user.Fullname, true)
	if err != nil {
		return nil, "", "", err
	}
	return user, token, refreshToken, nil
}
