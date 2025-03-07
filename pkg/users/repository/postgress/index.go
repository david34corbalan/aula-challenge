package users

import (
	"uala/pkg/common"
	users "uala/pkg/users/domain"
)

func (r *Repository) Index(params common.QuerysParamsPaginate) (usersFind []*users.Users, count int, err error) {

	query := r.Client.Model(users.Users{}).Order("id DESC")

	var counInt64 int64
	query.Count(&counInt64)
	count = int(counInt64)

	err = query.Limit(params.Limit).Offset(params.Offset).Find(&usersFind).Error
	if err != nil {
		return nil, 0, common.ErrRetrieve
	}

	return
}
