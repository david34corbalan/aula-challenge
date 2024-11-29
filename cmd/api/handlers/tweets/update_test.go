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

func Test_UpdateTweet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tweetsServiceMock := mocks.NewMockTweetsService(ctrl)
	h := tweets.Handler{
		TweetsService: tweetsServiceMock,
	}

	gin.SetMode(gin.TestMode)

	testCases := []tests.TestCaseHandlers{
		{
			Name:   "should update a tweet successfully",
			Method: "PUT",
			Url:    "/api/v1/tweets/1",
			Params: gin.Params{
				{Key: "id", Value: "1"},
			},
			Reqbody: `{"comment": "my first comment", "user_id": 1}`,
			Setup: func(mock ...interface{}) {
				mockService := mock[0].(*mocks.MockTweetsService)
				expectedTweet := domain.TweetsUpdate{
					Comment: "my first comment",
					UserID:  1,
				}
				tweetReturn := &domain.Tweets{
					ID:      1,
					Comment: "my first comment",
					UserID:  1,
				}
				mockService.EXPECT().Update(1, expectedTweet).Return(tweetReturn, nil)
			},
			Assertionfunc: func(subTest *testing.T, w *httptest.ResponseRecorder) {
				assert.Equal(subTest, http.StatusCreated, w.Code)
				var tweet domain.Tweets
				err := json.Unmarshal(w.Body.Bytes(), &tweet)
				assert.NoError(subTest, err)
				assert.Equal(subTest, uint(1), tweet.ID)
				assert.Equal(subTest, "my first comment", tweet.Comment)
				assert.Equal(subTest, 1, tweet.UserID)
			},
			Handler: h.UpdateTweet,
		},
		{
			Name:   "should return error for invalid request",
			Method: "PUT",
			Url:    "/api/v1/tweets/1",
			Params: gin.Params{
				{Key: "id", Value: "1"},
			},
			Reqbody: `{"comment": "my first comment"}`,
			Setup: func(mock ...interface{}) {
				// No setup needed for this test case
			},
			Assertionfunc: func(subTest *testing.T, w *httptest.ResponseRecorder) {
				assert.Equal(subTest, http.StatusUnprocessableEntity, w.Code)
			},
			Handler: h.UpdateTweet,
		},
		{
			Name:   "should return error update tweet",
			Method: "PUT",
			Url:    "/api/v1/tweets/1",
			Params: gin.Params{
				{Key: "id", Value: "1"},
			},
			Reqbody: `{"comment": "my first comment", "user_id": 1}`,
			Setup: func(mock ...interface{}) {
				mockService := mock[0].(*mocks.MockTweetsService)
				expectedTweet := domain.TweetsUpdate{
					Comment: "my first comment",
					UserID:  1,
				}
				mockService.EXPECT().Update(1, expectedTweet).Return(nil, common.ErrNotFound)
			},
			Assertionfunc: func(subTest *testing.T, w *httptest.ResponseRecorder) {
				assert.Equal(subTest, http.StatusInternalServerError, w.Code)
			},
			Handler: h.UpdateTweet,
		},
	}

	for _, tc := range testCases {
		tests.RunTest(t, tc, tweetsServiceMock)
	}
}
