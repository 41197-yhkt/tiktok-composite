package service

import (
	"context"
	"tiktok-composite/kitex_gen/composite"
	"time"
)

type FeedService struct {
	ctx context.Context
}

func NewFeedService(ctx context.Context) *FeedService {
	return &FeedService{ctx: ctx}
}

func (s *FeedService) Feed(req *composite.BasicFeedRequest) ([]*composite.Vedio, int64, error) {
	vedioDatabase := q.Vedio.WithContext(s.ctx)
	userDatabase := q.User.WithContext(s.ctx)
	userFavoriteDatabase := q.UserFavorite.WithContext(s.ctx)

	// 1. 查看 req.Last_time 是否存在，若存在使用 req.Last_time，否则使用 time.Now()
	lastest_time := time.Now()
	if req.LastestTime != nil {
		lastest_time = time.Unix(*req.LastestTime, 0)
	}

	// 2. 查找 vedio 表中更新时间小于 last_time 的最大 limit 个视频
	vedios, err := vedioDatabase.FindByUpdatedtime(lastest_time, 30)
	if err != nil {
		return nil, -1, err
	}

	// 3. 根据每个 vedio
	res := []*composite.Vedio{}
	for _, vedio := range vedios {
		// 3.1. user 表中根据 AuthorId 查作者信息
		author, err := userDatabase.FindByID(vedio.AuthorId)
		if err != nil {
			panic(err)
		}

		// 3.2. user_favorite 表中根据 user_id 和 vedio_id 查是否点赞了
		// TODO:  可以用 EXISTS 优化？
		var isFavorite bool
		err = userFavoriteDatabase.FindByUseridAndVedioid(vedio.AuthorId, int64(vedio.ID))
		if err == nil {
			isFavorite = true
		}

		// 4. 封装
		res = append(res, &composite.Vedio{
			Id: int64(vedio.ID),
			Author: &composite.User{
				Id:            int64(author.ID),
				Name:          author.Name,
				FollowCount:   &author.FollowCount,
				FollowerCount: &author.FollowerCount,
			},
			PlayUrl:    vedio.PlayUrl,
			CoverUrl:   vedio.CoverUrl,
			IsFavorite: isFavorite,
			Title:      vedio.Title,
		})
	}

	return res, vedios[0].UpdatedAt.Unix(), nil

}
