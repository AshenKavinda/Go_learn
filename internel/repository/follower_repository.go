package repository

import (
	"context"

	"github.com/ashenkavinda/go_social_app/internel/models"
	repositoryInterfaces "github.com/ashenkavinda/go_social_app/internel/repository/interfaces"
	"gorm.io/gorm"
)

type follower_repository struct {
	db *gorm.DB
}

func NewFolloerRepository(db *gorm.DB) repositoryInterfaces.Follower_repository {
	return &follower_repository{db: db}
}

func (r *follower_repository) Follow(ctx context.Context, follower int64, following int64) error {
	follow := &models.Follow{
		FollowerID:  follower,
		FollowingID: following,
	}

	return gorm.G[models.Follow](r.db).Create(ctx, follow)
}

func (r *follower_repository) Unfollow(ctx context.Context, follower int64, following int64) error {
	_, err := gorm.G[models.Follow](r.db).
		Where("follower_id = ? AND following_id = ?", follower, following).
		Delete(ctx)
	return err
}

func (r *follower_repository) Feed(ctx context.Context, userID int64) (*[]models.Post, error) {
	var posts []models.Post
	err := r.db.WithContext(ctx).
		Joins("INNER JOIN follows f ON f.following_id = posts.user_id").
		Where("f.follower_id = ?", userID).
		Order("posts.created_at DESC").
		Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return &posts, nil
}
