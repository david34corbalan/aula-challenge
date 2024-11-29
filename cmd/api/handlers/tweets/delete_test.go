package tweets_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"uala/cmd/api/handlers/tweets"
	"uala/mocks"
	"uala/pkg/common"
	domain "uala/pkg/tweets/domain"
	"uala/tests"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_DeleteTweet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tweetsServiceMock := mocks.NewMockTweetsService(ctrl)
	h := tweets.Handler{
		TweetsService: tweetsServiceMock,
	}

	gin.SetMode(gin.TestMode)

	testCases := []tests.TestCaseHandlers{
		{
			Name:   "should delete a tweet successfully",
			Method: "DELETE",
			Url:    "/api/v1/tweets/1",
			Params: gin.Params{
				{Key: "id", Value: "1"},
			},
			Setup: func(mock ...interface{}) {
				mockService := mock[0].(*mocks.MockTweetsService)
				tweet := &domain.Tweets{
					Comment: "my first comment",
					UserID:  1,
				}
				mockService.EXPECT().Delete(1).Return(tweet, nil)
			},
			Assertionfunc: func(subTest *testing.T, w *httptest.ResponseRecorder) {
				assert.Equal(subTest, http.StatusOK, w.Code)
				var tweet domain.Tweets
				err := json.Unmarshal(w.Body.Bytes(), &tweet)
				assert.NoError(subTest, err)
				assert.Equal(subTest, "my first comment", tweet.Comment)
				assert.Equal(subTest, 1, tweet.UserID)
			},
			Handler: h.DeleteTweet,
		},
		{
			Name:   "should return error 500",
			Method: "DELETE",
			Url:    "/api/v1/tweets/1",
			Params: gin.Params{
				{Key: "id", Value: "1"},
			},
			Setup: func(mock ...interface{}) {
				mockService := mock[0].(*mocks.MockTweetsService)
				mockService.EXPECT().Delete(1).Return(nil, common.ErrNotFound)
			},
			Assertionfunc: func(subTest *testing.T, w *httptest.ResponseRecorder) {
				assert.Equal(subTest, http.StatusInternalServerError, w.Code)
			},
			Handler: h.DeleteTweet,
		},
	}

	for _, tc := range testCases {
		tests.RunTest(t, tc, tweetsServiceMock)
	}
}
