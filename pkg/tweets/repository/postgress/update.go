package tweets

import (
	"errors"
	"uala/pkg/common"
	tweets "uala/pkg/tweets/domain"

	"gorm.io/gorm"
)

func (r *Repository) Update(id int, tweetUpdate tweets.TweetsUpdate) (tweet *tweets.Tweets, err error) {
	var userExist tweets.Tweets
	err = r.Client.Model(tweets.Tweets{}).Where("id = ?", id).First(&userExist).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrNotFound
		}
	}

	tweet = &tweets.Tweets{
		Comment: tweetUpdate.Comment,
		UserID:  userExist.UserID,
		ID:      userExist.ID,
	}

	err = r.Client.Model(&tweets.Tweets{}).Where("id = ?", id).Updates(&tweets.Tweets{Comment: tweet.Comment}).Error

	return
}
