package model

import (
	"time"

	"gorm.io/gen"
)

type VedioMethod interface {
	//where(id=@id)
	FindByID(id int) (gen.T, error)

	//sql(select * from @@table where updated_at < @lastTime limit @limit)
	FindByUpdatedtime(lastTime time.Time, limit int) (gen.T, error)
}

type UserFavoriteMethod interface {
	//where(userId=@userId)
	FindByUserid(userId int) (gen.T, error)

	//where(vedioId=@vedioId)
	FindByVedioId(vedioId int) (gen.T, error)
}
