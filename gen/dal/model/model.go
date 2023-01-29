package model

import (
	"time"

	"gorm.io/gorm"
)

type Vedio struct {
	gorm.Model
	Id         uint      `gorm:"primary_key"`
	authorId   string    `gorm:"column:author_id"`
	playUrl    string    `gorm:"column:play_url"`
	coverUrl   string    `gorm:"column:cover_url"`
	title      string    `gorm:"column:title"`
	created_at time.Time `gorm:"column:created_at"`
	updated_at time.Time `gorm:"column:updated_at"`
}

type UserFavorite struct {
	gorm.Model
	Id      uint `gorm:"primary_key"`
	userId  int  `gorm:"column:user_id"`
	vedioId int  `gorm:"column:vedio_id"`
}
