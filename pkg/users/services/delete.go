package users

import (
	"errors"
	"fmt"
	"uala/pkg/common"
	users "uala/pkg/users/domain"
)

func (s *Service) Delete(id int) (users *users.Users, err error) {

	users, err = s.Repo.Delete(id)
	if err != nil {

		if errors.Is(err, common.ErrNotFound) {
			return nil, common.NewAppError(
				common.ErrCodeNotFound,
				fmt.Sprintf("user with id '%d' not found", id))
		}

		return nil, fmt.Errorf("unexpected error getting user: %w", err)
	}

	return users, nil
}
