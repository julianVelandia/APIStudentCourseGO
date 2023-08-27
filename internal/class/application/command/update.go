package command

type Update struct {
	email   string
	classID string
	title   string
}

func (u Update) Email() string {
	return u.email
}

func (u Update) ClassID() string {
	return u.classID
}

func (u Update) Title() string {
	return u.title
}

func NewUpdate(email string, classID string, title string) *Update {
	return &Update{email: email, classID: classID, title: title}
}
