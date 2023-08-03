package domain

type Course struct {
	CourseID string
	//StudentID        string
	Title         string
	CreationDate  string
	Content       []string
	IsCodeContent []bool
	ReadTime      float64
	//Likes         int
	//Dislikes      int
	//Views         int
	OrderID int
	Tags    []Tag
}
