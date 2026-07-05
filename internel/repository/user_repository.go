package repository

import (
	"context"

	"github.com/ashenkavinda/go_social_app/internel/models"
	repositoryInterfaces "github.com/ashenkavinda/go_social_app/internel/repository/interfaces"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) repositoryInterfaces.UserRepository {
	return &UserRepository{DB: db}
}

func (s *UserRepository) Create(ctx context.Context, user *models.User) error {
	return nil
}
