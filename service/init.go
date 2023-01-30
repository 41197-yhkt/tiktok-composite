package service

import (
	"tiktok-composite/gen/dal"
	"tiktok-composite/gen/dal/query"
)

var q *query.Query

func Init() {
	q = query.Use(dal.DB.Debug())
}
