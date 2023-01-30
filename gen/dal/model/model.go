package model

import (
	"time"

	"gorm.io/gorm"
)

type Vedio struct {
	gorm.Model
	Id         uint      `gorm:"primary_key"`
	AuthorId   string    `gorm:"column:author_id"`
	PlayUrl    string    `gorm:"column:play_url"`
	CoverUrl   string    `gorm:"column:cover_url"`
	Title      string    `gorm:"column:title"`
	Created_at time.Time `gorm:"column:created_at"`
	Updated_at time.Time `gorm:"column:updated_at"`
}

type UserFavorite struct {
	gorm.Model
	Id      uint `gorm:"primary_key"`
	UserId  int  `gorm:"column:user_id"`
	VedioId int  `gorm:"column:vedio_id"`
}
