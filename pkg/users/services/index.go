package users

import (
	"errors"
	"fmt"
	"uala/pkg/common"
)

func (s *Service) Index(params common.QuerysParamsPaginate) (*common.Paginate, error) {

	users, count, err := s.Repo.Index(params)
	if err != nil {
		if errors.Is(err, common.ErrRetrieve) {
			return nil, common.NewAppError(
				common.ErrCodeInternalServer,
				err.Error())
		}

		return nil, fmt.Errorf("unexpected error getting example: %w", err)
	}

	pagi := common.NewPaginate(users, params.Limit, params.Offset, count)

	return pagi.Invoke(), nil
}
