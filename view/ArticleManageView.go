package view

type ArticleManageView struct {
	ID        int64  `json:"id" form:"id"`
	Title     string `json:"title" form:"title"`
	UpdatedAt string `json:"updated_at" form:"updated_at"`
}
