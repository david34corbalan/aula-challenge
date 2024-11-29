package tweets

import (
	"strconv"
	"time"

	"gorm.io/gorm"
)

type Tweets struct {
	gorm.Model `json:"-"`
	ID         uint   `json:"id"`
	Comment    string `json:"comment"`
	UserID     int    `json:"user_id"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type TweetsUpdate struct {
	Comment string `json:"comment" binding:"required,max=280"`
	UserID  int    `json:"user_id" binding:"required"`
}

type TweetsCreate struct {
	Comment string `json:"comment" binding:"required,max=280"`
	UserID  int    `json:"user_id" binding:"required"`
}

type TweetsUser struct {
	gorm.Model   `json:"-"`
	ID           uint   `json:"id"`
	Comment      string `json:"comment"`
	UserID       int    `json:"user_id"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	UserName     string `json:"user_name"`
	UserLastName string `json:"user_last_name"`
	UserEmail    string `json:"user_email"`
}

func SeedTweets() []Tweets {
	var users []Tweets
	id := 1
	for i := 1; i < 10000; i++ {
		users = append(users, Tweets{
			Comment:   "comment " + strconv.Itoa(i),
			UserID:    id,
			CreatedAt: time.Now().Add(time.Duration(i) * time.Hour).Format("2006-01-02 15:04:05"),
		})
		if id == 5 {
			id = 1
		} else {
			id++
		}
	}
	return users
}
