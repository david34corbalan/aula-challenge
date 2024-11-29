package tweets

import (
	tweets "uala/pkg/tweets/domain"

	"gorm.io/gorm"
)

func (r *Repository) Timeline(ids []int, limit int, offset int) (tweetsUsers []*tweets.TweetsUser, err error) {
	batchSize := 100
	err = r.Client.
		Model(&tweets.Tweets{}).
		Select("tweets.id, tweets.comment, tweets.user_id, tweets.created_at, tweets.updated_at, users.id as user_id, users.last_name as user_last_name,users.name as user_name, users.email as user_email").
		Joins("JOIN users ON users.id = tweets.user_id").
		Where("tweets.user_id IN (?)", ids).
		Where("tweets.deleted_at IS NULL").
		Order("tweets.created_at DESC").
		Offset(offset).
		Limit(limit).
		FindInBatches(&tweetsUsers, batchSize, func(tx *gorm.DB, batch int) error {
			return nil
		}).Error

	return
}
