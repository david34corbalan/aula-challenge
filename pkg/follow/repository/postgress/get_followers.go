package follow

import (
	followDomain "uala/pkg/follow/domain"
)

func (r *Repository) GetFollowers(id int) (ids []int, err error) {

	err = r.Client.Debug().Model(followDomain.Follow{}).Unscoped().Where("user_id = ?", id).Pluck("follow_id", &ids).Error
	return
}
