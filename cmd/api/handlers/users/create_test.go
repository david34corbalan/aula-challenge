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

func Test_CreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usersServiceMock := mocks.NewMockUsersService(ctrl)
	h := users.Handler{
		UsersService: usersServiceMock,
	}

	gin.SetMode(gin.TestMode)

	testCases := []tests.TestCaseHandlers{
		{
			Name:   "should create a user successfully",
			Method: "POST",
			Url:    "/api/v1/users",
			Reqbody: `{
				"name": "new_user",
				"last_name": "new_last_name",
				"email": "new@test.com"
			}`,
			Setup: func(mock ...interface{}) {
				mockService := mock[0].(*mocks.MockUsersService)
				userCreate := domain.UserCreate{
					Name:     "new_user",
					LastName: "new_last_name",
					Email:    "new@test.com",
				}
				createdUser := &domain.Users{
					ID:       1,
					Name:     "new_user",
					LastName: "new_last_name",
					Email:    "new@test.com",
				}
				mockService.EXPECT().Create(userCreate).Return(createdUser, nil)
			},
			Assertionfunc: func(subTest *testing.T, w *httptest.ResponseRecorder) {
				assert.Equal(subTest, http.StatusCreated, w.Code)
				var user domain.Users
				err := json.Unmarshal(w.Body.Bytes(), &user)
				assert.NoError(subTest, err)
				assert.Equal(subTest, uint(1), user.ID)
				assert.Equal(subTest, "new_user", user.Name)
				assert.Equal(subTest, "new_last_name", user.LastName)
				assert.Equal(subTest, "new@test.com", user.Email)
			},
			Handler: h.CreateUser,
		},
		{
			Name:   "should return error 422 for invalid request body",
			Method: "POST",
			Url:    "/api/v1/users",
			Reqbody: `{
				"name": "",
				"last_name": "new_last_name",
				"email": "new@test.com"
			}`,
			Setup: func(mock ...interface{}) {},
			Assertionfunc: func(subTest *testing.T, w *httptest.ResponseRecorder) {
				assert.Equal(subTest, http.StatusUnprocessableEntity, w.Code)
			},
			Handler: h.CreateUser,
		},
	}

	for _, tc := range testCases {
		tests.RunTest(t, tc, usersServiceMock)
	}
}
