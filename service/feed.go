package service

import (
	"context"
	"time"

	"github.com/41197/tiktok-composite/kitex_gen/composite"
	"github.com/41197/tiktok-composite/pack"
)

type FeedService struct {
	ctx context.Context
}

func NewFeedService(ctx context.Context) *FeedService {
	return &FeedService{ctx: ctx}
}

func (s *FeedService) Feed(req *composite.BasicFeedRequest) ([]*composite.Vedio, int64, error) {
	vedioDatabase := q.Vedio.WithContext(s.ctx)

	// 1. 查看 req.Last_time 是否存在，若存在使用 req.Last_time，否则使用 time.Now()
	lastest_time := time.Now()
	if req.LastestTime != nil {
		lastest_time = time.Unix(*req.LastestTime, 0)
	}

	// 2. 查找 vedio 表中更新时间小于 last_time 的最大 limit 个视频
	eligibleVedios, err := vedioDatabase.FindByUpdatedtime(lastest_time, 30)
	if err != nil {
		return nil, -1, err
	}
	// 3.1. vedio 表中查视频信息
	// TODO: 接入 vedio 服务
	vedioIds := pack.VedioIds(eligibleVedios)
	vedios := make([]*composite.Vedio, 0)
	for i := 0; i < len(vedioIds); i++ {
		vedios = append(vedios, &composite.Vedio{
			Id: vedioIds[i],
		})
	}

	// 3.2. user 表中查作者信息
	// TODO: 接入 user 服务
	authorIds := pack.VedioAuthorIds(eligibleVedios)
	authors := make([]*composite.User, len(vedioIds))
	for i := 0; i < len(authorIds); i++ {
		authors[i] = &composite.User{}
		authors[i].Id = authorIds[i]
	}

	res := pack.Vedios(vedios, authors)
	return res, eligibleVedios[0].UpdatedAt.Unix(), nil

}
