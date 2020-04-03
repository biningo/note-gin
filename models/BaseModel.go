package models

import (
	"time"
)

type BaseModel struct {
	ID          int64     ` form:"id" json:"id" gorm:"primary_key"`
	CreatedAt   time.Time `form:"create_at" json:"created_at"`
	UpdatedAt   time.Time `form:"updated_at" json:"updated_at"`
	Deleted     bool      `form:"deleted" json:"deleted"`
	DeletedTime time.Time `form:"deleted_time" json:"deleted_time"`
}
