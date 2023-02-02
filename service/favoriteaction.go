package service

import (
	"context"

	"github.com/41197/tiktok-composite/gen/dal/model"
	"github.com/41197/tiktok-composite/kitex_gen/composite"
)

type FavoriteActionService struct {
	ctx context.Context
}

func NewFavoriteActionService(ctx context.Context) *FavoriteActionService {
	return &FavoriteActionService{ctx: ctx}
}

func (s *FavoriteActionService) FavoriteAction(req *composite.BasicFavoriteActionRequest) error {
	userFavoriteDatabase := q.UserFavorite.WithContext(s.ctx)

	// 1. 判断 ActionType
	if req.ActionType == 1 {
		// 2. 增加一条点赞记录
		return userFavoriteDatabase.Create(&model.UserFavorite{UserId: req.UserId, VedioId: req.VedioId})
	} else {
		// 3. 删除点赞记录
		return userFavoriteDatabase.DeleteByUseridAndVedioid(req.UserId, req.VedioId)
	}

}
