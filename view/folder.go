package view

/**
*@Author lyer
*@Date 2/22/21 16:34
*@Describe
**/

type FolderInfo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	UpdatedAt string `json:"updated_at"`
}

type FolderSelectView struct {
	Value    int                `json:"value"`
	Label    string             `json:"label"`
	Leaf     bool               `json:"leaf"`
	Children []FolderSelectView `json:"children"`
}
