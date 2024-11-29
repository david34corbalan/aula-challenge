package example

import (
	"errors"
	"fmt"
	"log"
	example "uala/pkg/example/domain"
)

func (s *Service) Get(id string) (example *example.Example, err error) {

	if id == "" {
		return nil, errors.New("id is required")
	}

	example, err = s.Repo.Get(id)
	if err != nil {
		// if errors.Is(err, domain.ErrNotFound) {
		// 	return nil, domain.NewAppError(
		// 		domain.ErrCodeNotFound,
		// 		fmt.Sprintf("league with id '%s' not found", id))
		// }
		// if errors.Is(err, domain.ErrTimeout) {
		// 	return nil, domain.NewAppError(
		// 		domain.ErrCodeTimeout,
		// 		"timeout error, try again later")
		// }
		log.Println(err.Error())
		return nil, fmt.Errorf("unexpected error getting example: %w", err)
	}

	return example, nil
}
