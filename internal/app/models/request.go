package models

type (
	PaginationRequest struct {
		Page   int    `query:"page"`
		Limit  int    `query:"limit"`
		Search string `query:"search"`
	}
)

func (p *PaginationRequest) SetDefaultPage() {
	p.Page = 1
}

func (p *PaginationRequest) SetDefaultLimit() {
	p.Limit = 10
}

func (p *PaginationRequest) SetDefaults() {
	p.SetDefaultPage()
	p.SetDefaultLimit()
}
