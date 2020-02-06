package model



type User struct {
	BaseModel
	UserName string `form:"name" json:"name"`
	PassWord string `form:"password" json:"password"`
	Type int 	`form:"type" json:"type"` //用户类型
}

