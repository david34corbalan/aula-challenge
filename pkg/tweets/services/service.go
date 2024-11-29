package tweets

import (
	follow "uala/pkg/follow/ports"
	tweets "uala/pkg/tweets/ports"
	users "uala/pkg/users/ports"
)

var _ tweets.TweetsService = &Service{}

type Service struct {
	Repo       tweets.TweetsRepository
	RepoUser   users.UsersRepository
	RepoFollow follow.FollowRepository
}
