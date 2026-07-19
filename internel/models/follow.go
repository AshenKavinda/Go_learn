package models

type Follow struct {
	FollowerID  int64 `gorm:"primaryKey"`
	FollowingID int64 `gorm:"primaryKey"`

	Follower  User `gorm:"foreignKey:ID;references:FollowerID"`
	Following User `gorm:"foreignKey:ID;references:FollowingID"`
}
