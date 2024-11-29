package users

import (
	"errors"
	"fmt"
	"uala/pkg/common"
	users "uala/pkg/users/domain"
)

func (s *Service) Create(user users.UserCreate) (users *users.Users, err error) {

	users, err = s.Repo.Create(user)
	if err != nil {
		if errors.Is(err, common.ErrCreate) {
			return nil, common.NewAppError(
				common.ErrCodeInternalServer,
				err.Error(),
			)
		}

		return nil, fmt.Errorf("unexpected error getting user: %w", err)
	}

	return users, err
}
