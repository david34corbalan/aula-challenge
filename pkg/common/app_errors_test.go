package common_test

import (
	"fmt"
	"testing"
	"uala/pkg/common"

	"github.com/stretchr/testify/assert"
)

func Test_App_Errors(t *testing.T) {

	testTable := map[string]struct {
		setup         func() common.AppError
		nameTest      string
		assertionFunc func(subTest *testing.T, errApp common.AppError)
	}{
		"error ErrCreate type": {
			setup: func() common.AppError {
				return common.NewAppError(500, common.ErrCreate.Error())
			},
			nameTest: "ErrCreate",
			assertionFunc: func(subTest *testing.T, errApp common.AppError) {
				assert.Equal(subTest, 500, errApp.Code)
				assert.Equal(subTest, "error creating record", errApp.Msg)
			},
		},
		"error ErridNotValid type": {
			setup: func() common.AppError {
				return common.NewAppError(500, common.ErridNotValid.Error())
			},
			nameTest: "ErridNotValid",
			assertionFunc: func(subTest *testing.T, errApp common.AppError) {
				assert.Equal(subTest, 500, errApp.Code)
				assert.Equal(subTest, "id is not valid", errApp.Msg)
			},
		},
		"error ErrNotFound type": {
			setup: func() common.AppError {
				return common.NewAppError(404, common.ErrNotFound.Error())
			},
			nameTest: "ErrNotFound",
			assertionFunc: func(subTest *testing.T, errApp common.AppError) {
				assert.Equal(subTest, 404, errApp.Code)
				assert.Equal(subTest, "record not found error", errApp.Msg)
			},
		},
		"error ErrTimeout type": {
			setup: func() common.AppError {
				return common.NewAppError(500, common.ErrTimeout.Error())
			},
			nameTest: "ErrTimeout",
			assertionFunc: func(subTest *testing.T, errApp common.AppError) {
				assert.Equal(subTest, 500, errApp.Code)
				assert.Equal(subTest, "timeout error", errApp.Msg)
			},
		},
		"valdiate interface Error() ": {
			setup: func() common.AppError {
				return common.NewAppError(500, common.ErrTimeout.Error())
			},
			nameTest: "ErrTimeout",
			assertionFunc: func(subTest *testing.T, errApp common.AppError) {
				assert.Equal(subTest, fmt.Sprintf("%d: %s", errApp.Code, errApp.Msg), errApp.Error())
			},
		},
	}

	for name, test := range testTable {
		t.Run(name, func(subTest *testing.T) {
			test.setup()
			test.nameTest = name
			test.assertionFunc(subTest, test.setup())
		})
	}
}
