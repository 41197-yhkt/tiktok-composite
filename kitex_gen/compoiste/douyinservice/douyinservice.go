// Code generated by Kitex v0.4.4. DO NOT EDIT.

package douyinservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	compoiste "tiktok/kitex_gen/compoiste"
)

func serviceInfo() *kitex.ServiceInfo {
	return douyinServiceServiceInfo
}

var douyinServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "DouyinService"
	handlerType := (*compoiste.DouyinService)(nil)
	methods := map[string]kitex.MethodInfo{
		"BasicFavoriteActionMethod": kitex.NewMethodInfo(basicFavoriteActionMethodHandler, newDouyinServiceBasicFavoriteActionMethodArgs, newDouyinServiceBasicFavoriteActionMethodResult, false),
		"BasicFavoriteListMethod":   kitex.NewMethodInfo(basicFavoriteListMethodHandler, newDouyinServiceBasicFavoriteListMethodArgs, newDouyinServiceBasicFavoriteListMethodResult, false),
		"BasicFeedMethod":           kitex.NewMethodInfo(basicFeedMethodHandler, newDouyinServiceBasicFeedMethodArgs, newDouyinServiceBasicFeedMethodResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "compoiste",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.4.4",
		Extra:           extra,
	}
	return svcInfo
}

func basicFavoriteActionMethodHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*compoiste.DouyinServiceBasicFavoriteActionMethodArgs)
	realResult := result.(*compoiste.DouyinServiceBasicFavoriteActionMethodResult)
	success, err := handler.(compoiste.DouyinService).BasicFavoriteActionMethod(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newDouyinServiceBasicFavoriteActionMethodArgs() interface{} {
	return compoiste.NewDouyinServiceBasicFavoriteActionMethodArgs()
}

func newDouyinServiceBasicFavoriteActionMethodResult() interface{} {
	return compoiste.NewDouyinServiceBasicFavoriteActionMethodResult()
}

func basicFavoriteListMethodHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*compoiste.DouyinServiceBasicFavoriteListMethodArgs)
	realResult := result.(*compoiste.DouyinServiceBasicFavoriteListMethodResult)
	success, err := handler.(compoiste.DouyinService).BasicFavoriteListMethod(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newDouyinServiceBasicFavoriteListMethodArgs() interface{} {
	return compoiste.NewDouyinServiceBasicFavoriteListMethodArgs()
}

func newDouyinServiceBasicFavoriteListMethodResult() interface{} {
	return compoiste.NewDouyinServiceBasicFavoriteListMethodResult()
}

func basicFeedMethodHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*compoiste.DouyinServiceBasicFeedMethodArgs)
	realResult := result.(*compoiste.DouyinServiceBasicFeedMethodResult)
	success, err := handler.(compoiste.DouyinService).BasicFeedMethod(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newDouyinServiceBasicFeedMethodArgs() interface{} {
	return compoiste.NewDouyinServiceBasicFeedMethodArgs()
}

func newDouyinServiceBasicFeedMethodResult() interface{} {
	return compoiste.NewDouyinServiceBasicFeedMethodResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) BasicFavoriteActionMethod(ctx context.Context, req *compoiste.BasicFavoriteActionRequest) (r *compoiste.BasicFavoriteActionResponse, err error) {
	var _args compoiste.DouyinServiceBasicFavoriteActionMethodArgs
	_args.Req = req
	var _result compoiste.DouyinServiceBasicFavoriteActionMethodResult
	if err = p.c.Call(ctx, "BasicFavoriteActionMethod", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) BasicFavoriteListMethod(ctx context.Context, req *compoiste.BasicFavoriteListRequest) (r *compoiste.BasicFavoriteListResponse, err error) {
	var _args compoiste.DouyinServiceBasicFavoriteListMethodArgs
	_args.Req = req
	var _result compoiste.DouyinServiceBasicFavoriteListMethodResult
	if err = p.c.Call(ctx, "BasicFavoriteListMethod", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) BasicFeedMethod(ctx context.Context, req *compoiste.BasicFeedRequest) (r *compoiste.BasicFeedResponse, err error) {
	var _args compoiste.DouyinServiceBasicFeedMethodArgs
	_args.Req = req
	var _result compoiste.DouyinServiceBasicFeedMethodResult
	if err = p.c.Call(ctx, "BasicFeedMethod", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
