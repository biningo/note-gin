package ArticleView

type ArticleEditView struct {
	ID          int64   `form:"id" json:"id"`
	CreatedAt   string  `json:"created_at" form:"created_at"`
	UpdatedAt   string  `json:"updated_at" form:"updated_at"`
	Title       string  `json:"title" form:"title"`
	DirPath     []int64 `json:"dir_path" form:"dir_path"`
	FolderID    int64   `json:"folder_id" form:"folder_id"`
	FolderTitle string  `json:"folder_title" form:"folder_title"`
	MkValue     string  `form:"mkValue" json:"mkValue"`
}
