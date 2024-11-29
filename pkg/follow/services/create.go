package follow

import (
	"errors"
	"fmt"
	"log"
	"uala/pkg/common"
	follow "uala/pkg/follow/domain"
)

func (s *Service) Create(followCreate follow.FollowUser) (follow *follow.Follow, err error) {

	if followCreate.FollowID == followCreate.UserID {
		return nil, common.NewAppError(
			common.ErrCodeInternalServer,
			fmt.Errorf("follower and following can't be the same user").Error(),
		)
	}

	_, err = s.RepoUser.Get(int(followCreate.FollowID))
	if err != nil {

		if errors.Is(err, common.ErrNotFound) {
			return nil, common.NewAppError(
				common.ErrCodeNotFound,
				fmt.Sprintf("following with id '%d' not found", int(followCreate.FollowID)))
		}
		log.Println(err.Error())
		return nil, fmt.Errorf("unexpected error getting user: %w", err)
	}

	_, err = s.RepoUser.Get(int(followCreate.UserID))
	if err != nil {

		if errors.Is(err, common.ErrNotFound) {
			return nil, common.NewAppError(
				common.ErrCodeNotFound,
				fmt.Sprintf("follower with id '%d' not found", int(followCreate.UserID)))
		}
		log.Println(err.Error())
		return nil, fmt.Errorf("unexpected error getting user: %w", err)
	}

	follow, err = s.Repo.Create(followCreate)
	if err != nil {
		if errors.Is(err, common.ErrCreate) {
			return nil, common.NewAppError(
				common.ErrCodeInternalServer,
				err.Error(),
			)
		}

		return nil, fmt.Errorf("unexpected error creating follow: %w", err)
	}

	return follow, nil
}
