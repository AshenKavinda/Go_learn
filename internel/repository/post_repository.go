package repository

import (
	"context"

	"github.com/ashenkavinda/go_social_app/internel/models"
	repositoryInterfaces "github.com/ashenkavinda/go_social_app/internel/repository/interfaces"
	"gorm.io/gorm"
)

type PostRepository struct {
	DB *gorm.DB
}

func NewPostRepository(db *gorm.DB) repositoryInterfaces.PostRepository {
	return &PostRepository{DB: db}
}

func (r *PostRepository) Create(ctx context.Context, post *models.Post) error {
	result := r.DB.WithContext(ctx).Create(&post) // pass pointer of data to Create

	if result.Error != nil {
		return result.Error
	}

	return nil
}
