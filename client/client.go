package main

import (
	"context"
	"fmt"
	"tiktok-composite/kitex_gen/composite"
	"tiktok-composite/kitex_gen/composite/compositeservice"

	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/pkg/constants"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var userFavoriteClient compositeservice.Client

// 初始化 client
func initUserFavoriteRpc() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := compositeservice.NewClient(
		"composite",
		client.WithResolver(r),
	)
	if err != nil {
		panic(err)
	}
	userFavoriteClient = c
}

// client 模板
func main() {
	initUserFavoriteRpc()
	// 自己提供 ctx
	ctx := context.Background()
	// 定义请求
	req := &composite.BasicFavoriteActionRequest{VedioId: 1, UserId: 1}
	// 调用 BasicFavoriteActionMethod 方法
	resp, err := userFavoriteClient.BasicFavoriteActionMethod(ctx, req)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
