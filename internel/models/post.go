package models

import (
	"time"

	"github.com/lib/pq"
)

type Post struct {
	ID        int64          `gorm:"primaryKey;autoIncrement" json:"id"`
	Content   string         `json:"content"`
	Title     string         `json:"title"`
	Tags      pq.StringArray `gorm:"type:text[]" json:"tags"`
	UserID    int64          `json:"user_id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`

	User User `gorm:"foreignKey:UserID"`
}
