package tweets

import (
	tweets "uala/pkg/tweets/ports"

	"gorm.io/gorm"
)

var _ tweets.TweetsRepository = &Repository{}

type Repository struct {
	Client *gorm.DB
}
