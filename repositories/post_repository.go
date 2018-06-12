package repositories

import (
	"github.com/dima-kov/go-architecture/models"
	"github.com/jinzhu/gorm"
)

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return PostRepository{db: db}
}

func (r *PostRepository) Get(postID uint) (*models.Post, error) {
	post := &models.Post{}
	return post, r.db.First(post, "id = ?", postID).Error
}
