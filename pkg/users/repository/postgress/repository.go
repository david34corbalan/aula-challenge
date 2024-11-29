package users

import (
	users "uala/pkg/users/ports"

	"gorm.io/gorm"
)

var _ users.UsersRepository = &Repository{}

type Repository struct {
	Client *gorm.DB
}
