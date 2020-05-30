package models

type RoleType string

const (
	Admin   RoleType = "admin"
	User	RoleType = "user"
	Visitor RoleType = "visitor"
)

type Account struct {
	BaseModel
	LoginName string
	PassWord  string
	Role RoleType
}





