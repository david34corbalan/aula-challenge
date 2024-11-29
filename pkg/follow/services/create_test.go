package follow_test

import (
	"errors"
	"fmt"
	"testing"
	mocks "uala/mocks"
	"uala/pkg/common"
	follow "uala/pkg/follow/domain"
	followServices "uala/pkg/follow/services"
	users "uala/pkg/users/domain"

	"uala/tests"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestService_Create(t *testing.T) {

	followResponse := follow.Follow{
		FollowID: 2,
		UserID:   1,
	}

	tweet := follow.FollowUser{
		FollowID: 2,
		UserID:   1,
	}

	testCases := []tests.TestCaseService{
		{
			Name: "error same user and follow",
			Setup: func(mock ...interface{}) {
			},
			AssertionFunc: func(subTest *testing.T, result interface{}, err error) {
				p := result.(*follow.Follow)
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
				mockTweetRepo := mock[1].(*mocks.MockFollowRepository)
				s := &followServices.Service{
					Repo:     mockTweetRepo,
					RepoUser: mockUsersRepo,
				}
				tweetSame := follow.FollowUser{
					FollowID: 1,
					UserID:   1,
				}

				return s.Create(tweetSame)
			},
		},
		{
			Name: "success",
			Setup: func(mock ...interface{}) {
				mockRepo := mock[0].(*mocks.MockUsersRepository)
				mockRepoFollow := mock[1].(*mocks.MockFollowRepository)
				mockRepo.EXPECT().Get(gomock.Any()).Return(&users.Users{}, nil)
				mockRepo.EXPECT().Get(gomock.Any()).Return(&users.Users{}, nil)
				mockRepoFollow.EXPECT().Create(tweet).Return(&followResponse, nil)
			},
			AssertionFunc: func(subTest *testing.T, result interface{}, err error) {
				p := result.(*follow.Follow)
				assert.NotNil(subTest, p)
				assert.Equal(subTest, followResponse, *p)
				assert.NoError(subTest, err)
			},
			TestFunc: func(mock ...interface{}) (interface{}, error) {
				mockUsersRepo := mock[0].(*mocks.MockUsersRepository)
				mockTweetRepo := mock[1].(*mocks.MockFollowRepository)
				s := &followServices.Service{
					Repo:     mockTweetRepo,
					RepoUser: mockUsersRepo,
				}
				return s.Create(tweet)
			},
		},
		{
			Name: "error same follower",
			Setup: func(mock ...interface{}) {
				mockRepo := mock[0].(*mocks.MockUsersRepository)
				mockRepo.EXPECT().Get(gomock.Any()).Return(nil, errors.New(fmt.Sprintf("follower with id '%d' not found", int(followResponse.FollowID))))
			},
			AssertionFunc: func(subTest *testing.T, result interface{}, err error) {
				assert.NotNil(subTest, err)
				assert.Contains(subTest, err.Error(), "follower with id")
			},
			TestFunc: func(mock ...interface{}) (interface{}, error) {
				mockUsersRepo := mock[0].(*mocks.MockUsersRepository)
				mockTweetRepo := mock[1].(*mocks.MockFollowRepository)
				s := &followServices.Service{
					Repo:     mockTweetRepo,
					RepoUser: mockUsersRepo,
				}
				return s.Create(tweet)
			},
		},
		{
			Name: "error same user",
			Setup: func(mock ...interface{}) {
				mockRepo := mock[0].(*mocks.MockUsersRepository)
				// mockRepoFollow := mock[1].(*mocks.MockFollowRepository)
				mockRepo.EXPECT().Get(gomock.Any()).Return(&users.Users{}, nil)
				mockRepo.EXPECT().Get(gomock.Any()).Return(nil, errors.New(fmt.Sprintf("following with id '%d' not found", int(followResponse.UserID))))
			},
			AssertionFunc: func(subTest *testing.T, result interface{}, err error) {
				assert.NotNil(subTest, err)
				assert.Contains(subTest, err.Error(), "unexpected error")
			},
			TestFunc: func(mock ...interface{}) (interface{}, error) {
				mockUsersRepo := mock[0].(*mocks.MockUsersRepository)
				mockTweetRepo := mock[1].(*mocks.MockFollowRepository)
				s := &followServices.Service{
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
		mockFollowRepo := mocks.NewMockFollowRepository(ctrl)

		tests.RunTestService(t, tc, mockUsersRepo, mockFollowRepo)
	}
}
