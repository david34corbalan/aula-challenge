package users

import (
	"errors"
	"uala/pkg/common"
	users "uala/pkg/users/domain"

	"gorm.io/gorm"
)

func (r *Repository) Get(id int) (user *users.Users, err error) {

	err = r.Client.Where("id = ?", id).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrNotFound
		}
	}

	return
}
