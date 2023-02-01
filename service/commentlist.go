package service

import (
	"context"
	"tiktok-composite/gen/dal"
	"tiktok-composite/gen/dal/model"
	"tiktok-composite/kitex_gen/composite"
)

type CommentListService struct {
	ctx context.Context
}

func NewCommentListService(ctx context.Context) *CommentListService {
	return &CommentListService{ctx: ctx}
}

func (s *CommentListService) CommentList(req *composite.BasicCommentListRequest) ([]*composite.Comment, error) {
	userDatabase := q.User.WithContext(s.ctx)

	// 1. 从 comments 中根据 vedio_id 查 comments
	var comments []*model.Comment
	dal.DB.Where("vedio_id = ?", req.VedioId).Find(&comments)

	// 2. 对于每个 comments 中的 user_id
	res := []*composite.Comment{}
	for _, comment := range comments {
		// 2.1. users 表中查 Name、 follow_count 和 follower_count
		commentauthor, err := userDatabase.FindByID(comment.UserId)
		if err != nil {
			panic(err)
		}

		// 2.2. relationship 表中查 user_id 和 author_id 的关注关系
		// var isFollow bool

		// 3. 封装
		res = append(res, &composite.Comment{
			Id: int64(comment.ID),
			User: &composite.User{
				Id:            int64(commentauthor.ID),
				Name:          commentauthor.Name,
				FollowCount:   &commentauthor.FollowCount,
				FollowerCount: &commentauthor.FollowerCount,
				IsFollow:      true,
			},
			Content:    comment.Content,
			CreateDate: comment.CreatedAt.Format("01-02"),
		})
	}

	return res, nil
}
