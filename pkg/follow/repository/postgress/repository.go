package follow

import (
	follow "uala/pkg/follow/ports"

	"gorm.io/gorm"
)

var _ follow.FollowRepository = &Repository{}

type Repository struct {
	Client *gorm.DB
}
