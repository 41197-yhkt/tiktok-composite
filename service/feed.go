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

	lastest_time := time.Now()
	if req.LastestTime != nil {
		lastest_time = time.Unix(*req.LastestTime, 0)
	}

	vedios, err := vedioDatabase.FindByUpdatedtime(lastest_time, 30)
	if err != nil {
		return nil, -1, err
	}
	// fmt.Println(vedios)
	res := []*composite.Vedio{}
	for _, vedio := range vedios {
		// 根据 AuthorId 查作者信息
		author, err := userDatabase.FindByID(vedio.AuthorId)
		if err != nil {
			panic(err)
		}

		// TODO: 根据 AuthorId 查作者的粉丝和关注数
		// followCount, followercount := 1, 1

		// 根据 user_id 和 vedio_id 查是否点赞了
		var isFavorite bool
		err = userFavoriteDatabase.FindByUseridAndVedioid(vedio.AuthorId, int64(vedio.ID))
		if err == nil {
			isFavorite = true
		}

		res = append(res, &composite.Vedio{
			Id: int64(vedio.ID),
			Author: &composite.User{
				Id:   author.Id,
				Name: author.Name,
				// FollowCount: followCount,
				// FollowerCount: &followercount,
			},
			PlayUrl:    vedio.PlayUrl,
			CoverUrl:   vedio.CoverUrl,
			IsFavorite: isFavorite,
			Title:      vedio.Title,
		})
	}

	return res, vedios[0].UpdatedAt.Unix(), nil

}
