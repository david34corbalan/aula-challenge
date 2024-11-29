package users_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"uala/cmd/api/handlers/users"
	"uala/mocks"
	domain "uala/pkg/users/domain"
	"uala/tests"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_UpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usersServiceMock := mocks.NewMockUsersService(ctrl)
	h := users.Handler{
		UsersService: usersServiceMock,
	}

	gin.SetMode(gin.TestMode)

	testCases := []tests.TestCaseHandlers{
		{
			Name:   "should update a user successfully",
			Method: "PUT",
			Url:    "/api/v1/users/1",
			Params: gin.Params{
				{Key: "id", Value: "1"},
			},
			Reqbody: `{
				"name": "updated_name",
				"last_name": "updated_last_name",
				"email": "updated@test.com"
			}`,
			Setup: func(mock ...interface{}) {
				mockService := mock[0].(*mocks.MockUsersService)
				userUpdate := domain.UserUpdate{
					Name:     "updated_name",
					LastName: "updated_last_name",
					Email:    "updated@test.com",
				}
				updatedUser := &domain.Users{
					ID:       1,
					Name:     "updated_name",
					LastName: "updated_last_name",
					Email:    "updated@test.com",
				}
				mockService.EXPECT().Update(1, userUpdate).Return(updatedUser, nil)
			},
			Assertionfunc: func(subTest *testing.T, w *httptest.ResponseRecorder) {
				assert.Equal(subTest, http.StatusOK, w.Code)
				var user domain.Users
				err := json.Unmarshal(w.Body.Bytes(), &user)
				assert.NoError(subTest, err)
				assert.Equal(subTest, uint(1), user.ID)
				assert.Equal(subTest, "updated_name", user.Name)
				assert.Equal(subTest, "updated_last_name", user.LastName)
				assert.Equal(subTest, "updated@test.com", user.Email)
			},
			Handler: h.UpdateUser,
		},
		{
			Name:   "should return error 422 for invalid request body",
			Method: "PUT",
			Url:    "/api/v1/users/1",
			Params: gin.Params{
				{Key: "id", Value: "1"},
			},
			Reqbody: `{
				"name": 123,
				"last_name": "updated_last_name",
				"email": "updated@test.com"
			}`,
			Setup: func(mock ...interface{}) {},
			Assertionfunc: func(subTest *testing.T, w *httptest.ResponseRecorder) {
				assert.Equal(subTest, http.StatusUnprocessableEntity, w.Code)
			},
			Handler: h.UpdateUser,
		},
	}

	for _, tc := range testCases {
		tests.RunTest(t, tc, usersServiceMock)
	}
}
