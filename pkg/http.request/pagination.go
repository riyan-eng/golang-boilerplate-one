package httprequest

type PaginationReq struct {
	Page   int `query:"page"`
	Limit  int `query:"per_page"`
	Search string `query:"search"`
	Order  string `query:"order"`
}

func (p PaginationReq) Init() PaginationReq {
	p.Page = 1
	p.Limit = 10
	p.Order = "desc"
	return p
}