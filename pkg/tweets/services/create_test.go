package tweets_test

import (
	"errors"
	"testing"
	mocks "uala/mocks"
	"uala/pkg/common"
	tweets "uala/pkg/tweets/domain"
	tweetsServices "uala/pkg/tweets/services"
	users "uala/pkg/users/domain"

	"uala/tests"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestService_Create(t *testing.T) {

	userResponse := tweets.Tweets{
		Comment: "my first comment",
		UserID:  1,
	}

	tweet := tweets.TweetsCreate{
		Comment: "my first comment",
		UserID:  1,
	}

	testCases := []tests.TestCaseService{
		{
			Name: "error create database",
			Setup: func(mock ...interface{}) {
				mockRepo := mock[0].(*mocks.MockUsersRepository)
				mockRepoTweet := mock[1].(*mocks.MockTweetsRepository)
				mockRepo.EXPECT().Get(gomock.Any()).Return(&users.Users{}, nil)
				mockRepoTweet.EXPECT().Create(tweet).Return(nil, common.ErrCreate)
			},
			AssertionFunc: func(subTest *testing.T, result interface{}, err error) {
				p := result.(*tweets.Tweets)
				assert.Nil(subTest, p)
				if errors.Is(err, common.ErrCreate) {
					assert.Equal(subTest, common.NewAppError(
						common.ErrCodeInternalServer,
						err.Error(),
					), err)
				}
			},
			TestFunc: func(mock ...interface{}) (interface{}, error) {
				mockUsersRepo := mock[0].(*mocks.MockUsersRepository)
				mockTweetRepo := mock[1].(*mocks.MockTweetsRepository)
				s := &tweetsServices.Service{
					Repo:     mockTweetRepo,
					RepoUser: mockUsersRepo,
				}
				return s.Create(tweet)
			},
		},
		{
			Name: "success",
			Setup: func(mock ...interface{}) {
				mockRepo := mock[0].(*mocks.MockUsersRepository)
				mockRepoTweet := mock[1].(*mocks.MockTweetsRepository)
				mockRepo.EXPECT().Get(gomock.Any()).Return(&users.Users{}, nil)
				mockRepoTweet.EXPECT().Create(tweet).Return(&userResponse, nil)
			},
			AssertionFunc: func(subTest *testing.T, result interface{}, err error) {
				p := result.(*tweets.Tweets)
				assert.NotNil(subTest, p)
				assert.Equal(subTest, userResponse.Comment, p.Comment)
				assert.NoError(subTest, err)
			},
			TestFunc: func(mock ...interface{}) (interface{}, error) {
				mockUsersRepo := mock[0].(*mocks.MockUsersRepository)
				mockTweetRepo := mock[1].(*mocks.MockTweetsRepository)
				s := &tweetsServices.Service{
					Repo:     mockTweetRepo,
					RepoUser: mockUsersRepo,
				}
				return s.Create(tweet)
			},
		},
		{
			Name: "error unexpected tweet",
			Setup: func(mock ...interface{}) {
				mockRepo := mock[0].(*mocks.MockUsersRepository)
				mockRepoTweet := mock[1].(*mocks.MockTweetsRepository)
				mockRepo.EXPECT().Get(gomock.Any()).Return(&users.Users{}, nil)
				mockRepoTweet.EXPECT().Create(tweet).Return(nil, errors.New("unexpected error getting tweet"))
			},
			AssertionFunc: func(subTest *testing.T, result interface{}, err error) {
				assert.NotNil(subTest, err)
				assert.Contains(subTest, err.Error(), "unexpected error getting tweet")
			},
			TestFunc: func(mock ...interface{}) (interface{}, error) {
				mockUsersRepo := mock[0].(*mocks.MockUsersRepository)
				mockTweetRepo := mock[1].(*mocks.MockTweetsRepository)
				s := &tweetsServices.Service{
					Repo:     mockTweetRepo,
					RepoUser: mockUsersRepo,
				}
				return s.Create(tweet)
			},
		},
		{
			Name: "error unexpected user",
			Setup: func(mock ...interface{}) {
				mockRepo := mock[0].(*mocks.MockUsersRepository)
				mockRepo.EXPECT().Get(gomock.Any()).Return(nil, errors.New("unexpected error"))
			},
			AssertionFunc: func(subTest *testing.T, result interface{}, err error) {
				assert.NotNil(subTest, err)
				assert.Contains(subTest, err.Error(), "unexpected error")
			},
			TestFunc: func(mock ...interface{}) (interface{}, error) {
				mockUsersRepo := mock[0].(*mocks.MockUsersRepository)
				mockTweetRepo := mock[1].(*mocks.MockTweetsRepository)
				s := &tweetsServices.Service{
					Repo:     mockTweetRepo,
					RepoUser: mockUsersRepo,
				}
				return s.Create(tweet)
			},
		},
	}

	for _, tc := range testCases {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUsersRepo := mocks.NewMockUsersRepository(ctrl)
		mockTweetRepo := mocks.NewMockTweetsRepository(ctrl)

		tests.RunTestService(t, tc, mockUsersRepo, mockTweetRepo)
	}
}
