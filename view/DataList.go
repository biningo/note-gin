package view

type DataList struct {
	Items interface{} `form:"items" json:"items"`
	Total uint64	`form:"total" json:"total"`
}
