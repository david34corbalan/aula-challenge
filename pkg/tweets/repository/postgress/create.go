package tweets

import (
	"uala/pkg/common"
	tweets "uala/pkg/tweets/domain"
)

func (r *Repository) Create(tweetCreate tweets.TweetsCreate) (tweet *tweets.Tweets, err error) {

	tweet = &tweets.Tweets{
		Comment: tweetCreate.Comment,
		UserID:  tweetCreate.UserID,
	}

	err = r.Client.Model(tweets.Tweets{}).Create(&tweet).Error
	if err != nil {
		return nil, common.ErrCreate
	}

	return
}
