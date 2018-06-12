package services

import (
	"github.com/dima-kov/go-architecture/models"
	"github.com/dima-kov/go-architecture/repositories"
)

type PostService struct {
	repository repositories.PostRepository
}

func NewPostService(repository repositories.PostRepository) PostService {
	return PostService{repository: repository}
}

func (s *PostService) GetPost(postID uint) (*models.Post, error) {
	return s.repository.Get(postID)
}
