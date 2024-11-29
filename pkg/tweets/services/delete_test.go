package tweets_test

import (
	"errors"
	"fmt"
	"testing"
	mocks "uala/mocks"
	"uala/pkg/common"
	usersDomain "uala/pkg/users/domain"
	usersServices "uala/pkg/users/services"

	"uala/tests"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestService_Delete(t *testing.T) {

	id := 1

	user := usersDomain.Users{

		Name:     "David",
		LastName: "Gonzalez",
		Email:    "email@email.com.ar",
	}

	testCases := []tests.TestCaseService{
		{
			Name:  "empty id",
			Setup: func(mock ...interface{}) {},
			TestFunc: func(mock ...interface{}) (interface{}, error) {
				mockUsersRepo := mock[0].(*mocks.MockUsersRepository)
				s := &usersServices.Service{
					Repo: mockUsersRepo,
				}
				return s.Delete(0)
			},
			AssertionFunc: func(subTest *testing.T, result interface{}, err error) {
				p := result.(*usersDomain.Users)
				assert.Nil(subTest, p)
				assert.EqualError(subTest, err, "id is required")
			},
		},
		{
			Name: "user not found",
			Setup: func(mock ...interface{}) {
				mockRepo := mock[0].(*mocks.MockUsersRepository)
				mockRepo.EXPECT().Delete(id).Return(nil, common.ErrNotFound)
			},
			TestFunc: func(mock ...interface{}) (interface{}, error) {
				mockUsersRepo := mock[0].(*mocks.MockUsersRepository)
				s := &usersServices.Service{
					Repo: mockUsersRepo,
				}
				return s.Delete(id)
			},
			AssertionFunc: func(subTest *testing.T, result interface{}, err error) {
				p := result.(*usersDomain.Users)
				assert.Nil(subTest, p)
				if errors.As(err, &common.AppError{}) {
					assert.Equal(subTest, err, common.NewAppError(
						common.ErrCodeNotFound,
						fmt.Sprintf("user with id '%d' not found", id)))
				}
			},
		},
		{
			Name: "success",
			Setup: func(mock ...interface{}) {
				mockRepo := mock[0].(*mocks.MockUsersRepository)
				mockRepo.EXPECT().Delete(id).Return(&user, nil)
			},
			TestFunc: func(mock ...interface{}) (interface{}, error) {
				mockUsersRepo := mock[0].(*mocks.MockUsersRepository)
				s := &usersServices.Service{
					Repo: mockUsersRepo,
				}
				return s.Delete(id)
			},
			AssertionFunc: func(subTest *testing.T, result interface{}, err error) {
				p := result.(*usersDomain.Users)
				assert.NotNil(subTest, p)
				assert.Equal(subTest, "David", p.Name)
				assert.NoError(subTest, err)
			},
		},
		{
			Name: "error unexpected",
			Setup: func(mock ...interface{}) {
				mockRepo := mock[0].(*mocks.MockUsersRepository)
				mockRepo.EXPECT().Delete(id).Return(nil, errors.New("unexpected error"))
			},
			TestFunc: func(mock ...interface{}) (interface{}, error) {
				mockUsersRepo := mock[0].(*mocks.MockUsersRepository)
				s := &usersServices.Service{
					Repo: mockUsersRepo,
				}
				return s.Delete(id)
			},
			AssertionFunc: func(subTest *testing.T, result interface{}, err error) {
				assert.NotNil(subTest, err)
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
