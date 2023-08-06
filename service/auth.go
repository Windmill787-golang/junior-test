package service

import (
	"crypto/sha1"
	"fmt"
	"os"
	"time"

	"github.com/Windmill787-golang/junior-test/entities"
	"github.com/Windmill787-golang/junior-test/repository"
	"github.com/dgrijalva/jwt-go"
)

type userClaims struct {
	jwt.StandardClaims
	ID int `json:"id"`
}

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

func (s *AuthService) GenerateToken(user entities.User) (string, error) {
	id, err := s.repo.GetUserId(user.Username, s.generatePassword(user.Password))
	if err != nil {
		return "", err
	}

	if id == 0 {
		return "", nil
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &userClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		id,
	})

	return token.SignedString([]byte(os.Getenv("JWT_SIGN_KEY")))
}
