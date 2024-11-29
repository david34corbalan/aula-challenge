package common

import "math"

type Paginate struct {
	Data       any   `json:"data"`
	Limit      int64 `json:"limit" default:"10"`
	Offset     int64 `json:"offset" default:"0"`
	Count      int64 `json:"count"`
	LastOffset int64 `json:"last_offset"`

	Total      int64     `json:"total"`
	TotalPages int64     `json:"total_pages"`
	Page       int64     `json:"page"`
	Links      Links     `json:"links"`
	NextPages  NextPages `json:"next_pages"`
	PrevPages  PrevPages `json:"prev_pages"`
}

type Links struct {
	NextOffset int64 `json:"next_offset"`
	PrevOffset int64 `json:"prev_offset"`
}

type NextPages struct {
	NextPage1 int64 `json:"next_page_1"`
	NextPage2 int64 `json:"next_page_2"`
}

type PrevPages struct {
	PrevPage1 int64 `json:"prev_page_1"`
	PrevPage2 int64 `json:"prev_page_2"`
}

type QuerysParamsPaginate struct {
	Offset int    `form:"offset" json:"offset" binding:"limit_offset"`
	Limit  int    `form:"limit" json:"limit" binding:"limit_offset"`
	Search string `form:"search" json:"search,omitempty"`
}

func NewPaginate(data any, limit int, offset int, count int) *Paginate {
	return &Paginate{Data: data, Limit: int64(limit), Offset: int64(offset), Count: int64(count)}
}

func (p *Paginate) Invoke() *Paginate {
	return &Paginate{
		Data:       p.Data,
		Count:      int64(p.Count),
		Limit:      int64(p.Limit),
		Offset:     int64(int(p.Offset)),
		LastOffset: getLastPage(int(p.Limit), int(p.Count)),
		Links: Links{
			NextOffset: getNextPage(int(p.Offset), int(p.Limit), int(p.Count)),
			PrevOffset: getPrevPages(int(p.Offset), int(p.Limit)),
		},
		NextPages: NextPages{
			NextPage1: getNextPage(int(p.Offset), int(p.Limit), int(p.Count)),
			NextPage2: getNextsPage(int(p.Offset), int(p.Limit), int(p.Count)),
		},
		PrevPages: PrevPages{
			PrevPage1: getPrevPages(int(p.Offset), int(p.Limit)),
			PrevPage2: getPrevsPages(int(p.Offset), int(p.Limit)),
		},
		Page:       int64(int(p.Offset)/int(p.Limit) + 1),
		TotalPages: int64(int(p.Count)/int(p.Limit) + 1),
		Total:      int64(int(p.Count)),
	}
}

func getNextPage(offset int, limit int, count int) int64 {
	if offset+limit < count {
		return int64(offset + limit)
	}
	return 0
}

func getNextsPage(offset int, limit int, count int) int64 {
	if offset+limit*2 < count {
		return int64(offset + limit*2)
	}
	return 0
}

func getPrevPages(offset int, limit int) int64 {
	if offset-limit > 0 {
		return int64(offset - limit)
	}
	return 0
}

func getPrevsPages(offset int, limit int) int64 {
	if offset-limit*2 > 0 {
		return int64(offset - limit*2)
	}
	return 0
}

func getLastPage(limit int, count int) int64 {
	if count > 0 {
		return int64(math.Ceil(float64(count)/float64(limit)-1) * float64(limit))
	}
	return 0
}
