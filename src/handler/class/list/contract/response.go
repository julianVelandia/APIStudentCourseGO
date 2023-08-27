package contract

type Response struct {
	Classes []Class `json:"classes"`
}

type Class struct {
	ClassID      string   `json:"class_id"`
	Title        string   `json:"title"`
	Content      []string `json:"content"`
	CreationDate string   `json:"creation_date"`
	ReadTime     int      `json:"read_time"`
}
