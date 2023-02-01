package pack

import (
	"tiktok-composite/gen/dal/model"
	"tiktok-composite/kitex_gen/composite"
)

// 打包 vedio
// TODO: author 不应该是 *model.User 而是 *User.User
func Vedio(vedio *composite.Vedio, author *composite.User) *composite.Vedio {
	if vedio == nil || author == nil {
		return nil
	}

	return &composite.Vedio{
		Id: vedio.Id,
		Author: &composite.User{
			Id:            int64(author.Id),
			Name:          author.Name,
			FollowCount:   author.FollowCount,
			FollowerCount: author.FollowerCount,
			// TODO: 这里也不应该自己传参，后续从 user 那直接拿
			IsFollow: false,
		},
		PlayUrl:  vedio.PlayUrl,
		CoverUrl: vedio.CoverUrl,
		Title:    vedio.Title,
	}
}

// 打包 vedio list
// TODO: authors 不应该是 []*model.User 而是 []*User.User
func Vedios(vedios []*composite.Vedio, authors []*composite.User) []*composite.Vedio {
	res := make([]*composite.Vedio, 0)
	for i := 0; i < len(vedios); i++ {
		if v := Vedio(vedios[i], authors[i]); v != nil {
			res = append(res, v)
		}
	}
	return res
}

func VedioAuthorIds(vedios []*model.Vedio) []int64 {
	res := make([]int64, 0)
	for _, v := range vedios {
		if v != nil {
			res = append(res, v.AuthorId)
		}
	}
	return res
}

func VedioIds(vedios []*model.Vedio) []int64 {
	res := make([]int64, 0)
	for _, v := range vedios {
		if v != nil {
			res = append(res, int64(v.ID))
		}
	}
	return res
}
