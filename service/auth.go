package service

import (
	"crypto/sha1"
	"fmt"
	"os"

	"github.com/Windmill787-golang/junior-test/entities"
	"github.com/Windmill787-golang/junior-test/repository"
)

type AuthService struct {
	repo repository.Auth
}

func NewAuthService(repo repository.Auth) *AuthService {
	return &AuthService{repo}
}

func (s *AuthService) CreateUser(user entities.User) (int, error) {
	user.PasswordHash = s.generatePassword(user.Password)

	return s.repo.CreateUser(user)
}

func (s *AuthService) generatePassword(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(os.Getenv("AUTH_SALT"))))
}
