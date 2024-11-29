package tweets_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"uala/cmd/api/handlers/tweets"
	"uala/mocks"
	"uala/pkg/common/validators"
	"uala/tests"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_TimelineTweet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tweetServiceMock := mocks.NewMockTweetsService(ctrl)
	h := tweets.Handler{
		TweetsService: tweetServiceMock,
	}

	gin.SetMode(gin.TestMode)
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("limit_offset", validators.LimitAndOffset)
	}

	testCases := []tests.TestCaseHandlers{
		{
			Name:   "should return status code 422",
			Method: "GET",
			Url:    "/api/v1/tweets/1/timeline?offset=0&limit=0",
			Setup: func(mock ...interface{}) {
				// No setup needed for this test case
			},
			Assertionfunc: func(subTest *testing.T, w *httptest.ResponseRecorder) {
				assert.Equal(subTest, http.StatusUnprocessableEntity, w.Code)
			},
			Handler: h.TimelieTweet,
		},
		// {
		// 	Name:   "should return status code 200",
		// 	Method: "GET",
		// 	Url:    "/api/v1/tweets/1/timeline?offset=0&limit=10",
		// 	Setup: func(mock ...interface{}) {
		// 		mockService := mock[0].(*mocks.MockTweetsService)

		// 		data := []*tweetDomain.Tweets{
		// 			{
		// 				Comment: "test tweet",
		// 				UserID:  1,
		// 			},
		// 		}

		// 		mockService.EXPECT().Timeline(1, 10, 0).Return(data, nil)
		// 	},
		// 	Assertionfunc: func(subTest *testing.T, w *httptest.ResponseRecorder) {
		// 		var response []*tweetDomain.Tweets
		// 		json.Unmarshal(w.Body.Bytes(), &response)
		// 		assert.Equal(subTest, http.StatusOK, w.Code)
		// 	},
		// 	Handler: h.TimelieTweet,
		// },
		// {
		// 	Name:   "should return status code 500",
		// 	Method: "GET",
		// 	Url:    "/api/v1/tweets/1/timeline?offset=0&limit=10",
		// 	Setup: func(mock ...interface{}) {
		// 		mockService := mock[0].(*mocks.MockTweetsService)
		// 		mockService.EXPECT().Timeline(1, 10, 0).Return(nil, common.ErrRetrieve)
		// 	},
		// 	Assertionfunc: func(subTest *testing.T, w *httptest.ResponseRecorder) {
		// 		assert.Equal(subTest, http.StatusInternalServerError, w.Code)
		// 	},
		// 	Handler: h.TimelieTweet,
		// },
	}

	for _, tc := range testCases {
		tests.RunTest(t, tc, tweetServiceMock)
	}
}
