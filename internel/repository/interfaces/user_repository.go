package repositoryInterfaces

import (
	"context"

	"github.com/ashenkavinda/go_social_app/internel/models"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) (*int64, error)
	GetAll(ctx context.Context) (*[]models.User, error)
	GetByID(ctx context.Context, id int64) (*models.User, error)
	Update(ctx context.Context, user *models.User) (int, error)
	Delete(ctx context.Context, id int64) (int, error)
}
