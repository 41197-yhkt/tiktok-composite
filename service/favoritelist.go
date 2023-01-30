package service

import (
	"context"
	"fmt"
	"tiktok-composite/gen/dal"
	"tiktok-composite/gen/dal/model"
	"tiktok-composite/kitex_gen/composite"
)

type FavoriteListService struct {
	ctx context.Context
}

func NewFavoriteListService(ctx context.Context) *FavoriteListService {
	return &FavoriteListService{ctx: ctx}
}

func (s *FavoriteListService) FavoriteList(req *composite.BasicFavoriteListRequest) ([]*composite.Vedio, error) {
	vedioDatabase := q.Vedio.WithContext(s.ctx)
	userFavoriteDatabase := q.UserFavorite.WithContext(s.ctx)

	// 先根据 user_id 选出 vedios
	var userFavorites []*model.UserFavorite
	dal.DB.Select("vedio_id").Where("user_id = ?", req.QueryId).Find(&userFavorites)
	fmt.Println(userFavorites)

	// 根据 vedio_id 查 Vedio
	res := []*composite.Vedio{}
	for _, favorite := range userFavorites {
		vedio, err := vedioDatabase.FindByID(favorite.VedioId)
		if err != nil {
			panic(err)
		}

		// 查询点赞数目
		var favoriteCount int64
		dal.DB.Where("vedio_id = ?", favorite.VedioId).Count(&favoriteCount)

		// TODO: 查询评论数
		// var commentCount int 64

		// 查询自己是不是也点了赞
		var isFavorite bool
		err = userFavoriteDatabase.FindByUseridAndVedioid(req.UserId, favorite.VedioId)
		if err != nil {
			isFavorite = false
		} else {
			isFavorite = true
		}

		// TODO: 查询作者信息

		// 封装
		res = append(res, &composite.Vedio{
			Id: int64(favorite.VedioId),
			Author: &composite.User{
				Id: vedio.AuthorId,
			},
			PlayUrl:       vedio.PlayUrl,
			CoverUrl:      vedio.CoverUrl,
			FavoriteCount: favoriteCount,
			// TODO:
			CommentCount: 1,
			// 查询点赞视频肯定为 true 吧
			IsFavorite: isFavorite,
			Title:      vedio.Title,
		})
	}

	return res, nil
}
