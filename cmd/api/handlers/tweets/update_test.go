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

func Test_UpdateTweets(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tweetsServiceMock := mocks.NewMockTweetsService(ctrl)
	h := tweets.Handler{
		TweetsService: tweetsServiceMock,
	}

	gin.SetMode(gin.TestMode)

	id := "1"
	expectedTweet := domain.TweetsUpdate{
		Comment: "my first comment",
		UserID:  1,
	}

	tweetReturn := &domain.Tweets{
		ID:      1,
		Comment: "my first comment",
		UserID:  1,
	}

	testCases := []tests.TestCaseHandlers{
		{
			Name:    "should update a tweet successfully",
			Method:  "POST",
			Url:     "/api/v1/tweets/" + id,
			Reqbody: `{"comment": "my first comment", "user_id": 1}`,
			Setup: func(mock ...interface{}) {
				mockService := mock[0].(*mocks.MockTweetsService)
				mockService.EXPECT().Update(1, expectedTweet).Return(tweetReturn, nil)
			},
			Assertionfunc: func(subTest *testing.T, w *httptest.ResponseRecorder) {
				assert.Equal(subTest, http.StatusCreated, w.Code)
				var tweet domain.Tweets
				err := json.Unmarshal(w.Body.Bytes(), &tweet)
				assert.NoError(subTest, err)
				assert.Equal(subTest, tweetReturn, &tweet)
			},
			Handler: h.UpdateTweet,
		},
		{
			Name:    "should return error for invalid request",
			Method:  "POST",
			Url:     "/api/v1/tweets/" + id,
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
			Name:    "should return error update tweet",
			Method:  "POST",
			Url:     "/api/v1/tweets/" + id,
			Reqbody: `{"comment": "my first comment", "user_id": 1}`,
			Setup: func(mock ...interface{}) {
				mockService := mock[0].(*mocks.MockTweetsService)
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
