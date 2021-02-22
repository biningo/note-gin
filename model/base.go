package model

import (
	"time"
)

type BaseModel struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updateAt"`
	DeletedAt time.Time `json:"deletedAt"`
	Status    int       `json:"status"`
}
