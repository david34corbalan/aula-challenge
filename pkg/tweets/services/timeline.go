package tweets

import (
	"errors"
	"fmt"
	"log"
	"uala/pkg/common"
	tweets "uala/pkg/tweets/domain"
)

func (s *Service) Timeline(id int, limit int, offset int) (tweetsUsers []*tweets.TweetsUser, err error) {

	// que exista el usuario
	_, err = s.RepoUser.Get(id)
	if err != nil {
		if errors.Is(err, common.ErrNotFound) {
			return nil, common.NewAppError(
				common.ErrCodeNotFound,
				fmt.Sprintf("user with id '%d' not found", id))
		}
		log.Println(err.Error())
		return nil, fmt.Errorf("unexpected error getting user: %w", err)
	}

	// tener los idsF de los followers
	idsF, err := s.RepoFollow.GetFollowers(id)
	if err != nil {

		if errors.Is(err, common.ErrNotFound) {
			return nil, common.NewAppError(
				common.ErrCodeNotFound,
				fmt.Sprintf("user with id '%d' not found", id))
		}
		log.Println(err.Error())
		return nil, fmt.Errorf("unexpected error getting user: %w", err)
	}

	tweetsUsers, err = s.Repo.Timeline(idsF, limit, offset)
	if err != nil {

		return nil, fmt.Errorf("unexpected error getting tweets: %w", err)
	}

	return
}
