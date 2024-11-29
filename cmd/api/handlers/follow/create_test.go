package follow_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"uala/cmd/api/handlers/follow"
	"uala/mocks"
	"uala/pkg/common"
	domain "uala/pkg/follow/domain"
	"uala/tests"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_CreateTweet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	followServiceMock := mocks.NewMockFollowService(ctrl)
	h := follow.Handler{
		FollowService: followServiceMock,
	}

	gin.SetMode(gin.TestMode)

	testCases := []tests.TestCaseHandlers{
		{
			Name:    "should create a follow successfully",
			Method:  "POST",
			Url:     "/api/v1/follow",
			Reqbody: `{"user_id": 1, "follow_id":2}`,
			Setup: func(mock ...interface{}) {
				mockService := mock[0].(*mocks.MockFollowService)
				expectedTweet := domain.FollowUser{
					UserID:   1,
					FollowID: 2,
				}
				mockService.EXPECT().Create(expectedTweet).Return(&domain.Follow{
					UserID:   1,
					FollowID: 2,
				}, nil)
			},
			Assertionfunc: func(subTest *testing.T, w *httptest.ResponseRecorder) {
				assert.Equal(subTest, http.StatusCreated, w.Code)
				var follow domain.Follow
				err := json.Unmarshal(w.Body.Bytes(), &follow)
				assert.NoError(subTest, err)
			},
			Handler: h.CreateFollow,
		},
		{
			Name:    "should return error for invalid request body",
			Method:  "POST",
			Url:     "/api/v1/follow",
			Reqbody: `{"user_id": 1}`,
			Setup: func(mock ...interface{}) {
				// No setup needed for this test case
			},
			Assertionfunc: func(subTest *testing.T, w *httptest.ResponseRecorder) {
				assert.Equal(subTest, http.StatusUnprocessableEntity, w.Code)
			},
			Handler: h.CreateFollow,
		},
		{
			Name:    "should return error create follow",
			Method:  "POST",
			Url:     "/api/v1/follow",
			Reqbody: `{"user_id": 1, "follow_id":2}`,
			Setup: func(mock ...interface{}) {
				mockService := mock[0].(*mocks.MockFollowService)
				expectedFollow := domain.FollowUser{
					UserID:   1,
					FollowID: 2,
				}
				mockService.EXPECT().Create(expectedFollow).Return(nil, common.ErrCreate)
			},
			Assertionfunc: func(subTest *testing.T, w *httptest.ResponseRecorder) {
				assert.Equal(subTest, http.StatusInternalServerError, w.Code)
			},
			Handler: h.CreateFollow,
		},
	}

	for _, tc := range testCases {
		tests.RunTest(t, tc, followServiceMock)
	}
}
