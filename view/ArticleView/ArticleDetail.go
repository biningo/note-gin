package ArticleView

type ArticleDetail struct {
	ID      int64  `form:"id" json:"id"`
	Title   string `form:"title" json:"title"`
	MkValue string `form:"mkValue" json:"mkValue"`
}
