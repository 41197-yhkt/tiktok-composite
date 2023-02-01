// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"tiktok-composite/gen/dal/model"
)

func newComment(db *gorm.DB, opts ...gen.DOOption) comment {
	_comment := comment{}

	_comment.commentDo.UseDB(db, opts...)
	_comment.commentDo.UseModel(&model.Comment{})

	tableName := _comment.commentDo.TableName()
	_comment.ALL = field.NewAsterisk(tableName)
	_comment.ID = field.NewUint(tableName, "id")
	_comment.CreatedAt = field.NewTime(tableName, "created_at")
	_comment.UpdatedAt = field.NewTime(tableName, "updated_at")
	_comment.DeletedAt = field.NewField(tableName, "deleted_at")
	_comment.UserId = field.NewInt64(tableName, "user_id")
	_comment.VedioId = field.NewInt64(tableName, "vedio_id")
	_comment.Content = field.NewString(tableName, "content")

	_comment.fillFieldMap()

	return _comment
}

type comment struct {
	commentDo commentDo

	ALL       field.Asterisk
	ID        field.Uint
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field
	UserId    field.Int64
	VedioId   field.Int64
	Content   field.String

	fieldMap map[string]field.Expr
}

func (c comment) Table(newTableName string) *comment {
	c.commentDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c comment) As(alias string) *comment {
	c.commentDo.DO = *(c.commentDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *comment) updateTableName(table string) *comment {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewUint(table, "id")
	c.CreatedAt = field.NewTime(table, "created_at")
	c.UpdatedAt = field.NewTime(table, "updated_at")
	c.DeletedAt = field.NewField(table, "deleted_at")
	c.UserId = field.NewInt64(table, "user_id")
	c.VedioId = field.NewInt64(table, "vedio_id")
	c.Content = field.NewString(table, "content")

	c.fillFieldMap()

	return c
}

func (c *comment) WithContext(ctx context.Context) *commentDo { return c.commentDo.WithContext(ctx) }

func (c comment) TableName() string { return c.commentDo.TableName() }

func (c comment) Alias() string { return c.commentDo.Alias() }

func (c *comment) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *comment) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 7)
	c.fieldMap["id"] = c.ID
	c.fieldMap["created_at"] = c.CreatedAt
	c.fieldMap["updated_at"] = c.UpdatedAt
	c.fieldMap["deleted_at"] = c.DeletedAt
	c.fieldMap["user_id"] = c.UserId
	c.fieldMap["vedio_id"] = c.VedioId
	c.fieldMap["content"] = c.Content
}

func (c comment) clone(db *gorm.DB) comment {
	c.commentDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c comment) replaceDB(db *gorm.DB) comment {
	c.commentDo.ReplaceDB(db)
	return c
}

type commentDo struct{ gen.DO }

// where(id=@id)
func (c commentDo) FindByID(id int64) (result model.Comment, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, id)
	generateSQL.WriteString("id=? ")

	var executeSQL *gorm.DB

	executeSQL = c.UnderlyingDB().Where(generateSQL.String(), params...).Take(&result)
	err = executeSQL.Error
	return
}

// where(user_id = @userId)
func (c commentDo) FindByUserid(userId int64) (result []model.Comment, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, userId)
	generateSQL.WriteString("user_id = ? ")

	var executeSQL *gorm.DB

	executeSQL = c.UnderlyingDB().Where(generateSQL.String(), params...).Find(&result)
	err = executeSQL.Error
	return
}

// where(vedio_id = @vedioId)
func (c commentDo) FindByVedioid(vedioId int64) (result []*model.Comment, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, vedioId)
	generateSQL.WriteString("vedio_id = ? ")

	var executeSQL *gorm.DB

	executeSQL = c.UnderlyingDB().Where(generateSQL.String(), params...).Find(&result)
	err = executeSQL.Error
	return
}

// where(user_id = @userId and vedio_id = @vedioId)
func (c commentDo) FindByUseridAndVedioid(userId int64, vedioId int64) (result []model.Comment, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, userId)
	params = append(params, vedioId)
	generateSQL.WriteString("user_id = ? and vedio_id = ? ")

	var executeSQL *gorm.DB

	executeSQL = c.UnderlyingDB().Where(generateSQL.String(), params...).Find(&result)
	err = executeSQL.Error
	return
}

