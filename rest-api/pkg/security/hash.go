package security

import (
	"golang.org/x/crypto/bcrypt"
)

type (
	HashService interface {
		HashPassword(password string) (string, error)
		CheckPasswordHash(password, hash string) bool
	}
	hashService struct{}
)

func NewHashService() HashService {
	return &hashService{}
}

func (h *hashService) HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}
func (h *hashService) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
