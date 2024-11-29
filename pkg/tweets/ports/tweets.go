package tweets

import (
	"uala/pkg/common"
	tweets "uala/pkg/tweets/domain"
)

// mockgen -destination=mocks/mock_tweets_port.go -package=mocks --source=pkg/tweets/ports/tweets.go
type TweetsService interface {
	Index(params common.QuerysParamsPaginate) (paginate *common.Paginate, err error)
	Create(tweet tweets.TweetsCreate) (tweets *tweets.Tweets, err error)
	Get(id int) (tweets *tweets.Tweets, err error)
	Timeline(id int, limit int, offset int) (tweets []*tweets.TweetsUser, err error)
	Update(id int, tweet tweets.TweetsUpdate) (tweets *tweets.Tweets, err error)
	Delete(id int) (tweets *tweets.Tweets, err error)
}

type TweetsRepository interface {
	Index(params common.QuerysParamsPaginate) (tweets []*tweets.Tweets, count int, err error)
	Create(user tweets.TweetsCreate) (tweets *tweets.Tweets, err error)
	Get(id int) (tweets *tweets.Tweets, err error)
	Timeline(ids []int, limit int, offset int) (tweets []*tweets.TweetsUser, err error)
	Update(id int, tweet tweets.TweetsUpdate) (tweets *tweets.Tweets, err error)
	Delete(id int) (tweets *tweets.Tweets, err error)
}