// sql(delete from @@table where id = @id)
func (c commentDo) DeleteById(id int64) (err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, id)
	generateSQL.WriteString("delete from comments where id = ? ")

	var executeSQL *gorm.DB

	executeSQL = c.UnderlyingDB().Exec(generateSQL.String(), params...)
	err = executeSQL.Error
	return
}

// sql(select count(*) from @@table where vedio_id = @vedioId)
func (c commentDo) CountByVedioid(vedioId int64) (result int64, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, vedioId)
	generateSQL.WriteString("select count(*) from comments where vedio_id = ? ")

	var executeSQL *gorm.DB

	executeSQL = c.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result)
	err = executeSQL.Error
	return
}

// sql(SELECT LAST_INSERT_ID())
func (c commentDo) LastInsertID() (result uint, err error) {
	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT LAST_INSERT_ID() ")

	var executeSQL *gorm.DB

	executeSQL = c.UnderlyingDB().Raw(generateSQL.String()).Take(&result)
	err = executeSQL.Error
	return
}

func (c commentDo) Debug() *commentDo {
	return c.withDO(c.DO.Debug())
}

func (c commentDo) WithContext(ctx context.Context) *commentDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c commentDo) ReadDB() *commentDo {
	return c.Clauses(dbresolver.Read)
}

func (c commentDo) WriteDB() *commentDo {
	return c.Clauses(dbresolver.Write)
}

func (c commentDo) Session(config *gorm.Session) *commentDo {
	return c.withDO(c.DO.Session(config))
}

func (c commentDo) Clauses(conds ...clause.Expression) *commentDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c commentDo) Returning(value interface{}, columns ...string) *commentDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c commentDo) Not(conds ...gen.Condition) *commentDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c commentDo) Or(conds ...gen.Condition) *commentDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c commentDo) Select(conds ...field.Expr) *commentDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c commentDo) Where(conds ...gen.Condition) *commentDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c commentDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *commentDo {
	return c.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (c commentDo) Order(conds ...field.Expr) *commentDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c commentDo) Distinct(cols ...field.Expr) *commentDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c commentDo) Omit(cols ...field.Expr) *commentDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c commentDo) Join(table schema.Tabler, on ...field.Expr) *commentDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c commentDo) LeftJoin(table schema.Tabler, on ...field.Expr) *commentDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c commentDo) RightJoin(table schema.Tabler, on ...field.Expr) *commentDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c commentDo) Group(cols ...field.Expr) *commentDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c commentDo) Having(conds ...gen.Condition) *commentDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c commentDo) Limit(limit int) *commentDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c commentDo) Offset(offset int) *commentDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c commentDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *commentDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c commentDo) Unscoped() *commentDo {
	return c.withDO(c.DO.Unscoped())
}

func (c commentDo) Create(values ...*model.Comment) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c commentDo) CreateInBatches(values []*model.Comment, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c commentDo) Save(values ...*model.Comment) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c commentDo) First() (*model.Comment, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Comment), nil
	}
}

func (c commentDo) Take() (*model.Comment, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Comment), nil
	}
}

func (c commentDo) Last() (*model.Comment, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Comment), nil
	}
}

func (c commentDo) Find() ([]*model.Comment, error) {
	result, err := c.DO.Find()
	return result.([]*model.Comment), err
}

func (c commentDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Comment, err error) {
	buf := make([]*model.Comment, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c commentDo) FindInBatches(result *[]*model.Comment, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c commentDo) Attrs(attrs ...field.AssignExpr) *commentDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c commentDo) Assign(attrs ...field.AssignExpr) *commentDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c commentDo) Joins(fields ...field.RelationField) *commentDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c commentDo) Preload(fields ...field.RelationField) *commentDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c commentDo) FirstOrInit() (*model.Comment, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Comment), nil
	}
}

func (c commentDo) FirstOrCreate() (*model.Comment, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Comment), nil
	}
}

func (c commentDo) FindByPage(offset int, limit int) (result []*model.Comment, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c commentDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c commentDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c commentDo) Delete(models ...*model.Comment) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *commentDo) withDO(do gen.Dao) *commentDo {
	c.DO = *do.(*gen.DO)
	return c
}
