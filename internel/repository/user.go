package repository

import (
	"context"
	"database/sql"

	"github.com/ashenkavinda/go_social_app/internel/models"
)

type UserStore struct {
	DB *sql.DB
}

func (s *UserStore) Create(ctx context.Context, user *models.User) error {
	return nil
}
