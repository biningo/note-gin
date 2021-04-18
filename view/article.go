package view

/**
*@Author lyer
*@Date 2/22/21 16:30
*@Describe
**/
type ArticleEditView struct {
	ID          int    `json:"id"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
	Title       string `json:"title"`
	DirPath     []int  `json:"dirPath"`
	FolderID    int    `json:"folderId"`
	FolderTitle string `json:"folderTitle"`
	Content     string `json:"content"`
}

type ArticleDetail struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type ArticleInfo struct {
	ID        int      `json:"id"`
	Title     string   `json:"title"`
	UpdatedAt string   `json:"updatedAt"`
	Tags      []string `json:"tags" form:"tags"`
}
