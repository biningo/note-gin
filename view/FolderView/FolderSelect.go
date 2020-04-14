package FolderView

type FolderSelectView struct {
	Value    int64          `json:"value" form:"value"`
	Label    string         `json:"label" form:"label"`
	Leaf     bool           `json:"leaf" form:"leaf"`
	Children []FolderSelectView `json:"children" form:"children"`
}
