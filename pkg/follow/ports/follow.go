package follow

import follow "uala/pkg/follow/domain"

// mockgen -destination=mocks/mock_follow_port.go -package=mocks --source=pkg/follow/ports/follow.go
type FollowService interface {
	Create(followCreate follow.FollowUser) (follow *follow.Follow, err error)
	// GetFollowers(id int) (ids []*int, err error)
}

type FollowRepository interface {
	Create(followCreate follow.FollowUser) (follow *follow.Follow, err error)
	GetFollowers(id int) (ids []int, err error)
}
