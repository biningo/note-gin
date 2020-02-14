package view

type FolderSelect struct {
	Value int64  `json:"value" form:"value"`
	Label string `json:"label" form:"label"`
	Leaf  bool   `json:"leaf" form:"leaf"`
}
