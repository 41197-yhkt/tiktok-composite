package service

import (
	"context"
	"testing"
	"tiktok-composite/kitex_gen/composite"
)

func TestFeed(t *testing.T) {
	compositeClient := setupClient()
	ctx := context.Background()

	// 定义获取视频列表请求
	var req *composite.BasicFeedRequest
	// 获取视频列表测试
	// lastTime := time.Now().Unix()
	req = &composite.BasicFeedRequest{}
	resp, err := compositeClient.BasicFeedMethod(ctx, req)
	if err != nil {
		t.Error(err)
	}
	if resp.BaseResp.StatusCode != 0 {
		t.Errorf("Error Code: %v, Error Message: %v", resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	t.Logf("Favorite List: %v", resp.VedioList)
}
