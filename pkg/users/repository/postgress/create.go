package users

import (
	"uala/pkg/common"
	users "uala/pkg/users/domain"
)

func (r *Repository) Create(userCreate users.UserCreate) (user *users.Users, err error) {

	user = &users.Users{
		Name:     userCreate.Name,
		LastName: userCreate.LastName,
		Email:    userCreate.Email,
	}

	err = r.Client.Model(users.Users{}).Create(&user).Error
	if err != nil {
		return nil, common.ErrCreate
	}

	return
}
