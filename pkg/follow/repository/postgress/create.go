package follow

import (
	"time"
	"uala/pkg/common"
	followDomain "uala/pkg/follow/domain"
)

func (r *Repository) Create(followCreate followDomain.FollowUser) (follow *followDomain.Follow, err error) {
	var followExist followDomain.Follow

	follow = &followDomain.Follow{
		UserID:    followCreate.UserID,
		FollowID:  followCreate.FollowID,
		CreatedAt: time.Now().String(),
	}

	r.Client.Unscoped().Where("user_id = ? AND follow_id = ?", follow.UserID, follow.FollowID).First(&followExist)

	if followExist.ID != 0 {
		err = r.Client.Debug().Unscoped().Delete(&followDomain.Follow{}, followExist.ID).Error
		if err != nil {
			return nil, common.ErrCreate
		}

		follow = &followDomain.Follow{
			UserID:   0,
			FollowID: 0,
		}
		return
	}

	err = r.Client.Model(followDomain.Follow{}).Create(&follow).Error
	if err != nil {
		return nil, common.ErrCreate
	}

	return
}
