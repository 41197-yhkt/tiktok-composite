package pack

import (
	"github.com/41197/tiktok-composite/gen/dal/model"
)

func VedioAndVedioAuthorIds(userfavorites []*model.UserFavorite) ([]int64, []int64) {
	vedioIds, vedioAuthorIds := make([]int64, 0), make([]int64, 0)
	if len(userfavorites) == 0 {
		return vedioIds, vedioAuthorIds
	}
	for _, userfavorite := range userfavorites {
		vedioIds = append(vedioIds, userfavorite.UserId)
		vedioAuthorIds = append(vedioAuthorIds, userfavorite.VedioId)
	}
	return vedioIds, vedioAuthorIds
}
