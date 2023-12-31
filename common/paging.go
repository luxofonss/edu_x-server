package common

import "strings"

type Paging struct {
	Page       int    `json:"page" form:"page"`
	Limit      int    `json:"limit" form:"limit"`
	Total      int    `json:"total" form:"total"`
	FakeCursor string `json:"fake_cursor" form:"fake_cursor"`
	NextCursor string `json:"next_cursor" form:"next_cursor"`
}

func (p *Paging) Fulfill() {
	if p.Page == 0 {
		p.Page = 1
	}
	if p.Limit == 0 {
		p.Limit = 20
	}

	p.FakeCursor = strings.TrimSpace(p.FakeCursor)
}
