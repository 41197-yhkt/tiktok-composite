package main

import (
	"context"
	compoiste "tiktok/kitex_gen/compoiste"
)

// DouyinServiceImpl implements the last service interface defined in the IDL.
type DouyinServiceImpl struct{}

// BasicFavoriteActionMethod implements the DouyinServiceImpl interface.
func (s *DouyinServiceImpl) BasicFavoriteActionMethod(ctx context.Context, req *compoiste.BasicFavoriteActionRequest) (resp *compoiste.BasicFavoriteActionResponse, err error) {
	// TODO: Your code here...
	return
}

// BasicFavoriteListMethod implements the DouyinServiceImpl interface.
func (s *DouyinServiceImpl) BasicFavoriteListMethod(ctx context.Context, req *compoiste.BasicFavoriteListRequest) (resp *compoiste.BasicFavoriteListResponse, err error) {
	// TODO: Your code here...
	return
}

// BasicFeedMethod implements the DouyinServiceImpl interface.
func (s *DouyinServiceImpl) BasicFeedMethod(ctx context.Context, req *compoiste.BasicFeedRequest) (resp *compoiste.BasicFeedResponse, err error) {
	// TODO: Your code here...
	return
}
