package service

import (
	"context"
	"tiktok-composite/gen/dal/model"
	"tiktok-composite/kitex_gen/composite"
)

type FavoriteActionService struct {
	ctx context.Context
}

func NewFavoriteActionService(ctx context.Context) *FavoriteActionService {
	return &FavoriteActionService{ctx: ctx}
}

func (s *FavoriteActionService) FavoriteAction(req *composite.BasicFavoriteActionRequest) error {
	userFavoriteDatabase := q.UserFavorite.WithContext(s.ctx)

	userFavoriteData := &model.UserFavorite{UserId: req.UserId, VedioId: req.VedioId}

	return userFavoriteDatabase.Create(userFavoriteData)
}
