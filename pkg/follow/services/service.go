package follow

import (
	follow "uala/pkg/follow/ports"
	users "uala/pkg/users/ports"
)

var _ follow.FollowService = &Service{}

type Service struct {
	Repo     follow.FollowRepository
	RepoUser users.UsersRepository
}
