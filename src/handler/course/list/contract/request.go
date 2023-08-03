package contract

type URLParams struct {
	Page string `form:"page" json:"page"`
	Tags string `form:"tags" json:"tags"`
	Sort string `form:"sort" json:"sort"`
}
