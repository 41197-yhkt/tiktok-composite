package service

import (
	"context"
	"testing"
	"tiktok-composite/kitex_gen/composite"
)

func TestFavoriteAction(t *testing.T) {
	compositeClient := setupClient()
	ctx := context.Background()

	// 定义点赞请求
	var req *composite.BasicFavoriteActionRequest
	// 点赞测试
	req = &composite.BasicFavoriteActionRequest{VedioId: 5, UserId: 1}
	resp, err := compositeClient.BasicFavoriteActionMethod(ctx, req)
	if err != nil {
		t.Error(err)
	}
	if resp.BaseResp.StatusCode != 0 {
		t.Errorf("Error Code: %v, Error Message: %v", resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
}
