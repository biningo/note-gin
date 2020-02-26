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

func (this MyMood) GetMany(page int) (moods []MyMood, total int) {
	db.Offset((page - 1) * 14).Limit(14).Order("created_at desc").Find(&moods)
	db.Table("my_mood").Count(&total)
	return
}

func (this *MyMood) Delete() {
	db.Delete(&this, "id=?", this.ID)
}

func (this MyMood) DeleteMany(ids []string) {
	db.Table("my_mood").Delete(this, "id in (?)", ids)
}

func (this *MyMood) Add() {
	db.Create(this)
}
