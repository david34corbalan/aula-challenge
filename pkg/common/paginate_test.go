package common_test

import (
	"testing"
	"uala/pkg/common"

	"github.com/stretchr/testify/assert"
)

type Data struct {
	Id   int
	Name string
}

func Test_Paginate(t *testing.T) {

	data := []Data{
		{Id: 1, Name: "David"},
		{Id: 2, Name: "Gonzalez"},
		{Id: 3, Name: "email"},
		{Id: 4, Name: "email"},
	}
	dataEmpty := []Data{}

	limit := 10
	offset := 0
	count := 4
	countEmpty := 0

	testTable := map[string]struct {
		setup         func() *common.Paginate
		nameTest      string
		assertionFunc func(subTest *testing.T, pag *common.Paginate)
	}{
		"get paginate": {
			setup: func() *common.Paginate {
				pag := common.NewPaginate(data, limit, offset, count)
				return pag.Invoke()
			},
			nameTest: "get paginate",
			assertionFunc: func(subTest *testing.T, paginate *common.Paginate) {
				assert.Equal(subTest, paginate.Data, data)
				assert.Equal(subTest, paginate.Limit, int64(limit))
				assert.Equal(subTest, paginate.Offset, int64(offset))
				assert.Equal(subTest, paginate.Count, int64(count))
			},
		},
		"paginate empty": {
			setup: func() *common.Paginate {
				return common.NewPaginate(dataEmpty, limit, offset, countEmpty)
			},
			nameTest: "paginate empty",
			assertionFunc: func(subTest *testing.T, paginate *common.Paginate) {
				assert.Equal(subTest, paginate.Data, dataEmpty)
				assert.Equal(subTest, paginate.Limit, int64(limit))
				assert.Equal(subTest, paginate.Offset, int64(offset))
				assert.Equal(subTest, paginate.Count, int64(countEmpty))
			},
		},
		"should return 0 for paginate getNextPage": {
			setup: func() *common.Paginate {
				return common.NewPaginate(data, limit, offset, count)
			},
			nameTest: "getNextPage 0",
			assertionFunc: func(subTest *testing.T, paginate *common.Paginate) {
				pag := paginate.Invoke()
				assert.Equal(subTest, int64(0), pag.Links.NextOffset)
			},
		},
		"should return 10 for paginate getNextPage": {
			setup: func() *common.Paginate {
				return common.NewPaginate(data, limit, offset, 100)
			},
			nameTest: "getNextPage 20",
			assertionFunc: func(subTest *testing.T, paginate *common.Paginate) {
				pag := paginate.Invoke()
				assert.Equal(subTest, int64(offset+limit), pag.Links.NextOffset)
			},
		},
		"should return 140 for paginate getPrevPages": {
			setup: func() *common.Paginate {
				return common.NewPaginate(data, 10, 150, count)
			},
			nameTest: "getPrevPages 140",
			assertionFunc: func(subTest *testing.T, paginate *common.Paginate) {
				pag := paginate.Invoke()
				assert.Equal(subTest, int64(140), pag.PrevPages.PrevPage1)
			},
		},
		"should return 130 for paginate getPrevPages": {
			setup: func() *common.Paginate {
				return common.NewPaginate(data, 10, 150, 100)
			},
			nameTest: "getPrevPages 130",
			assertionFunc: func(subTest *testing.T, paginate *common.Paginate) {
				pag := paginate.Invoke()
				assert.Equal(subTest, int64(130), pag.PrevPages.PrevPage2)
			},
		},
		"should return LastOffset": {
			setup: func() *common.Paginate {
				return common.NewPaginate(data, 10, 0, 100)
			},
			nameTest: "LastOffset",
			assertionFunc: func(subTest *testing.T, paginate *common.Paginate) {
				pag := paginate.Invoke()
				assert.Equal(subTest, int64(90), pag.LastOffset)
			},
		},
		"should return 0 LastOffset": {
			setup: func() *common.Paginate {
				return common.NewPaginate(data, 10, 0, 0)
			},
			nameTest: "LastOffset 0",
			assertionFunc: func(subTest *testing.T, paginate *common.Paginate) {
				pag := paginate.Invoke()
				assert.Equal(subTest, int64(0), pag.LastOffset)
			},
		},
	}

	for name, test := range testTable {
		t.Run(name, func(subTest *testing.T) {
			test.nameTest = name
			test.assertionFunc(subTest, test.setup())
		})
	}
}
