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

func (s *UserRepository) Create(ctx context.Context, user *models.User) (*int64, error) {
	err := gorm.G[models.User](s.DB).Create(ctx, user)

	if err != nil {
		return nil, err
	}

	return &user.ID, nil
}

func (s *UserRepository) GetAll(ctx context.Context) (*[]models.User, error) {
	posts, err := gorm.G[models.User](s.DB).Find(ctx)
	if err != nil {
		return nil, err
	}

	return &posts, nil
}

func (s *UserRepository) GetByID(ctx context.Context, id int64) (*models.User, error) {
	user, err := gorm.G[models.User](s.DB).Where("id = ?", id).First(ctx)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *UserRepository) Update(ctx context.Context, user *models.User) (*int, error) {
	ra, err := gorm.G[models.User](s.DB).Where("id = ?", user.ID).Updates(ctx, *user)
	if err != nil {
		return nil, err
	}

	return &ra, nil
}

func (s *UserRepository) Delete(ctx context.Context, id int64) (*int, error) {
	ra, err := gorm.G[models.User](s.DB).Where("id = ?", id).Delete(ctx)
	if err != nil {
		return nil, err
	}

	return &ra, nil

}
