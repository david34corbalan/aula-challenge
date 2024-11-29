package tweets_test

import (
	"errors"
	"testing"
	mocks "uala/mocks"
	"uala/pkg/common"
	tweets "uala/pkg/tweets/domain"
	tweetsServices "uala/pkg/tweets/services"

	"uala/tests"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestService_Index(t *testing.T) {
	data := []*tweets.Tweets{}

	for i := 0; i < 20; i++ {
		tweetsData := tweets.Tweets{
			ID:        uint(i),
			Comment:   "my frist tweet",
			UserID:    1,
			CreatedAt: "",
			UpdatedAt: "",
		}
		data = append(data, &tweetsData)
	}

	params := common.QuerysParamsPaginate{
		Limit:  10,
		Offset: 0,
	}

	testCases := []tests.TestCaseService{
		{
			Name: "valid limit and offset",
			Setup: func(mock ...interface{}) {
				mockRepo := mock[0].(*mocks.MockTweetsRepository)
				mockRepo.EXPECT().Index(params).Return(nil, 0, nil)
			},
			TestFunc: func(mock ...interface{}) (interface{}, error) {
				mockTweetsRepo := mock[0].(*mocks.MockTweetsRepository)
				s := &tweetsServices.Service{
					Repo: mockTweetsRepo,
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
				mockRepo := mock[0].(*mocks.MockTweetsRepository)
				mockRepo.EXPECT().Index(params).Return(data, 20, nil)
			},
			TestFunc: func(mock ...interface{}) (interface{}, error) {
				mockTweetsRepo := mock[0].(*mocks.MockTweetsRepository)
				s := &tweetsServices.Service{
					Repo: mockTweetsRepo,
				}
				return s.Index(params)
			},
			AssertionFunc: func(subTest *testing.T, result interface{}, err error) {
				p := result.(*common.Paginate)
				assert.Equal(subTest, int(p.Limit), 10)
				assert.Equal(subTest, int(p.Offset), 0)
				assert.Equal(subTest, int(p.Count), 20)
				assert.Equal(subTest, len(p.Data.([]*tweets.Tweets)), 20)
			},
		},
		{
			Name: "error unexpected",
			Setup: func(mock ...interface{}) {
				mockRepo := mock[0].(*mocks.MockTweetsRepository)
				mockRepo.EXPECT().Index(params).Return(data, 20, errors.New("unexpected error"))
			},
			TestFunc: func(mock ...interface{}) (interface{}, error) {
				mockTweetsRepo := mock[0].(*mocks.MockTweetsRepository)
				s := &tweetsServices.Service{
					Repo: mockTweetsRepo,
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
				mockRepo := mock[0].(*mocks.MockTweetsRepository)
				mockRepo.EXPECT().Index(params).Return(data, 20, common.ErrRetrieve)
			},
			TestFunc: func(mock ...interface{}) (interface{}, error) {
				mockTweetsRepo := mock[0].(*mocks.MockTweetsRepository)
				s := &tweetsServices.Service{
					Repo: mockTweetsRepo,
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

		mockTweetsRepo := mocks.NewMockTweetsRepository(ctrl)

		tests.RunTestService(t, tc, mockTweetsRepo)
	}
}
