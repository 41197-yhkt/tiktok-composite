package dal

import (
	"sync"
	"tiktok/gen/dal/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var once sync.Once

func init() {
	once.Do(func() {
		DB = ConnctDB().Debug()
		_ = DB.AutoMigrate(&model.Vedio{}, &model.UserFavorite{})
	})
}

func ConnctDB() (conn *gorm.DB) {
	conn, err := gorm.Open(mysql.Open("gorm:gorm@tcp(localhost:9910)/gorm?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return conn
}
