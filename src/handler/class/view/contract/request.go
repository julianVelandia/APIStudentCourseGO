package contract

type Request struct {
	Email   string `json:"email" binding:"required"`
	ClassID string `json:"class_id" binding:"required"`
}
