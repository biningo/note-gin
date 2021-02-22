package common

type DataList struct {
	Items interface{} `form:"items" json:"items"`
	Total int64       `form:"total" json:"total"`
}
