package main

import (
	"context"
	"tiktok-composite/kitex_gen/composite"
	"tiktok-composite/service"

	"tiktok-composite/pack"

	"github.com/41197-yhkt/pkg/errno"
)

// CompositeServiceImpl implements the last service interface defined in the IDL.
type CompositeServiceImpl struct{}

// BasicFavoriteActionMethod implements the CompositeServiceImpl interface.
func (s *CompositeServiceImpl) BasicFavoriteActionMethod(ctx context.Context, req *composite.BasicFavoriteActionRequest) (resp *composite.BasicFavoriteActionResponse, err error) {
	resp = new(composite.BasicFavoriteActionResponse)

	// 检验参数是否符合规范
	if req.VedioId <= 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.InvalidParams)
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
	resp = new(composite.BasicFavoriteListResponse)

	// 检验参数是否符合规范
	if req.UserId <= 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.UserNotExist)
		return resp, nil
	}

	resp.VedioList, err = service.NewFavoriteListService(ctx).FavoriteList(req)
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return
}

// BasicFeedMethod implements the CompositeServiceImpl interface.
func (s *CompositeServiceImpl) BasicFeedMethod(ctx context.Context, req *composite.BasicFeedRequest) (resp *composite.BasicFeedResponse, err error) {
	resp = new(composite.BasicFeedResponse)

	vedioList, nextTime, err := service.NewFeedService(ctx).Feed(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return nil, err
	}

	resp.VedioList = vedioList
	resp.NextTime = &nextTime
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// BasicCommentActionMethod implements the CompositeServiceImpl interface.
func (s *CompositeServiceImpl) BasicCommentActionMethod(ctx context.Context, req *composite.BasicCommentActionRequest) (resp *composite.BasicCommentActionResponse, err error) {
	resp = new(composite.BasicCommentActionResponse)

	// 检验参数是否符合规范
	if (req.ActionType == 1 && (req.VedioId <= 0 || req.UserId <= 0)) || (req.ActionType == 2 && (*req.CommentId <= 0)) {
		resp.BaseResp = pack.BuildBaseResp(errno.InvalidParams)
		return resp, nil
	}

	resp.Comment, err = service.NewCommentActionService(ctx).CommentAction(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// BasicCommentListMethod implements the CompositeServiceImpl interface.
func (s *CompositeServiceImpl) BasicCommentListMethod(ctx context.Context, req *composite.BasicCommentListRequest) (resp *composite.BasicCommentListResponse, err error) {
	resp = new(composite.BasicCommentListResponse)

	if req.VedioId <= 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.InvalidParams)
		return resp, nil
	}

	resp.CommentList, err = service.NewCommentListService(ctx).CommentList(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}
