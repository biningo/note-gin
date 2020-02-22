package model

import "time"

//1 2 3
const (
	_ = iota
	Good
	Middle
	Bad
)

type MyMood struct {
	ID        int64     `form:"id" json:"id"`
	Content   string    `form:"content" json:"content"`
	Status    int       `form:"status" json:"status"` //1 good 2middle 3bad
	CreatedAt time.Time `form:"created_at" json:"created_at"`
}
