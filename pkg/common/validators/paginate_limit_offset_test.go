package validators_test

import (
	"testing"
	"uala/pkg/common/validators"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

type QuerysParamsPaginateTest struct {
	Offset int    `form:"offset" json:"offset" validate:"limitAndOffset"`
	Limit  int    `form:"limit" json:"limit" validate:"limitAndOffset"`
	Search string `form:"search" json:"search,omitempty"`
}

func TestLimitAndOffset(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("limitAndOffset", validators.LimitAndOffset)

	tests := []struct {
		name   string
		params QuerysParamsPaginateTest
		want   bool
	}{
		{
			name:   "Valid params",
			params: QuerysParamsPaginateTest{Limit: 10, Offset: 5},
			want:   true,
		},
		{
			name:   "Invalid Limit",
			params: QuerysParamsPaginateTest{Limit: 0, Offset: 5},
			want:   false,
		},
		{
			name:   "Invalid Offset",
			params: QuerysParamsPaginateTest{Limit: 10, Offset: -1},
			want:   false,
		},
		{
			name:   "Empty Offset",
			params: QuerysParamsPaginateTest{Limit: 10, Offset: 0},
			want:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validate.Struct(tt.params)
			if tt.want {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}
