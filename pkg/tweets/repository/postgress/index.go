package tweets

import (
	"uala/pkg/common"
	tweets "uala/pkg/tweets/domain"

	"gorm.io/gorm"
)

func (r *Repository) Index(params common.QuerysParamsPaginate) (tweet []*tweets.Tweets, count int, err error) {

	query := r.Client.Model(tweets.Tweets{}).Scopes(Search(params.Search)).Order("id DESC")

	var counInt64 int64
	query.Count(&counInt64)
	count = int(counInt64)

	err = query.Limit(params.Limit).Offset(params.Offset).Find(&tweet).Error
	if err != nil {
		return nil, 0, common.ErrRetrieve
	}

	return
}

func Search(search string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if search != "" {
			db = db.Where("comment LIKE ?", "%"+search+"%")
		}
		return db
	}
}
