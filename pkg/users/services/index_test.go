package users_test

import (
	"errors"
	"fmt"
	"testing"
	mocks "uala/mocks"
	"uala/pkg/common"
	users "uala/pkg/users/domain"
	usersServices "uala/pkg/users/services"

	"uala/tests"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestService_Index(t *testing.T) {
	data := []*users.Users{}

	for i := 0; i < 20; i++ {
		newUser := users.Users{
			Name:     fmt.Sprintf("Name %d", i),
			Email:    fmt.Sprintf("email %d", i),
			LastName: fmt.Sprintf("lastname %d", i),
		}
		data = append(data, &newUser)
	}

	params := common.QuerysParamsPaginate{
		Limit:  10,
		Offset: 0,
	}

	testCases := []tests.TestCaseService{
		{
			Name: "valid limit and offset",
			Setup: func(mock ...interface{}) {
				mockRepo := mock[0].(*mocks.MockUsersRepository)
				mockRepo.EXPECT().Index(params).Return(nil, 0, nil)
			},
			TestFunc: func(mock ...interface{}) (interface{}, error) {
				mockUsersRepo := mock[0].(*mocks.MockUsersRepository)
				s := &usersServices.Service{
					Repo: mockUsersRepo,
				}
				return s.Index(params)
			},
			AssertionFunc: func(subTest *testing.T, result interface{}, err error) {
				p := result.(*common.Paginate)
				assert.Equal(subTest, int(p.Limit), 10)
				assert.Equal(subTest, int(p.Offset), 0)
				assert.Equal(subTest, int(p.Count), 0)
			},
		},
		{
			Name: "data should length 20",
			Setup: func(mock ...interface{}) {
				mockRepo := mock[0].(*mocks.MockUsersRepository)
				mockRepo.EXPECT().Index(params).Return(data, 20, nil)
			},
			TestFunc: func(mock ...interface{}) (interface{}, error) {
				mockUsersRepo := mock[0].(*mocks.MockUsersRepository)
				s := &usersServices.Service{
					Repo: mockUsersRepo,
				}
				return s.Index(params)
			},
			AssertionFunc: func(subTest *testing.T, result interface{}, err error) {
				p := result.(*common.Paginate)
				assert.Equal(subTest, int(p.Limit), 10)
				assert.Equal(subTest, int(p.Offset), 0)
				assert.Equal(subTest, int(p.Count), 20)
				assert.Equal(subTest, len(p.Data.([]*users.Users)), 20)
			},
		},
		{
			Name: "error unexpected",
			Setup: func(mock ...interface{}) {
				mockRepo := mock[0].(*mocks.MockUsersRepository)
				mockRepo.EXPECT().Index(params).Return(data, 20, errors.New("unexpected error"))
			},
			TestFunc: func(mock ...interface{}) (interface{}, error) {
				mockUsersRepo := mock[0].(*mocks.MockUsersRepository)
				s := &usersServices.Service{
					Repo: mockUsersRepo,
				}
				return s.Index(params)
			},
			AssertionFunc: func(subTest *testing.T, result interface{}, err error) {
				assert.NotNil(subTest, err)
			},
		},
		{
			Name: "should return error common.ErrRetrieve",
			Setup: func(mock ...interface{}) {
				mockRepo := mock[0].(*mocks.MockUsersRepository)
				mockRepo.EXPECT().Index(params).Return(data, 20, common.ErrRetrieve)
			},
			TestFunc: func(mock ...interface{}) (interface{}, error) {
				mockUsersRepo := mock[0].(*mocks.MockUsersRepository)
				s := &usersServices.Service{
					Repo: mockUsersRepo,
				}
				return s.Index(params)
			},
			AssertionFunc: func(subTest *testing.T, result interface{}, err error) {
				assert.NotNil(subTest, err)
				assert.Equal(subTest, err, common.NewAppError(
					common.ErrCodeInternalServer,
					common.ErrRetrieve.Error(),
				))
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
