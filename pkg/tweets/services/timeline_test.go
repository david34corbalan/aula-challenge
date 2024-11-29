package tweets_test

import (
	"errors"
	"testing"
	mocks "uala/mocks"
	"uala/pkg/common"
	tweets "uala/pkg/tweets/domain"
	tweetsServices "uala/pkg/tweets/services"
	usersDomain "uala/pkg/users/domain"

	"uala/tests"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestService_Timeline(t *testing.T) {
	data := []*tweets.TweetsUser{}
	user := usersDomain.Users{
		ID:       1,
		Name:     "test",
		LastName: "test",
		Email:    "email@email.com.ar",
	}

	for i := 0; i < 20; i++ {
		tweetsData := tweets.TweetsUser{
			ID:        uint(i),
			Comment:   "my first tweet",
			UserID:    1,
			CreatedAt: "",
			UpdatedAt: "",
		}
		data = append(data, &tweetsData)
	}
	ids := []int{2, 3}

	testCases := []tests.TestCaseService{
		{
			Name: "valid limit and offset",
			Setup: func(mock ...interface{}) {
				mockUserRepo := mock[0].(*mocks.MockUsersRepository)
				mockFollowRepo := mock[1].(*mocks.MockFollowRepository)
				mockTweetsRepo := mock[2].(*mocks.MockTweetsRepository)

				mockUserRepo.EXPECT().Get(gomock.Any()).Return(&user, nil)
				mockFollowRepo.EXPECT().GetFollowers(gomock.Any()).Return(ids, nil)
				mockTweetsRepo.EXPECT().Timeline(ids, 10, 0).Return(data, nil)
			},
			TestFunc: func(mock ...interface{}) (interface{}, error) {
				mockUserRepo := mock[0].(*mocks.MockUsersRepository)
				mockFollowRepo := mock[1].(*mocks.MockFollowRepository)
				mockTweetsRepo := mock[2].(*mocks.MockTweetsRepository)

				s := &tweetsServices.Service{
					Repo:       mockTweetsRepo,
					RepoUser:   mockUserRepo,
					RepoFollow: mockFollowRepo,
				}
				return s.Timeline(1, 10, 0)
			},
			AssertionFunc: func(subTest *testing.T, result interface{}, err error) {
				assert.Nil(subTest, err)
				assert.Equal(subTest, len(result.([]*tweets.TweetsUser)), 20)
			},
		},
		{
			Name: "user not found",
			Setup: func(mock ...interface{}) {
				mockUserRepo := mock[0].(*mocks.MockUsersRepository)
				mockUserRepo.EXPECT().Get(1).Return(nil, common.ErrNotFound)
			},
			TestFunc: func(mock ...interface{}) (interface{}, error) {
				mockUserRepo := mock[0].(*mocks.MockUsersRepository)
				mockFollowRepo := mock[1].(*mocks.MockFollowRepository)
				mockTweetsRepo := mock[2].(*mocks.MockTweetsRepository)

				s := &tweetsServices.Service{
					Repo:       mockTweetsRepo,
					RepoUser:   mockUserRepo,
					RepoFollow: mockFollowRepo,
				}
				return s.Timeline(1, 10, 0)
			},
			AssertionFunc: func(subTest *testing.T, result interface{}, err error) {
				assert.NotNil(subTest, err)
				assert.Equal(subTest, err, common.NewAppError(
					common.ErrCodeNotFound,
					"user with id '1' not found",
				))
			},
		},
		{
			Name: "unexpected error getting user",
			Setup: func(mock ...interface{}) {
				mockUserRepo := mock[0].(*mocks.MockUsersRepository)
				mockUserRepo.EXPECT().Get(1).Return(nil, errors.New("unexpected error"))
			},
			TestFunc: func(mock ...interface{}) (interface{}, error) {
				mockUserRepo := mock[0].(*mocks.MockUsersRepository)
				mockFollowRepo := mock[1].(*mocks.MockFollowRepository)
				mockTweetsRepo := mock[2].(*mocks.MockTweetsRepository)

				s := &tweetsServices.Service{
					Repo:       mockTweetsRepo,
					RepoUser:   mockUserRepo,
					RepoFollow: mockFollowRepo,
				}
				return s.Timeline(1, 10, 0)
			},
			AssertionFunc: func(subTest *testing.T, result interface{}, err error) {
				assert.NotNil(subTest, err)
			},
		},
		{
			Name: "unexpected error getting followers",
			Setup: func(mock ...interface{}) {
				mockUserRepo := mock[0].(*mocks.MockUsersRepository)
				mockFollowRepo := mock[1].(*mocks.MockFollowRepository)
				// mockTweetsRepo := mock[2].(*mocks.MockTweetsRepository)

				mockUserRepo.EXPECT().Get(1).Return(&user, nil)
				mockFollowRepo.EXPECT().GetFollowers(1).Return(nil, errors.New("unexpected error"))
			},
			TestFunc: func(mock ...interface{}) (interface{}, error) {
				mockUserRepo := mock[0].(*mocks.MockUsersRepository)
				mockFollowRepo := mock[1].(*mocks.MockFollowRepository)
				mockTweetsRepo := mock[2].(*mocks.MockTweetsRepository)

				s := &tweetsServices.Service{
					Repo:       mockTweetsRepo,
					RepoUser:   mockUserRepo,
					RepoFollow: mockFollowRepo,
				}
				return s.Timeline(1, 10, 0)
			},
			AssertionFunc: func(subTest *testing.T, result interface{}, err error) {
				assert.NotNil(subTest, err)
			},
		},
		{
			Name: "unexpected error getting tweets",
			Setup: func(mock ...interface{}) {
				mockUserRepo := mock[0].(*mocks.MockUsersRepository)
				mockFollowRepo := mock[1].(*mocks.MockFollowRepository)
				mockTweetsRepo := mock[2].(*mocks.MockTweetsRepository)

				mockUserRepo.EXPECT().Get(1).Return(&user, nil)
				mockFollowRepo.EXPECT().GetFollowers(1).Return(ids, nil)
				mockTweetsRepo.EXPECT().Timeline(ids, 10, 0).Return(nil, errors.New("unexpected error"))
			},
			TestFunc: func(mock ...interface{}) (interface{}, error) {
				mockUserRepo := mock[0].(*mocks.MockUsersRepository)
				mockFollowRepo := mock[1].(*mocks.MockFollowRepository)
				mockTweetsRepo := mock[2].(*mocks.MockTweetsRepository)

				s := &tweetsServices.Service{
					Repo:       mockTweetsRepo,
					RepoUser:   mockUserRepo,
					RepoFollow: mockFollowRepo,
				}
				return s.Timeline(1, 10, 0)
			},
			AssertionFunc: func(subTest *testing.T, result interface{}, err error) {
				assert.NotNil(subTest, err)
			},
		},
	}

	for _, tc := range testCases {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUserRepo := mocks.NewMockUsersRepository(ctrl)
		mockFollowRepo := mocks.NewMockFollowRepository(ctrl)
		mockTweetsRepo := mocks.NewMockTweetsRepository(ctrl)

		tests.RunTestService(t, tc, mockUserRepo, mockFollowRepo, mockTweetsRepo)
	}
}
