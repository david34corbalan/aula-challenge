package users_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"uala/cmd/api/handlers/users"
	"uala/mocks"
	"uala/pkg/common"
	"uala/pkg/common/validators"
	userDomain "uala/pkg/users/domain"
	"uala/tests"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_IndexUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userServiceMock := mocks.NewMockUsersService(ctrl)
	h := users.Handler{
		UsersService: userServiceMock,
	}

	gin.SetMode(gin.TestMode)
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("limit_offset", validators.LimitAndOffset)
	}

	testCases := []tests.TestCaseHandlers{
		{
			Name:   "should return status code 422",
			Method: "GET",
			Url:    "/api/v1/users?offset=0&limit=0",
			Setup: func(mock ...interface{}) {
				// No setup needed for this test case
			},
			Assertionfunc: func(subTest *testing.T, w *httptest.ResponseRecorder) {
				assert.Equal(subTest, http.StatusUnprocessableEntity, w.Code)
			},
			Handler: h.IndexUsers,
		},
		{
			Name:   "should return status code 200",
			Method: "GET",
			Url:    "/api/v1/users?offset=0&limit=10",
			Setup: func(mock ...interface{}) {
				mockService := mock[0].(*mocks.MockUsersService)

				data := []*userDomain.Users{
					{
						Name:     "test",
						LastName: "test",
						Email:    "test@test.com",
					},
				}

				pag := common.NewPaginate(data, 10, 0, 1)
				paginate := pag.Invoke()
				mockService.EXPECT().Index(gomock.Any()).Return(paginate, nil)
			},
			Assertionfunc: func(subTest *testing.T, w *httptest.ResponseRecorder) {
				var paginate *common.Paginate
				json.Unmarshal(w.Body.Bytes(), &paginate)
				assert.Equal(subTest, http.StatusOK, w.Code)
				assert.Equal(subTest, int64(10), paginate.Limit)
				assert.Equal(subTest, 1, len(paginate.Data.([]interface{})))
			},
			Handler: h.IndexUsers,
		},
		{
			Name:   "should return status code 500",
			Method: "GET",
			Url:    "/api/v1/users?offset=0&limit=10",
			Setup: func(mock ...interface{}) {
				mockService := mock[0].(*mocks.MockUsersService)
				mockService.EXPECT().Index(gomock.Any()).Return(nil, common.ErrRetrieve)
			},
			Assertionfunc: func(subTest *testing.T, w *httptest.ResponseRecorder) {
				assert.Equal(subTest, http.StatusInternalServerError, w.Code)
			},
			Handler: h.IndexUsers,
		},
	}

	for _, tc := range testCases {
		tests.RunTest(t, tc, userServiceMock)
	}
}
