package repositoryInterfaces

import (
	"context"

	"github.com/ashenkavinda/go_social_app/internel/models"
)

type Follower_repository interface {
	Follow(ctx context.Context, follower int64, following int64) error
	Unfollow(ctx context.Context, follower int64, following int64) error
	Feed(ctx context.Context, userID int64) (*[]models.Post, error)
}
