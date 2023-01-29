package main

import (
	"log"
	dal "tiktok-composite/gen/dal"
	composite "tiktok-composite/kitex_gen/composite/compositeservice"
)

func Init() {
	dal.Init()
}

func main() {
	Init()
	svr := composite.NewServer(new(CompositeServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
