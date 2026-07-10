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

func (r *PostRepository) GetByID(ctx context.Context, id int64) (*models.Post, error) {
	post, err := gorm.G[models.Post](r.DB).Where("id = ?", id).First(ctx)
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *PostRepository) Update(ctx context.Context, post *models.Post) (int, error) {
	ra, err := gorm.G[models.Post](r.DB).Where("id = ?", post.ID).Updates(ctx, *post)
	if err != nil {
		return 0, err
	}
	return ra, nil
}

func (r *PostRepository) Delete(ctx context.Context, id int64) (int, error) {
	ra, err := gorm.G[models.Post](r.DB).Where("id = ?", id).Delete(ctx)
	if err != nil {
		return 0, err
	}
	return ra, nil
}
