package model

import (
	"time"
)

type BaseModel struct {
	ID uint64 ` form:"id" json:"id" gorm:"primary_key"`
	CreatedAt time.Time `form:"create_at" json:"created_at"`
	UpdatedAt time.Time	`form:"updated_at" json:"updated_at"`
	DeletedAt *time.Time	`form:"deleted_at" json:"deleted_at" sql:"index"`
}
