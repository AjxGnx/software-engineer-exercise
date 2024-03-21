package dto

type Paginate struct {
	Page  int
	Limit int
}

func (p *Paginate) ValidateLimitAndPage() {
	if p.Page < 1 {
		p.Page = 1
	}

	if p.Limit == 0 {
		p.Limit = 10
	}
}

func (p *Paginate) CalculateOffset() int {
	var offset int

	if p.Page > 0 {
		offset = (p.Page - 1) * p.Limit
	}

	return offset
}
