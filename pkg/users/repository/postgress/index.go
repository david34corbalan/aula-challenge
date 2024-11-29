package users

import (
	"uala/pkg/common"
	users "uala/pkg/users/domain"

	"gorm.io/gorm"
)

func (r *Repository) Index(params common.QuerysParamsPaginate) (usersFind []*users.Users, count int, err error) {

	query := r.Client.Model(users.Users{}).Scopes(Search(params.Search)).Order("last_name DESC")

	var counInt64 int64
	query.Count(&counInt64)
	count = int(counInt64)

	err = query.Limit(params.Limit).Offset(params.Offset).Find(&usersFind).Error
	if err != nil {
		return nil, 0, common.ErrRetrieve
	}

	return
}

func Search(search string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if search != "" {
			db = db.Where("name LIKE ?", "%"+search+"%").Or("email LIKE ?", "%"+search+"%").Or("last_name LIKE ?", "%"+search+"%")
		}
		return db
	}
}
