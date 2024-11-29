package users

import (
	"errors"
	"uala/pkg/common"
	users "uala/pkg/users/domain"

	"gorm.io/gorm"
)

func (r *Repository) Update(id int, userUpdate users.UserUpdate) (user *users.Users, err error) {
	var userExist users.Users
	err = r.Client.Model(users.Users{}).Where("id = ?", id).First(&userExist).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrNotFound
		}
	}

	user = &users.Users{
		Name:     userUpdate.Name,
		LastName: userUpdate.LastName,
		Email:    userUpdate.Email,
	}

	err = r.Client.Model(users.Users{}).Where("id = ?", id).Updates(user).Error

	return
}
