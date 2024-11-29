package tweets_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"uala/cmd/api/handlers/tweets"
	"uala/mocks"
	domain "uala/pkg/tweets/domain"
	"uala/tests"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_GetTweet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tweetsServiceMock := mocks.NewMockTweetsService(ctrl)
	h := tweets.Handler{
		TweetsService: tweetsServiceMock,
	}

	gin.SetMode(gin.TestMode)

	testCases := []tests.TestCaseHandlers{
		{
			Name:   "should return a user successfully",
			Method: "GET",
			Url:    "/api/v1/tweets/1",
			Setup: func(mock ...interface{}) {
				mockService := mock[0].(*mocks.MockTweetsService)
				tweet := &domain.Tweets{
					Comment: "my first comment",
					UserID:  1,
				}
				mockService.EXPECT().Get(gomock.Any()).Return(tweet, nil)
			},
			Assertionfunc: func(subTest *testing.T, w *httptest.ResponseRecorder) {
				assert.Equal(subTest, http.StatusOK, w.Code)
				var tweet domain.Tweets
				err := json.Unmarshal(w.Body.Bytes(), &tweet)
				assert.NoError(subTest, err)
			},
			Handler: h.GetTweet,
		},
		// {
		// 	Name:   "should return error 500",
		// 	Method: "GET",
		// 	Url:    "/api/v1/tweets/1",
		// 	Setup: func(mock ...interface{}) {
		// 		mockService := mock[0].(*mocks.MockTweetsService)
		// 		mockService.EXPECT().Get(gomock.Any()).Return(nil, common.ErrNotFound)
		// 	},
		// 	Assertionfunc: func(subTest *testing.T, w *httptest.ResponseRecorder) {
		// 		assert.Equal(subTest, http.StatusInternalServerError, w.Code)
		// 	},
		// 	Handler: h.GetUser,
		// },
	}

	for _, tc := range testCases {
		tests.RunTest(t, tc, tweetsServiceMock)
	}
}
