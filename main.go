package main

import (
	"log"
	compoiste "tiktok/kitex_gen/compoiste/douyinservice"
)

func main() {
	svr := compoiste.NewServer(new(DouyinServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
