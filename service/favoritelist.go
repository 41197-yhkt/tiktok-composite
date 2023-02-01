package service

import (
	"context"
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
	userDatabase := q.User.WithContext(s.ctx)
	vedioDatabase := q.Vedio.WithContext(s.ctx)
	comentDatabase := q.Comment.WithContext(s.ctx)
	userFavoriteDatabase := q.UserFavorite.WithContext(s.ctx)

	// 1. 从 user_favorites 中根据 query_id 查 vedio_id
	var userFavorites []*model.UserFavorite
	userFavorites, err := userFavoriteDatabase.FindByUserid(req.QueryId)
	if err != nil {
		panic(err)
	}
	// dal.DB.Select("vedio_id").Where("user_id = ?", req.QueryId).Find(&userFavorites)

	// 2. 对于每个 vedio_id
	res := []*composite.Vedio{}
	for _, favorite := range userFavorites {
		// 2.1. vedios 表中查 vedio 具体信息
		vedio, err := vedioDatabase.FindByID(favorite.VedioId)
		if err != nil {
			panic(err)
		}

		// 2.2. comments  表中查评论数
		commentCount, err := comentDatabase.CountByVedioid(favorite.VedioId)
		if err != nil {
			panic(err)
		}

		// 2.3. user_favorites 表中查点赞数
		favoriteCount, err := userFavoriteDatabase.CountByVedioid(favorite.VedioId)
		if err != nil {
			panic(err)
		}
		// var favoriteCount int64
		// dal.DB.Where("vedio_id = ?", favorite.VedioId).Count(&favoriteCount)

		// 2.4. user_favorites 表中查是不是自己点赞了
		// TODO: 这里可以用 EXISTS 优化嘛？
		var isFavorite bool
		err = userFavoriteDatabase.FindByUseridAndVedioid(req.UserId, favorite.VedioId)
		if err != nil {
			isFavorite = false
		} else {
			isFavorite = true
		}

		// 2.5. users 中获取完整 author 信息
		author, err := userDatabase.FindByID(vedio.AuthorId)
		if err != nil {
			panic(err)
		}

		// 3. 封装
		res = append(res, &composite.Vedio{
			Id: int64(favorite.VedioId),
			Author: &composite.User{
				Id:            int64(author.ID),
				FollowCount:   &author.FollowCount,
				FollowerCount: &author.FollowerCount,
			},
			PlayUrl:       vedio.PlayUrl,
			CoverUrl:      vedio.CoverUrl,
			FavoriteCount: favoriteCount,
			CommentCount:  commentCount,
			IsFavorite:    isFavorite,
			Title:         vedio.Title,
		})
	}

	return res, nil
}
