package services

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dysdimas/go-clean-arch-tmpl/models"
)

type AuthService struct {
	userModel *models.UserModel
}

func NewAuthService(userModel *models.UserModel) *AuthService {
	return &AuthService{
		userModel: userModel,
	}
}

func (s *AuthService) RegisterUser(username, password, role string) error {
	// Check if the user already exists
	_, err := s.userModel.GetUserByUsername(username)
	if err == nil {
		return errors.New("user already exists")
	}

	// Create a new user
	user := &models.User{
		Username: username,
		Password: password,
		Role:     role,
	}

	err = s.userModel.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (s *AuthService) AuthenticateUser(username, password string) (string, error) {
	// Get the user from the database
	user, err := s.userModel.GetUserByUsernameAndPassword(username, password)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// Generate a JWT token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["role"] = user.Role
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expiration time

	tokenString, err := token.SignedString([]byte("your-secret-key"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
