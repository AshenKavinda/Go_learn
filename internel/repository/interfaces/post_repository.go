package repositoryInterfaces

import (
	"context"

	"github.com/ashenkavinda/go_social_app/internel/models"
)

type PostRepository interface {
	Create(ctx context.Context, post *models.Post) error
}
