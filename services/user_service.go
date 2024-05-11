package services

import (
	"alexnerotd/apps/crud_user/models"
	"alexnerotd/apps/crud_user/repositories"
)

type UserService struct {
	userRepository *repositories.UserRepository
}

func NewUserService(userRepository *repositories.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (s *UserService) GetUserById(id int) (*models.User, error) {
	return s.userRepository.GetUserById(id)
}
