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

	//sql(select * from @@table where updated_at < @lastTime order by updated_at limit @limit)
	FindByUpdatedtime(lastTime time.Time, limit int) ([]gen.T, error)
}

type UserFavoriteMethod interface {
	//sql(select vedio_id from @@table where user_id = @userId)
	FindByUserid(userId int64) (gen.T, error)

	//sql(select user_id from @@table where vedio_id = @vedioId)
	FindByVedioid(vedioId int64) (gen.T, error)

	//sql(select * from @@table where vedio_id = @vedioId and user_id = @userId)
	FindByUseridAndVedioid(userId, vedioId int64) error

	//sql(select count(*) from @@table where vedio_id = @vedioId)
	CountByVedioid(vedioId int64) (int64, error)
}

type CommentMethod interface {
	//where(id=@id)
	FindByID(id int64) (gen.T, error)

	//where(user_id = @userId)
	FindByUserid(userId int64) (gen.T, error)

	//where(vedio_id = @vedioId)
	FindByVedioid(vedioId int64) (gen.T, error)

	//where(user_id = @userId and vedio_id = @vedioId)
	FindByUseridAndVedioid(userId, vedioId int64) (gen.T, error)

	// sql(delete from @@table where id = @id)
	DeleteById(id int64) error

	//sql(select count(*) from @@table where vedio_id = @vedioId)
	CountByVedioid(vedioId int64) (int64, error)

	//sql(SELECT LAST_INSERT_ID())
	LastInsertID() (uint, error)
}
