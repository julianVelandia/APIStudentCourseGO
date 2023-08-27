package query

type View struct {
	email   string
	classID string
	title   string
}

func (v View) Title() string {
	return v.title
}

func (v View) Email() string {
	return v.email
}

func (v View) ClassID() string {
	return v.classID
}

func NewView(email string, classID string, title string) *View {
	return &View{email: email, classID: classID, title: title}
}
