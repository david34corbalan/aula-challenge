package example

import (
	example "uala/pkg/example/ports"

	"gorm.io/gorm"
)

var _ example.ExampleRepository = &Repository{}

type Repository struct {
	Client *gorm.DB
}
