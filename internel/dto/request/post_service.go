package request

import "github.com/lib/pq"

type PostRequest struct {
	Content string         `json:"content"`
	Title   string         `json:"title"`
	Tags    pq.StringArray `gorm:"type:text[]" json:"tags"`
	UserID  int64          `json:"user_id"`
}
