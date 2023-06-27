package services

import (
	"errors"
	"github.com/dysdimas/go-clean-arch-tmpl/models"
)

type UserService struct {
	userModel *models.UserModel
}

func NewUserService(userModel *models.UserModel) *UserService {
	return &UserService{
		userModel: userModel,
	}
}

func (s *UserService) GetUserByUsername(username string) (*models.User, error) {
	userByUsername, err := s.userModel.GetUserByUsername(username)
	if err != nil {
		return nil, errors.New("User not found")
	}

	return userByUsername, err
}

func (s *UserService) DeleteUserByUsername(username string) (string, error) {
	_, err := s.userModel.DeleteUserByUsername(username)
	if err != nil {
		return "", errors.New("Delete user fail")
	}
	return "", err
}

func (s *UserService) FetchAllUser() ([]map[string]interface{}, error) {
	// Get the user from the database
	user, err := s.userModel.FetchAllUser()
	if err != nil {
		return nil, errors.New("User not found")
	}

	return user, err
}
