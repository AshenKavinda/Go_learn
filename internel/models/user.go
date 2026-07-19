package models

import "time"

type User struct {
	ID        int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`

	Posts []Post

	// Users who follow me
	Followers []User `gorm:"many2many:follows;joinForeignKey:FollowingID;joinReferences:FollowerID"`

	// Users I follow
	Following []User `gorm:"many2many:follows;joinForeignKey:FollowerID;joinReferences:FollowingID"`
}
