package users_test

import (
	"errors"
	"testing"
	mocks "uala/mocks"
	"uala/pkg/common"
	usersDomain "uala/pkg/users/domain"
	usersServices "uala/pkg/users/services"

	"uala/tests"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestService_Create(t *testing.T) {

	userResponse := usersDomain.Users{

		Name:     "David",
		LastName: "Gonzalez",
		Email:    "email@email.com.ar",
	}

	user := usersDomain.UserCreate{
		Name:     userResponse.Name,
		LastName: userResponse.LastName,
		Email:    userResponse.Email,
	}

	testCases := []tests.TestCaseService{
		{
			Name: "error create database",
			Setup: func(mock ...interface{}) {
				mockRepo := mock[0].(*mocks.MockUsersRepository)
				mockRepo.EXPECT().Create(user).Return(nil, common.ErrCreate)
			},
			AssertionFunc: func(subTest *testing.T, result interface{}, err error) {
				p := result.(*usersDomain.Users)
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
				s := &usersServices.Service{
					Repo: mockUsersRepo,
				}
				return s.Create(user)
			},
		},
		{
			Name: "success",
			Setup: func(mock ...interface{}) {
				mockRepo := mock[0].(*mocks.MockUsersRepository)
				mockRepo.EXPECT().Create(user).Return(&userResponse, nil)
			},
			AssertionFunc: func(subTest *testing.T, result interface{}, err error) {
				p := result.(*usersDomain.Users)
				assert.NotNil(subTest, p)
				assert.Equal(subTest, "David", p.Name)
				assert.NoError(subTest, err)
			},
			TestFunc: func(mock ...interface{}) (interface{}, error) {
				mockUsersRepo := mock[0].(*mocks.MockUsersRepository)
				s := &usersServices.Service{
					Repo: mockUsersRepo,
				}
				return s.Create(user)
			},
		},
		{
			Name: "error unexpected",
			Setup: func(mock ...interface{}) {
				mockRepo := mock[0].(*mocks.MockUsersRepository)
				mockRepo.EXPECT().Create(user).Return(nil, errors.New("unexpected error"))
			},
			AssertionFunc: func(subTest *testing.T, result interface{}, err error) {
				assert.NotNil(subTest, err)
				assert.Contains(subTest, err.Error(), "unexpected error")
			},
			TestFunc: func(mock ...interface{}) (interface{}, error) {
				mockUsersRepo := mock[0].(*mocks.MockUsersRepository)
				s := &usersServices.Service{
					Repo: mockUsersRepo,
				}
				return s.Create(user)
			},
		},
	}

	for _, tc := range testCases {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUsersRepo := mocks.NewMockUsersRepository(ctrl)

		tests.RunTestService(t, tc, mockUsersRepo)
	}
}
