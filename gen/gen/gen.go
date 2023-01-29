package main

import (
	"tiktok/gen/dal/model"

	"gorm.io/gen"
)

func main() {
	// 我们通过 gen.NewGenerator 来构造一个【代码生成器】，指定我们要生成的代码要放到 dal 下面的 query 子包，生成模式暂时用 default 就ok
	g := gen.NewGenerator(gen.Config{
		OutPath: "../dal/query",
		Mode:    gen.WithDefaultQuery,
	})

	// 调用 ApplyBasic 基于 model 来生成基础 DAL 代码
	g.ApplyBasic(model.Vedio{}, model.UserFavorite{})

	// 调用 ApplyInterface，指明我们希望基于什么 model 和 interface 来生成自定义的接口实现
	g.ApplyInterface(func(model.VedioMethod) {}, model.Vedio{})
	g.ApplyInterface(func(model.UserFavoriteMethod) {}, model.UserFavorite{})

	// 最后调用 Execute 方法来触发生成
	g.Execute()
}
