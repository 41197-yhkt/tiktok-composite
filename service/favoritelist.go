package service

import (
	"context"
	"tiktok-composite/gen/dal/model"
	"tiktok-composite/kitex_gen/composite"
	"tiktok-composite/pack"
)

type FavoriteListService struct {
	ctx context.Context
}

func NewFavoriteListService(ctx context.Context) *FavoriteListService {
	return &FavoriteListService{ctx: ctx}
}

func (s *FavoriteListService) FavoriteList(req *composite.BasicFavoriteListRequest) ([]*composite.Vedio, error) {
	userFavoriteDatabase := q.UserFavorite.WithContext(s.ctx)

	// 1. 从 user_favorites 中根据 query_id 查 vedio_id
	var userFavorites []*model.UserFavorite
	userFavorites, err := userFavoriteDatabase.FindByUserid(req.QueryId)
	if err != nil {
		panic(err)
	}

	// 2. 对于每个 vedio_id
	vedioIds, authorIds := pack.VedioAndVedioAuthorIds(userFavorites)
	vedios := make([]*composite.Vedio, 0)
	authors := make([]*composite.User, 0)
	for i := 0; i < len(vedioIds); i++ {
		vedios = append(vedios, &composite.Vedio{
			Id: vedioIds[i],
		})
		authors = append(authors, &composite.User{
			Id: authorIds[i],
		})
	}
	res := pack.Vedios(vedios, authors)
	return res, nil
}
