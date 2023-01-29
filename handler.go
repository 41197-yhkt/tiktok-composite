package main

import (
	"context"
	composite "tiktok-composite/kitex_gen/composite"
)

// CompositeServiceImpl implements the last service interface defined in the IDL.
type CompositeServiceImpl struct{}

// BasicFavoriteActionMethod implements the CompositeServiceImpl interface.
func (s *CompositeServiceImpl) BasicFavoriteActionMethod(ctx context.Context, req *composite.BasicFavoriteActionRequest) (resp *composite.BasicFavoriteActionResponse, err error) {
	// TODO: Your code here...
	return
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
