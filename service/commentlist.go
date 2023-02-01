package service

import (
	"context"
	"tiktok-composite/gen/dal/model"
	"tiktok-composite/kitex_gen/composite"
	"tiktok-composite/pack"
)

type CommentListService struct {
	ctx context.Context
}

func NewCommentListService(ctx context.Context) *CommentListService {
	return &CommentListService{ctx: ctx}
}

func (s *CommentListService) CommentList(req *composite.BasicCommentListRequest) ([]*composite.Comment, error) {
	commentDatabase := q.Comment.WithContext(s.ctx)

	// 1. 从 comments 中根据 vedio_id 查 comments
	var comments []*model.Comment
	comments, err := commentDatabase.FindByVedioid(req.VedioId)
	if err != nil {
		return nil, err
	}

	// 2. 对于每个 comments 中的 user_id
	// 2.1. users 表中查信息
	// TODO: 后续接到 user 服务上
	authorIds := pack.AuthorIds(comments)
	authors := make([]*model.User, len(authorIds))
	for i := 0; i < len(authorIds); i++ {
		authors[i] = &model.User{}
		authors[i].ID = uint(authorIds[i])
	}

	// 2.2 relationship 表中查 user_id 和 author_id 的关注关系
	isFollows := make([]bool, len(authorIds))

	res := pack.Comments(comments, authors, isFollows)

	return res, nil
}
