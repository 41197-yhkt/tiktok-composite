package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id   int64  `gorm:"primary_key"`
	Name string `gorm:"column:name"`
}

type Vedio struct {
	gorm.Model
	Id         uint      `gorm:"primary_key"`
	AuthorId   int64     `gorm:"column:author_id"`
	PlayUrl    string    `gorm:"column:play_url"`
	CoverUrl   string    `gorm:"column:cover_url"`
	Title      string    `gorm:"column:title"`
	Created_at time.Time `gorm:"column:created_at"`
	Updated_at time.Time `gorm:"column:updated_at"`
}

type UserFavorite struct {
	gorm.Model
	Id      uint  `gorm:"primary_key"`
	UserId  int64 `gorm:"column:user_id"`
	VedioId int64 `gorm:"column:vedio_id"`
}

type Comment struct {
	gorm.Model
	Id      uint   `gorm:"primary_key"`
	UserId  int64  `gorm:"column:user_id"`
	VedioId int64  `gorm:"column:vedio_id"`
	Content string `gorm:"column:content"`
}
