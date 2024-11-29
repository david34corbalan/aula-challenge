package tweets

import (
	"errors"
	"fmt"
	"uala/pkg/common"
	tweets "uala/pkg/tweets/domain"
)

func (s *Service) Delete(id int) (tweet *tweets.Tweets, err error) {

	if id == 0 || id < 0 {
		return nil, errors.New("id invalid")
	}

	tweet, err = s.Repo.Delete(id)
	if err != nil {

		if errors.Is(err, common.ErrNotFound) {
			return nil, common.NewAppError(
				common.ErrCodeNotFound,
				fmt.Sprintf("tweet with id '%d' not found", id))
		}

		return nil, fmt.Errorf("unexpected error getting tweet: %w", err)
	}

	return
}
