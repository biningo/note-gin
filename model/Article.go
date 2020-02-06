package model

type Article struct {
	BaseModel
	Title string	`form:"title" json:"title"`
	FolderID int64
	Tags []Tag 	`many2many:"article_tag" form:"tags" json:"tags"`
	MkValue string `form:"mkValue" json:"mkValue" type:"text"`
	MkHtml string `form:"mkHtml" json:"mkHtml" type:"text"`
}


//Find


