package tweets_test

import (
	"errors"
	"fmt"
	"testing"
	mocks "uala/mocks"
	"uala/pkg/common"
	tweets "uala/pkg/tweets/domain"
	tweetsServices "uala/pkg/tweets/services"

	"uala/tests"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestService_Update(t *testing.T) {

	id := 1

	tweet := tweets.Tweets{
		ID:        1,
		Comment:   "my second tweet",
		UserID:    id,
		CreatedAt: "",
		UpdatedAt: "",
	}

	tweetUpdate := tweets.TweetsUpdate{
		Comment: "my second tweet",
		UserID:  id,
	}

	testCases := []tests.TestCaseService{
		{
			Name:  "empty id",
			Setup: func(mock ...interface{}) {},
			TestFunc: func(mock ...interface{}) (interface{}, error) {
				mockUsersRepo := mock[0].(*mocks.MockTweetsRepository)
				s := &tweetsServices.Service{
					Repo: mockUsersRepo,
				}
				return s.Update(0, tweetUpdate)
			},
			AssertionFunc: func(subTest *testing.T, result interface{}, err error) {
				p := result.(*tweets.Tweets)
				assert.Nil(subTest, p)
				assert.EqualError(subTest, err, "id invalid")
			},
		},
		{
			Name: "user not found",
			Setup: func(mock ...interface{}) {
				mockRepo := mock[0].(*mocks.MockTweetsRepository)
				mockRepo.EXPECT().Update(id, tweetUpdate).Return(nil, common.ErrNotFound)
			},
			TestFunc: func(mock ...interface{}) (interface{}, error) {
				mockUsersRepo := mock[0].(*mocks.MockTweetsRepository)
				s := &tweetsServices.Service{
					Repo: mockUsersRepo,
				}
				return s.Update(id, tweetUpdate)
			},
			AssertionFunc: func(subTest *testing.T, result interface{}, err error) {
				p := result.(*tweets.Tweets)
				assert.Nil(subTest, p)
				if errors.As(err, &common.AppError{}) {
					assert.Equal(subTest, err, common.NewAppError(
						common.ErrCodeNotFound,
						fmt.Sprintf("tweet with id '%d' not found", id)))
				}
			},
		},
		{
			Name: "error unexpected",
			Setup: func(mock ...interface{}) {
				mockRepo := mock[0].(*mocks.MockTweetsRepository)
				mockRepo.EXPECT().Update(id, tweetUpdate).Return(nil, errors.New("unexpected error"))
			},
			TestFunc: func(mock ...interface{}) (interface{}, error) {
				mockUsersRepo := mock[0].(*mocks.MockTweetsRepository)
				s := &tweetsServices.Service{
					Repo: mockUsersRepo,
				}
				return s.Update(id, tweetUpdate)
			},
			AssertionFunc: func(subTest *testing.T, result interface{}, err error) {
				assert.NotNil(subTest, err)
				assert.Contains(subTest, err.Error(), "unexpected error")
			},
		},
		{
			Name: "success",
			Setup: func(mock ...interface{}) {
				mockRepo := mock[0].(*mocks.MockTweetsRepository)
				mockRepo.EXPECT().Update(id, tweetUpdate).Return(&tweet, nil)
			},
			TestFunc: func(mock ...interface{}) (interface{}, error) {
				mockUsersRepo := mock[0].(*mocks.MockTweetsRepository)
				s := &tweetsServices.Service{
					Repo: mockUsersRepo,
				}
				return s.Update(id, tweetUpdate)
			},
			AssertionFunc: func(subTest *testing.T, result interface{}, err error) {
				p := result.(*tweets.Tweets)
				assert.NotNil(subTest, p)
				assert.Equal(subTest, tweetUpdate.Comment, p.Comment)
				assert.NoError(subTest, err)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(subTest *testing.T) {

			ctrl := gomock.NewController(subTest)
			defer ctrl.Finish()

			mockUsersRepo := mocks.NewMockTweetsRepository(ctrl)

			tests.RunTestService(t, tc, mockUsersRepo)
		})
	}
}
