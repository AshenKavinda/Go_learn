package store

import (
	"context"
	"database/sql"

	"github.com/ashenkavinda/go_social_app/internel/models"
	"github.com/ashenkavinda/go_social_app/internel/repository"
)

type Storage struct {
	Post interface {
		Create(ctx context.Context, post *models.Post) error
	}

	User interface {
		Create(ctx context.Context, post *models.User) error
	}
}

func NewPostgresStorage(db *sql.DB) Storage {
	return Storage{
		Post: &repository.PostStore{DB: db},
		User: &repository.UserStore{DB: db},
	}
}
