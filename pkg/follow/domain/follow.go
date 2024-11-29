package follow

import (
	"time"

	"gorm.io/gorm"
)

type Follow struct {
	gorm.Model `json:"-"`
	ID         uint   `json:"id"`
	UserID     uint   `json:"user_id"`
	FollowID   uint   `json:"follow_id"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at,omitempty"`
}

type FollowUser struct {
	UserID   uint `json:"user_id" binding:"required"`
	FollowID uint `json:"follow_id" binding:"required"`
}

func SeedFollowers() []Follow {
	var followers []Follow
	for i := 1; i <= 5; i++ {
		followers = append(followers, Follow{
			UserID:    1,
			FollowID:  uint(i),
			CreatedAt: time.Now().Add(time.Duration(i) * time.Hour).Format("2006-01-02 15:04:05"),
		},
		)
	}
	return followers
}
