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

func Test_CreateTweet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tweetServiceMock := mocks.NewMockTweetsService(ctrl)
	h := tweets.Handler{
		TweetsService: tweetServiceMock,
	}

	gin.SetMode(gin.TestMode)

	testCases := []tests.TestCaseHandlers{
		{
			Name:    "should create a tweet successfully",
			Method:  "POST",
			Url:     "/api/v1/tweets",
			Reqbody: `{"user_id": 1, "comment": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus lacinia odio vitae vestibulum vestibulum"}`,
			Setup: func(mock ...interface{}) {
				mockService := mock[0].(*mocks.MockTweetsService)
				expectedTweet := domain.TweetsCreate{
					UserID:  1,
					Comment: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus lacinia odio vitae vestibulum vestibulum",
				}
				mockService.EXPECT().Create(expectedTweet).Return(&domain.Tweets{
					UserID:  1,
					Comment: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus lacinia odio vitae vestibulum vestibulum",
				}, nil)
			},
			Assertionfunc: func(subTest *testing.T, w *httptest.ResponseRecorder) {
				assert.Equal(subTest, http.StatusCreated, w.Code)
				var tweet domain.Tweets
				err := json.Unmarshal(w.Body.Bytes(), &tweet)
				assert.NoError(subTest, err)
				assert.Equal(subTest, 1, tweet.UserID)
				assert.Equal(subTest, "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus lacinia odio vitae vestibulum vestibulum", tweet.Comment)
			},
			Handler: h.CreateTweet,
		},
		{
			Name:    "should return error for invalid comment",
			Method:  "POST",
			Url:     "/api/v1/tweets",
			Reqbody: `{"user_id": 1, "comment": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus lacinia odio vitae vestibulum vestibulum. Cras venenatis euismod malesuada. Nulla facilisi. Curabitur ac felis arcu. Sed vehicula, urna eu efficitur tincidunt, sapien libero hendrerit est, nec scelerisque scelerisque scelerisque scelerisque"}`,
			Setup: func(mock ...interface{}) {
				// No setup needed for this test case
			},
			Assertionfunc: func(subTest *testing.T, w *httptest.ResponseRecorder) {
				assert.Equal(subTest, http.StatusUnprocessableEntity, w.Code)
			},
			Handler: h.CreateTweet,
		},
		{
			Name:    "should return error create tweet",
			Method:  "POST",
			Url:     "/api/v1/tweets",
			Reqbody: `{"user_id": 1, "comment": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus lacinia odio vitae vestibulum vestibulum"}`,
			Setup: func(mock ...interface{}) {
				mockService := mock[0].(*mocks.MockTweetsService)
				expectedTweet := domain.TweetsCreate{
					UserID:  1,
					Comment: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus lacinia odio vitae vestibulum vestibulum",
				}
				mockService.EXPECT().Create(expectedTweet).Return(nil, common.ErrCreate)
			},
			Assertionfunc: func(subTest *testing.T, w *httptest.ResponseRecorder) {
				assert.Equal(subTest, http.StatusInternalServerError, w.Code)
			},
			Handler: h.CreateTweet,
		},
	}

	for _, tc := range testCases {
		tests.RunTest(t, tc, tweetServiceMock)
	}
}
