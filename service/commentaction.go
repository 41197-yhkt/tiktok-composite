package service

import (
	"context"
	"tiktok-composite/gen/dal/model"
	"tiktok-composite/kitex_gen/composite"
)

type CommentActionService struct {
	ctx context.Context
}

func NewCommentActionService(ctx context.Context) *CommentActionService {
	return &CommentActionService{ctx: ctx}
}

func (s *CommentActionService) CommentAction(req *composite.BasicCommentActionRequest) (*composite.Comment, error) {
	commentDatabase := q.Comment.WithContext(s.ctx)
	userDatabase := q.User.WithContext(s.ctx)

	if req.ActionType == 1 {
		commentData := &model.Comment{
			UserId:  req.UserId,
			VedioId: req.VedioId,
			Content: *req.CommentText,
		}
		err := commentDatabase.Create(commentData)
		if err != nil {
			return nil, err
		}
		// TODO: 这里高并发可能出现问题..
		lastComment, err := commentDatabase.Last()
		if err != nil {
			return nil, err
		}
		userInfo, err := userDatabase.FindByID(lastComment.UserId)
		if err != nil {
			return nil, err
		}
		resInfo := &composite.Comment{
			Id: int64(lastComment.Id),
			User: &composite.User{
				Id: userInfo.Id,
			},
			Content:    lastComment.Content,
			CreateDate: lastComment.CreatedAt.Format("01-02"),
		}
		return resInfo, nil
	} else {
		// 硬删除
		err := commentDatabase.DeleteById(*req.CommentId)
		if err != nil {
			return nil, err
		}
		return nil, err
	}
}
