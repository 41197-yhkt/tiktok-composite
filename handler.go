package main

import (
	"context"
	"tiktok-composite/kitex_gen/composite"
	"tiktok-composite/service"

	"tiktok-composite/pack"

	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/pkg/errno"
)

// CompositeServiceImpl implements the last service interface defined in the IDL.
type CompositeServiceImpl struct{}

// BasicFavoriteActionMethod implements the CompositeServiceImpl interface.
func (s *CompositeServiceImpl) BasicFavoriteActionMethod(ctx context.Context, req *composite.BasicFavoriteActionRequest) (resp *composite.BasicFavoriteActionResponse, err error) {
	resp = new(composite.BasicFavoriteActionResponse)

	// 检验参数是否符合规范
	if req.VedioId <= 0 || req.UserId <= 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	err = service.NewFavoriteActionService(ctx).FavoriteAction(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// BasicFavoriteListMethod implements the CompositeServiceImpl interface.
func (s *CompositeServiceImpl) BasicFavoriteListMethod(ctx context.Context, req *composite.BasicFavoriteListRequest) (resp *composite.BasicFavoriteListResponse, err error) {
	// TODO: Your code here...
	return
}

// BasicFeedMethod implements the CompositeServiceImpl interface.
func (s *CompositeServiceImpl) BasicFeedMethod(ctx context.Context, req *composite.BasicFeedRequest) (resp *composite.BasicFeedResponse, err error) {
	// TODO: Your code here...
	return
}
