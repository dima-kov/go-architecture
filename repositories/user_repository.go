package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/dima-kov/go-architecture/models"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{db: db}
}

func (r *UserRepository) Get(userID uint) (*models.User, error) {
	user := &models.User{}
	return user, r.db.First(user, "id = ?", userID).Error
}