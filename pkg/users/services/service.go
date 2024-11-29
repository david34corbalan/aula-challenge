package users

import users "uala/pkg/users/ports"

var _ users.UsersService = &Service{}

type Service struct {
	Repo users.UsersRepository
}
