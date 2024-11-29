package users

import (
	"errors"
	"fmt"
	"log"
	"uala/pkg/common"
	users "uala/pkg/users/domain"
)

func (s *Service) Update(id int, user users.UserUpdate) (users *users.Users, err error) {

	users, err = s.Repo.Update(id, user)
	if err != nil {

		if errors.Is(err, common.ErrNotFound) {
			return nil, common.NewAppError(
				common.ErrCodeNotFound,
				fmt.Sprintf("user with id '%d' not found", id))
		}
		log.Println(err.Error())
		return nil, fmt.Errorf("unexpected error getting user: %w", err)
	}

	return users, nil
}
