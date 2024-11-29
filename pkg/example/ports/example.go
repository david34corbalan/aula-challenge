package example

import example "uala/pkg/example/domain"

type ExampleService interface {
	Get(id string) (example *example.Example, err error)
}

type ExampleRepository interface {
	Get(id string) (example *example.Example, err error)
}
