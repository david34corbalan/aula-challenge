package users_test

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

func TestService_Update(t *testing.T) {

	id := 1

	user := usersDomain.Users{
		ID:       uint(id),
		Name:     "David 2",
		LastName: "Gonzalez 2",
		Email:    "email2@email.com.ar",
	}

	userUpdate := usersDomain.UserUpdate{
		Name:     "David 2",
		LastName: "Gonzalez 2",
		Email:    "email2@email.com.ar",
	}

	testCases := []tests.TestCaseService{
		{
			Name: "user not found",
			Setup: func(mock ...interface{}) {
				mockRepo := mock[0].(*mocks.MockUsersRepository)
				mockRepo.EXPECT().Update(id, userUpdate).Return(nil, common.ErrNotFound)
			},
			TestFunc: func(mock ...interface{}) (interface{}, error) {
				mockUsersRepo := mock[0].(*mocks.MockUsersRepository)
				s := &usersServices.Service{
					Repo: mockUsersRepo,
				}
				return s.Update(id, userUpdate)
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
			Name: "error unexpected",
			Setup: func(mock ...interface{}) {
				mockRepo := mock[0].(*mocks.MockUsersRepository)
				mockRepo.EXPECT().Update(id, userUpdate).Return(nil, errors.New("unexpected error"))
			},
			TestFunc: func(mock ...interface{}) (interface{}, error) {
				mockUsersRepo := mock[0].(*mocks.MockUsersRepository)
				s := &usersServices.Service{
					Repo: mockUsersRepo,
				}
				return s.Update(id, userUpdate)
			},
			AssertionFunc: func(subTest *testing.T, result interface{}, err error) {
				assert.NotNil(subTest, err)
				assert.Contains(subTest, err.Error(), "unexpected error")
			},
		},
		{
			Name: "success",
			Setup: func(mock ...interface{}) {
				mockRepo := mock[0].(*mocks.MockUsersRepository)
				mockRepo.EXPECT().Update(id, userUpdate).Return(&user, nil)
			},
			TestFunc: func(mock ...interface{}) (interface{}, error) {
				mockUsersRepo := mock[0].(*mocks.MockUsersRepository)
				s := &usersServices.Service{
					Repo: mockUsersRepo,
				}
				return s.Update(id, userUpdate)
			},
			AssertionFunc: func(subTest *testing.T, result interface{}, err error) {
				p := result.(*usersDomain.Users)
				assert.NotNil(subTest, p)
				assert.Equal(subTest, "David 2", p.Name)
				assert.NoError(subTest, err)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(subTest *testing.T) {

			ctrl := gomock.NewController(subTest)
			defer ctrl.Finish()

			mockUsersRepo := mocks.NewMockUsersRepository(ctrl)

			tests.RunTestService(t, tc, mockUsersRepo)
		})
	}
}
