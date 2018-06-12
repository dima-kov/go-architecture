package services

import (
	"github.com/dima-kov/go-architecture/repositories"
	"github.com/dima-kov/go-architecture/models"
)

type UserService struct {
	repository repositories.UserRepository
}

func NewUserService(repository repositories.UserRepository) UserService {
	return UserService{repository:repository}
}

func (s *UserService) GetUser(userID uint) (*models.User, error) {
	return s.repository.Get(userID)
}
