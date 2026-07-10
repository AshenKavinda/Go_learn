package repositoryInterfaces

import (
	"context"

	"github.com/ashenkavinda/go_social_app/internel/models"
)

type PostRepository interface {
	Create(ctx context.Context, post *models.Post) (*int64, error)
	GetAll(ctx context.Context) (*[]models.Post, error)
	//GetByID(ctx context.Context, id int64) (*models.Post, error)
	//Update(ctx context.Context, post *models.Post) error
	//Delete(ctx context.Context, id int64) error
}
