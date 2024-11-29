package tweets

import (
	"errors"
	"uala/pkg/common"
	tweets "uala/pkg/tweets/domain"

	"gorm.io/gorm"
)

func (r *Repository) Get(id int) (tweet *tweets.Tweets, err error) {

	err = r.Client.Where("id = ?", id).First(&tweet).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrNotFound
		}
	}

	return
}
