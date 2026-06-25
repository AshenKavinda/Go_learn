package models

import "github.com/lib/pq"

type Post struct {
	ID        int64          `json:"id"`
	Content   string         `json:"content"`
	Title     string         `json:"title"`
	Tags      pq.StringArray `gorm:"type:text[]" json:"tags"`
	UserID    int64          `json:"user_id"`
	CreatedAt string         `json:"created_at"`
	UpdatedAt string         `json:"updated_at"`
}