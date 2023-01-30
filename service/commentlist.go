package service

import (
	"context"
	"fmt"
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

	fmt.Println("根据 vedio_id 选出 comments")
	// 根据 vedio_id 选出 comments
	var comments []*model.Comment
	dal.DB.Where("vedio_id = ?", req.VedioId).Find(&comments)
	fmt.Println("评论列表如下： ", comments)

	res := []*composite.Comment{}
	for _, comment := range comments {
		commentauthor, err := userDatabase.FindByID(comment.UserId)
		if err != nil {
			panic(err)
		}

		// TODO: FollowCount 和 FollowerCount 计算在 relationship 表中查
		var followCount, followerCount int64 = 1, 1

		res = append(res, &composite.Comment{
			Id: int64(comment.Id),
			User: &composite.User{
				Id:            commentauthor.Id,
				Name:          commentauthor.Name,
				FollowCount:   &followCount,
				FollowerCount: &followerCount,
				IsFollow:      true,
			},
			Content:    comment.Content,
			CreateDate: comment.CreatedAt.Format("01-02"),
		})
	}

	return res, nil
}
