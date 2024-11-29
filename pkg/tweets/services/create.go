package tweets

import (
	"errors"
	"fmt"
	"uala/pkg/common"
	tweets "uala/pkg/tweets/domain"
)

func (s *Service) Create(tweet tweets.TweetsCreate) (tweets *tweets.Tweets, err error) {

	_, err = s.RepoUser.Get(tweet.UserID)
	if err != nil {
		if errors.Is(err, common.ErrNotFound) {
			return nil, common.NewAppError(
				common.ErrCodeNotFound,
				fmt.Sprintf("user with id '%d' not found", tweet.UserID),
			)
		}

		return nil, fmt.Errorf("unexpected error getting user: %w", err)
	}

	tweets, err = s.Repo.Create(tweet)
	if err != nil {
		if errors.Is(err, common.ErrCreate) {
			return nil, common.NewAppError(
				common.ErrCodeInternalServer,
				err.Error(),
			)
		}

		return nil, fmt.Errorf("unexpected error getting tweet: %w", err)
	}

	return tweets, err
}
