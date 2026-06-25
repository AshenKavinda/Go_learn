package store

import (
	"context"
	"database/sql"
)

type Storage struct {
	Post interface {
		Create(ctx context.Context) error
	}

	User interface {
		Create(ctx context.Context) error
	}
}

func NewPostgresStorage(db *sql.DB) Storage {
	return Storage{
		Post: &PostStore{db},
		User: &UserStore{db},
	}
}
