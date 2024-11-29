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

func TestService_Get(t *testing.T) {

	id := 1

	tweet := tweets.Tweets{
		ID:        1,
		Comment:   "my first tweet",
		UserID:    id,
		CreatedAt: "",
		UpdatedAt: "",
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
				return s.Get(0)
			},
			AssertionFunc: func(subTest *testing.T, result interface{}, err error) {
				p := result.(*tweets.Tweets)
				assert.Nil(subTest, p)
				assert.EqualError(subTest, err, "id invalid")
			},
		},
		{
			Name: "success",
			Setup: func(mock ...interface{}) {
				mockRepo := mock[0].(*mocks.MockTweetsRepository)
				mockRepo.EXPECT().Get(id).Return(&tweet, nil)
			},
			TestFunc: func(mock ...interface{}) (interface{}, error) {
				mockUsersRepo := mock[0].(*mocks.MockTweetsRepository)
				s := &tweetsServices.Service{
					Repo: mockUsersRepo,
				}
				return s.Get(id)
			},
			AssertionFunc: func(subTest *testing.T, result interface{}, err error) {
				p := result.(*tweets.Tweets)
				assert.NotNil(subTest, p)
				assert.Equal(subTest, tweet.Comment, p.Comment)
				assert.NoError(subTest, err)
			},
		},
		{
			Name: "get user not found",
			Setup: func(mock ...interface{}) {
				mockRepo := mock[0].(*mocks.MockTweetsRepository)
				mockRepo.EXPECT().Get(id).Return(nil, common.ErrNotFound)
			},
			TestFunc: func(mock ...interface{}) (interface{}, error) {
				mockUsersRepo := mock[0].(*mocks.MockTweetsRepository)
				s := &tweetsServices.Service{
					Repo: mockUsersRepo,
				}
				return s.Get(id)
			},
			AssertionFunc: func(subTest *testing.T, result interface{}, err error) {
				p := result.(*tweets.Tweets)
				assert.Nil(subTest, p)
				assert.NotNil(subTest, err)
				assert.Equal(subTest, err, common.NewAppError(
					common.ErrCodeNotFound,
					fmt.Sprintf("tweet with id '%d' not found", id)))
			},
		},
		{
			Name: "error unexpected",
			Setup: func(mock ...interface{}) {
				mockRepo := mock[0].(*mocks.MockTweetsRepository)
				mockRepo.EXPECT().Get(id).Return(nil, errors.New("unexpected error"))
			},
			TestFunc: func(mock ...interface{}) (interface{}, error) {
				mockUsersRepo := mock[0].(*mocks.MockTweetsRepository)
				s := &tweetsServices.Service{
					Repo: mockUsersRepo,
				}
				return s.Get(id)
			},
			AssertionFunc: func(subTest *testing.T, result interface{}, err error) {
				assert.NotNil(subTest, err)
				assert.Contains(subTest, err.Error(), "unexpected error")
			},
		},
	}

	for _, tc := range testCases {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockTweetsRepo := mocks.NewMockTweetsRepository(ctrl)

		tests.RunTestService(t, tc, mockTweetsRepo)
	}
}
