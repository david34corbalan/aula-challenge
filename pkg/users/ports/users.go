package users

import (
	"uala/pkg/common"
	users "uala/pkg/users/domain"
)

// mockgen -destination=mocks/mock_users_port.go -package=mocks --source=pkg/users/ports/users.go
type UsersService interface {
	Index(params common.QuerysParamsPaginate) (paginate *common.Paginate, err error)
	Create(user users.UserCreate) (users *users.Users, err error)
	Get(id int) (users *users.Users, err error)
	Update(id int, user users.UserUpdate) (users *users.Users, err error)
	Delete(id int) (users *users.Users, err error)
}

type UsersRepository interface {
	Index(params common.QuerysParamsPaginate) (users []*users.Users, count int, err error)
	Create(user users.UserCreate) (users *users.Users, err error)
	Get(id int) (users *users.Users, err error)
	Update(id int, user users.UserUpdate) (users *users.Users, err error)
	Delete(id int) (users *users.Users, err error)
}
