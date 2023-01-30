package model

import (
	"time"

	"gorm.io/gen"
)

type UserMethod interface {
	//where(id=@id)
	FindByID(id int64) (gen.T, error)
}

type VedioMethod interface {
	//where(id=@id)
	FindByID(id int64) (gen.T, error)

	//sql(select * from @@table where updated_at < @lastTime limit @limit)
	FindByUpdatedtime(lastTime time.Time, limit int) (gen.T, error)
}

type UserFavoriteMethod interface {
	//sql(select vedio_id from @@table where user_id = @userId)
	FindByUserid(userId int64) (gen.T, error)

	//sql(select user_id from @@table where vedio_id = @vedioId)
	FindByVedioid(vedioId int64) (gen.T, error)

	//sql(select * from @@table where vedio_id = @vedioId and user_id = @userId)
	FindByUseridAndVedioid(userId, vedioId int64) error
}
