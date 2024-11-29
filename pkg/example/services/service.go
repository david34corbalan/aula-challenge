package example

import example "uala/pkg/example/ports"

// Make sure Service implements the LeagueService interface
// at compile time.
var _ example.ExampleService = &Service{}

type Service struct {
	Repo example.ExampleRepository
}
