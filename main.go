package main

import (
	"log"
	"net"
	"tiktok-composite/gen/dal"
	composite "tiktok-composite/kitex_gen/composite/compositeservice"
	"tiktok-composite/service"

	trace "github.com/41197-yhkt/pkg/trace"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/pkg/constants"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	internal_opentracing "github.com/kitex-contrib/tracer-opentracing"
)

func Init() {
	dal.Init()
	service.Init()
	trace.InitJaeger("kitex-server")
}

func main() {
	r, err := etcd.NewEtcdRegistry([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8889")
	if err != nil {
		panic(err)
	}
	Init()
	svr := composite.NewServer(new(CompositeServiceImpl),
		server.WithSuite(internal_opentracing.NewDefaultServerSuite()),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "composite"}),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
	)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
