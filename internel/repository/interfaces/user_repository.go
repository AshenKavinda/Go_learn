package repositoryInterfaces

import (
	"context"

	"github.com/ashenkavinda/go_social_app/internel/models"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
}
