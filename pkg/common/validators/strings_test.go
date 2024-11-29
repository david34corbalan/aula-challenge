package validators_test

import (
	"testing"
	"uala/pkg/common/validators"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

type Strings struct {
	Name string `form:"name" json:"name" validate:"strings"`
}

func Test_Strings(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("strings", validators.Strings)

	tests := []struct {
		name   string
		params Strings
		want   bool
	}{
		{
			name:   "Valid params",
			params: Strings{Name: "fdasmjhfgkda 32132 fdas fdsa2"},
			want:   true,
		},
		{
			name:   "Invalid name",
			params: Strings{Name: "kjhfgdak @!#@!#$$%&^<>,./;'[]\\=-0987654321~!@#$%^&*()_+{}|:<>?"},
			want:   false,
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
