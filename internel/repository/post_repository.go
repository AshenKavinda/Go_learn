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

func (r *PostRepository) Create(ctx context.Context, post *models.Post) (*int64, error) {
	err := gorm.G[models.Post](r.DB).Create(ctx, post) // pass pointer of data to Create

	if err != nil {
		return nil, err
	}

	return &post.ID, nil
}

func (r *PostRepository) GetAll(ctx context.Context) (*[]models.Post, error) {

	posts, err := gorm.G[models.Post](r.DB).Find(ctx)
	if err != nil {
		return nil, err
	}

	return &posts, nil
}
