package view

type DataList struct {
	Items interface{} `form:"items" json:"items"`
	Total int         `form:"total" json:"total"`
}
