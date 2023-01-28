package main

import (
	"log"
	composite "tiktok/kitex_gen/composite/compositeservice"
)

func main() {
	svr := composite.NewServer(new(CompositeServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
