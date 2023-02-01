package pack

import (
	"tiktok-composite/gen/dal/model"
	"tiktok-composite/kitex_gen/composite"
)

// 打包 comment
func Comment(comment *model.Comment, author *model.User, isFollow bool) *composite.Comment {
	if comment == nil {
		return nil
	}
	if author == nil {
		return &composite.Comment{
			Id:         int64(comment.ID),
			User:       nil,
			Content:    comment.Content,
			CreateDate: comment.CreatedAt.Format("01-02"),
		}
	}

	return &composite.Comment{
		Id: int64(comment.ID),
		User: &composite.User{
			Id:            int64(author.ID),
			Name:          author.Name,
			FollowCount:   &author.FollowCount,
			FollowerCount: &author.FollowerCount,
			IsFollow:      isFollow,
		},
		Content:    comment.Content,
		CreateDate: comment.CreatedAt.Format("01-02"),
	}
}

// 打包 comment list
func Comments(comments []*model.Comment, authors []*model.User, isFollows []bool) []*composite.Comment {
	res := make([]*composite.Comment, 0)
	for i := 0; i < len(comments); i++ {
		if c := Comment(comments[i], authors[i], isFollows[i]); c != nil {
			res = append(res, c)
		}
	}
	return res
}

// 根据 comment 返回 author list 帮助查询
func AuthorIds(comments []*model.Comment) []int64 {
	res := make([]int64, 0)
	if len(comments) == 0 {
		return res
	}
	for _, comment := range comments {
		res = append(res, comment.UserId)
	}
	return res
}
