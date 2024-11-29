package example

import (
	example "uala/pkg/example/domain"
)

func (r *Repository) Get(id string) (exam *example.Example, err error) {

	exam = &example.Example{
		ID:   id,
		Name: "Example",
	}

	return
}
