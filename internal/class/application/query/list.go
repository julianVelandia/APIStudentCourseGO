package query

type List struct {
	page int
}

func (l List) Page() int {
	return l.page
}

func NewList(page int) *List {
	return &List{page: page}
}
