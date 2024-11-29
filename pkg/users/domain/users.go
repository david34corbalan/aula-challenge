package users

import (
	"fmt"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model `json:"-"`
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type UserUpdate struct {
	Name     string `json:"name" binding:"required"`
	LastName string `json:"last_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

type UserCreate struct {
	Name     string `json:"name" binding:"required"`
	LastName string `json:"last_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

func SeedUsers() []Users {
	var users []Users
	for i := 1; i <= 5; i++ {
		users = append(users, Users{
			Name:      "name " + strconv.Itoa(i),
			LastName:  "last_name" + strconv.Itoa(i),
			Email:     fmt.Sprintf("email%s@email.@com", strconv.Itoa(i)),
			CreatedAt: time.Now().Add(time.Duration(i) * time.Hour).Format("2006-01-02 15:04:05"),
		},
		)
	}
	return users
}
