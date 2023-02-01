package service

import (
	"context"
	"log"
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

	// 1. 判断 ActionType
	if req.ActionType == 1 {
		// 2.1. 创建 Comment 数据模型，插入 comments 表中
		commentData := &model.Comment{
			UserId:  req.UserId,
			VedioId: req.VedioId,
			Content: *req.CommentText,
		}

		err := commentDatabase.Create(commentData)
		if err != nil {
			return nil, err
		}

		// 2.2. 从 comments 拿到最后一个插入的记录
		lastComment, err := commentDatabase.Last()
		if err != nil {
			return nil, err
		}

		// 2.3. users 表中查评论作者的信息
		userInfo, err := userDatabase.FindByID(lastComment.UserId)
		if err != nil {
			return nil, err
		}

		// 3. 封装
		resInfo := &composite.Comment{
			Id: int64(lastComment.ID),
			User: &composite.User{
				Id: int64(userInfo.ID),
			},
			Content:    lastComment.Content,
			CreateDate: lastComment.CreatedAt.Format("01-02"),
		}
		return resInfo, nil
	} else {
		// 软删除
		comment := new(model.Comment)
		comment.ID = uint(*req.CommentId)
		resInfo, err := commentDatabase.Delete(comment)

		// 硬删除
		// err := commentDatabase.DeleteById(*req.CommentId)

		if err != nil {
			return nil, err
		}
		log.Println(resInfo)
		return nil, err
	}
}
