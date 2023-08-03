package query

type List struct {
	page int
	tag  string
	sort string
	//withAnswer bool
	batch    int
	orderID  int
	courseID string
}

func NewListDefault(page int, tag string, sort string, batch int, orderID int, courseID string) *List {
	return &List{courseID: courseID, page: page, tag: tag, sort: sort, batch: batch, orderID: orderID}
}

func (l List) Batch() int {
	return l.batch
}

func (l List) OffSet() int {
	return (l.page - 1) * l.batch
}

func (l List) Tag() string {
	return l.tag
}

func (l List) Sort() string {
	return l.sort
}

func (l List) OrderID() int {
	return l.orderID
}

func (l List) CourseID() string {
	return l.courseID
}
