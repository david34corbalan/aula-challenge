package users_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"uala/cmd/api/handlers/users"
	"uala/mocks"
	"uala/pkg/common"
	domain "uala/pkg/users/domain"
	"uala/tests"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_DeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usersServiceMock := mocks.NewMockUsersService(ctrl)
	h := users.Handler{
		UsersService: usersServiceMock,
	}

	gin.SetMode(gin.TestMode)

	testCases := []tests.TestCaseHandlers{
		{
			Name:   "should delete a user successfully",
			Method: "DELETE",
			Url:    "/api/v1/users/:id",
			Params: gin.Params{
				{Key: "id", Value: "2"},
			},
			Setup: func(mock ...interface{}) {
				mockService := mock[0].(*mocks.MockUsersService)
				user := &domain.Users{
					ID:       1,
					Name:     "test",
					LastName: "test",
					Email:    "test@test.com",
				}
				mockService.EXPECT().Delete(gomock.Any()).Return(user, nil)
			},
			Assertionfunc: func(subTest *testing.T, w *httptest.ResponseRecorder) {
				assert.Equal(subTest, http.StatusOK, w.Code)
				var user domain.Users
				err := json.Unmarshal(w.Body.Bytes(), &user)
				assert.NoError(subTest, err)
				assert.Equal(subTest, uint(1), user.ID)
			},
			Handler: h.DeleteUser,
		},
		{
			Name:   "should return error 500",
			Method: "DELETE",
			Url:    "/api/v1/users/1",
			Params: gin.Params{
				{Key: "id", Value: "1"},
			},
			Setup: func(mock ...interface{}) {
				mockService := mock[0].(*mocks.MockUsersService)
				mockService.EXPECT().Delete(gomock.Any()).Return(nil, common.ErrNotFound)
			},
			Assertionfunc: func(subTest *testing.T, w *httptest.ResponseRecorder) {
				assert.Equal(subTest, http.StatusInternalServerError, w.Code)
			},
			Handler: h.DeleteUser,
		},
	}

	for _, tc := range testCases {
		tests.RunTest(t, tc, usersServiceMock)
	}
}
