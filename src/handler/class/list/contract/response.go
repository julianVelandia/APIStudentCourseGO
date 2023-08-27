package contract

type Response struct {
	Courses []Course `json:"courses"`
}

type Course struct {
	CourseID string   `json:"course_id"`
	Title    string   `json:"title"`
	Content  string   `json:"content"`
	Tags     []string `json:"tags"`
}
