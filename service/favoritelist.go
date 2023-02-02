package service

import (
	"context"

	"github.com/41197/tiktok-composite/gen/dal/model"
	"github.com/41197/tiktok-composite/kitex_gen/composite"
	"github.com/41197/tiktok-composite/pack"

	"github.com/41197-yhkt/pkg/errno"
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
		return nil, errno.UserNotExist
	}

	// 2. 对于每个 vedio_id
	// TODO: 接到 user 和 vedio 服务上
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
