package repository

import (
	"context"
	"database/sql"

	"github.com/ashenkavinda/go_social_app/internel/models"
)

type PostStore struct {
	DB *sql.DB
}

func (s *PostStore) Create(ctx context.Context, post *models.Post) error {
	return nil
}
