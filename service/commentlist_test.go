package service

import (
	"context"
	"testing"
	"tiktok-composite/kitex_gen/composite"
)

func TestCommentList(t *testing.T) {
	compositeClient := setupClient()
	ctx := context.Background()

	// 定义评论请求
	var req *composite.BasicCommentListRequest
	// 获取评论列表测试
	req = &composite.BasicCommentListRequest{
		VedioId: 1,
	}
	resp, err := compositeClient.BasicCommentListMethod(ctx, req)
	if err != nil {
		t.Error(err)
	}
	if resp.BaseResp.StatusCode != 0 {
		t.Errorf("Error Code: %v, Error Message: %v", resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	t.Log(resp.CommentList)
}
