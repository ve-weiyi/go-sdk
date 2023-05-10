// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/ve-weiyi/go-sdk/utils/generate/blog/entity"
)

func newCOLUMNS(db *gorm.DB, opts ...gen.DOOption) cOLUMNS {
	_cOLUMNS := cOLUMNS{}

	_cOLUMNS.cOLUMNSDo.UseDB(db, opts...)
	_cOLUMNS.cOLUMNSDo.UseModel(&entity.COLUMNS{})

	tableName := _cOLUMNS.cOLUMNSDo.TableName()
	_cOLUMNS.ALL = field.NewAsterisk(tableName)
	_cOLUMNS.TABLECATALOG = field.NewString(tableName, "TABLE_CATALOG")
	_cOLUMNS.TABLESCHEMA = field.NewString(tableName, "TABLE_SCHEMA")
	_cOLUMNS.TABLENAME = field.NewString(tableName, "TABLE_NAME")
	_cOLUMNS.COLUMNNAME = field.NewString(tableName, "COLUMN_NAME")
	_cOLUMNS.ORDINALPOSITION = field.NewUint(tableName, "ORDINAL_POSITION")
	_cOLUMNS.COLUMNDEFAULT = field.NewString(tableName, "COLUMN_DEFAULT")
	_cOLUMNS.ISNULLABLE = field.NewString(tableName, "IS_NULLABLE")
	_cOLUMNS.DATATYPE = field.NewString(tableName, "DATA_TYPE")
	_cOLUMNS.CHARACTERMAXIMUMLENGTH = field.NewInt(tableName, "CHARACTER_MAXIMUM_LENGTH")
	_cOLUMNS.CHARACTEROCTETLENGTH = field.NewInt(tableName, "CHARACTER_OCTET_LENGTH")
	_cOLUMNS.NUMERICPRECISION = field.NewUint(tableName, "NUMERIC_PRECISION")
	_cOLUMNS.NUMERICSCALE = field.NewUint(tableName, "NUMERIC_SCALE")
	_cOLUMNS.DATETIMEPRECISION = field.NewUint(tableName, "DATETIME_PRECISION")
	_cOLUMNS.CHARACTERSETNAME = field.NewString(tableName, "CHARACTER_SET_NAME")
	_cOLUMNS.COLLATIONNAME = field.NewString(tableName, "COLLATION_NAME")
	_cOLUMNS.COLUMNTYPE = field.NewString(tableName, "COLUMN_TYPE")
	_cOLUMNS.COLUMNKEY = field.NewString(tableName, "COLUMN_KEY")
	_cOLUMNS.EXTRA = field.NewString(tableName, "EXTRA")
	_cOLUMNS.PRIVILEGES = field.NewString(tableName, "PRIVILEGES")
	_cOLUMNS.COLUMNCOMMENT = field.NewString(tableName, "COLUMN_COMMENT")
	_cOLUMNS.GENERATIONEXPRESSION = field.NewString(tableName, "GENERATION_EXPRESSION")
	_cOLUMNS.SRSID = field.NewUint(tableName, "SRS_ID")

	_cOLUMNS.fillFieldMap()

	return _cOLUMNS
}

type cOLUMNS struct {
	cOLUMNSDo

	ALL                    field.Asterisk
	TABLECATALOG           field.String
	TABLESCHEMA            field.String
	TABLENAME              field.String
	COLUMNNAME             field.String
	ORDINALPOSITION        field.Uint
	COLUMNDEFAULT          field.String
	ISNULLABLE             field.String
	DATATYPE               field.String
	CHARACTERMAXIMUMLENGTH field.Int
	CHARACTEROCTETLENGTH   field.Int
	NUMERICPRECISION       field.Uint
	NUMERICSCALE           field.Uint
	DATETIMEPRECISION      field.Uint
	CHARACTERSETNAME       field.String
	COLLATIONNAME          field.String
	COLUMNTYPE             field.String
	COLUMNKEY              field.String
	EXTRA                  field.String
	PRIVILEGES             field.String
	COLUMNCOMMENT          field.String
	GENERATIONEXPRESSION   field.String
	SRSID                  field.Uint

	fieldMap map[string]field.Expr
}

func (c cOLUMNS) Table(newTableName string) *cOLUMNS {
	c.cOLUMNSDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c cOLUMNS) As(alias string) *cOLUMNS {
	c.cOLUMNSDo.DO = *(c.cOLUMNSDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *cOLUMNS) updateTableName(table string) *cOLUMNS {
	c.ALL = field.NewAsterisk(table)
	c.TABLECATALOG = field.NewString(table, "TABLE_CATALOG")
	c.TABLESCHEMA = field.NewString(table, "TABLE_SCHEMA")
	c.TABLENAME = field.NewString(table, "TABLE_NAME")
	c.COLUMNNAME = field.NewString(table, "COLUMN_NAME")
	c.ORDINALPOSITION = field.NewUint(table, "ORDINAL_POSITION")
	c.COLUMNDEFAULT = field.NewString(table, "COLUMN_DEFAULT")
	c.ISNULLABLE = field.NewString(table, "IS_NULLABLE")
	c.DATATYPE = field.NewString(table, "DATA_TYPE")
	c.CHARACTERMAXIMUMLENGTH = field.NewInt(table, "CHARACTER_MAXIMUM_LENGTH")
	c.CHARACTEROCTETLENGTH = field.NewInt(table, "CHARACTER_OCTET_LENGTH")
	c.NUMERICPRECISION = field.NewUint(table, "NUMERIC_PRECISION")
	c.NUMERICSCALE = field.NewUint(table, "NUMERIC_SCALE")
	c.DATETIMEPRECISION = field.NewUint(table, "DATETIME_PRECISION")
	c.CHARACTERSETNAME = field.NewString(table, "CHARACTER_SET_NAME")
	c.COLLATIONNAME = field.NewString(table, "COLLATION_NAME")
	c.COLUMNTYPE = field.NewString(table, "COLUMN_TYPE")
	c.COLUMNKEY = field.NewString(table, "COLUMN_KEY")
	c.EXTRA = field.NewString(table, "EXTRA")
	c.PRIVILEGES = field.NewString(table, "PRIVILEGES")
	c.COLUMNCOMMENT = field.NewString(table, "COLUMN_COMMENT")
	c.GENERATIONEXPRESSION = field.NewString(table, "GENERATION_EXPRESSION")
	c.SRSID = field.NewUint(table, "SRS_ID")

	c.fillFieldMap()

	return c
}

func (c *cOLUMNS) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *cOLUMNS) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 22)
	c.fieldMap["TABLE_CATALOG"] = c.TABLECATALOG
	c.fieldMap["TABLE_SCHEMA"] = c.TABLESCHEMA
	c.fieldMap["TABLE_NAME"] = c.TABLENAME
	c.fieldMap["COLUMN_NAME"] = c.COLUMNNAME
	c.fieldMap["ORDINAL_POSITION"] = c.ORDINALPOSITION
	c.fieldMap["COLUMN_DEFAULT"] = c.COLUMNDEFAULT
	c.fieldMap["IS_NULLABLE"] = c.ISNULLABLE
	c.fieldMap["DATA_TYPE"] = c.DATATYPE
	c.fieldMap["CHARACTER_MAXIMUM_LENGTH"] = c.CHARACTERMAXIMUMLENGTH
	c.fieldMap["CHARACTER_OCTET_LENGTH"] = c.CHARACTEROCTETLENGTH
	c.fieldMap["NUMERIC_PRECISION"] = c.NUMERICPRECISION
	c.fieldMap["NUMERIC_SCALE"] = c.NUMERICSCALE
	c.fieldMap["DATETIME_PRECISION"] = c.DATETIMEPRECISION
	c.fieldMap["CHARACTER_SET_NAME"] = c.CHARACTERSETNAME
	c.fieldMap["COLLATION_NAME"] = c.COLLATIONNAME
	c.fieldMap["COLUMN_TYPE"] = c.COLUMNTYPE
	c.fieldMap["COLUMN_KEY"] = c.COLUMNKEY
	c.fieldMap["EXTRA"] = c.EXTRA
	c.fieldMap["PRIVILEGES"] = c.PRIVILEGES
	c.fieldMap["COLUMN_COMMENT"] = c.COLUMNCOMMENT
	c.fieldMap["GENERATION_EXPRESSION"] = c.GENERATIONEXPRESSION
	c.fieldMap["SRS_ID"] = c.SRSID
}

func (c cOLUMNS) clone(db *gorm.DB) cOLUMNS {
	c.cOLUMNSDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c cOLUMNS) replaceDB(db *gorm.DB) cOLUMNS {
	c.cOLUMNSDo.ReplaceDB(db)
	return c
}

type cOLUMNSDo struct{ gen.DO }

type ICOLUMNSDo interface {
	gen.SubQuery
	Debug() ICOLUMNSDo
	WithContext(ctx context.Context) ICOLUMNSDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ICOLUMNSDo
	WriteDB() ICOLUMNSDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ICOLUMNSDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ICOLUMNSDo
	Not(conds ...gen.Condition) ICOLUMNSDo
	Or(conds ...gen.Condition) ICOLUMNSDo
	Select(conds ...field.Expr) ICOLUMNSDo
	Where(conds ...gen.Condition) ICOLUMNSDo
	Order(conds ...field.Expr) ICOLUMNSDo
	Distinct(cols ...field.Expr) ICOLUMNSDo
	Omit(cols ...field.Expr) ICOLUMNSDo
	Join(table schema.Tabler, on ...field.Expr) ICOLUMNSDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ICOLUMNSDo
	RightJoin(table schema.Tabler, on ...field.Expr) ICOLUMNSDo
	Group(cols ...field.Expr) ICOLUMNSDo
	Having(conds ...gen.Condition) ICOLUMNSDo
	Limit(limit int) ICOLUMNSDo
	Offset(offset int) ICOLUMNSDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ICOLUMNSDo
	Unscoped() ICOLUMNSDo
	Create(values ...*entity.COLUMNS) error
	CreateInBatches(values []*entity.COLUMNS, batchSize int) error
	Save(values ...*entity.COLUMNS) error
	First() (*entity.COLUMNS, error)
	Take() (*entity.COLUMNS, error)
	Last() (*entity.COLUMNS, error)
	Find() ([]*entity.COLUMNS, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.COLUMNS, err error)
	FindInBatches(result *[]*entity.COLUMNS, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*entity.COLUMNS) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ICOLUMNSDo
	Assign(attrs ...field.AssignExpr) ICOLUMNSDo
	Joins(fields ...field.RelationField) ICOLUMNSDo
	Preload(fields ...field.RelationField) ICOLUMNSDo
	FirstOrInit() (*entity.COLUMNS, error)
	FirstOrCreate() (*entity.COLUMNS, error)
	FindByPage(offset int, limit int) (result []*entity.COLUMNS, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ICOLUMNSDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (c cOLUMNSDo) Debug() ICOLUMNSDo {
	return c.withDO(c.DO.Debug())
}

func (c cOLUMNSDo) WithContext(ctx context.Context) ICOLUMNSDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c cOLUMNSDo) ReadDB() ICOLUMNSDo {
	return c.Clauses(dbresolver.Read)
}

func (c cOLUMNSDo) WriteDB() ICOLUMNSDo {
	return c.Clauses(dbresolver.Write)
}

func (c cOLUMNSDo) Session(config *gorm.Session) ICOLUMNSDo {
	return c.withDO(c.DO.Session(config))
}

func (c cOLUMNSDo) Clauses(conds ...clause.Expression) ICOLUMNSDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c cOLUMNSDo) Returning(value interface{}, columns ...string) ICOLUMNSDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c cOLUMNSDo) Not(conds ...gen.Condition) ICOLUMNSDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c cOLUMNSDo) Or(conds ...gen.Condition) ICOLUMNSDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c cOLUMNSDo) Select(conds ...field.Expr) ICOLUMNSDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c cOLUMNSDo) Where(conds ...gen.Condition) ICOLUMNSDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c cOLUMNSDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ICOLUMNSDo {
	return c.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (c cOLUMNSDo) Order(conds ...field.Expr) ICOLUMNSDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c cOLUMNSDo) Distinct(cols ...field.Expr) ICOLUMNSDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c cOLUMNSDo) Omit(cols ...field.Expr) ICOLUMNSDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c cOLUMNSDo) Join(table schema.Tabler, on ...field.Expr) ICOLUMNSDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c cOLUMNSDo) LeftJoin(table schema.Tabler, on ...field.Expr) ICOLUMNSDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c cOLUMNSDo) RightJoin(table schema.Tabler, on ...field.Expr) ICOLUMNSDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c cOLUMNSDo) Group(cols ...field.Expr) ICOLUMNSDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c cOLUMNSDo) Having(conds ...gen.Condition) ICOLUMNSDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c cOLUMNSDo) Limit(limit int) ICOLUMNSDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c cOLUMNSDo) Offset(offset int) ICOLUMNSDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c cOLUMNSDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ICOLUMNSDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c cOLUMNSDo) Unscoped() ICOLUMNSDo {
	return c.withDO(c.DO.Unscoped())
}

func (c cOLUMNSDo) Create(values ...*entity.COLUMNS) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c cOLUMNSDo) CreateInBatches(values []*entity.COLUMNS, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c cOLUMNSDo) Save(values ...*entity.COLUMNS) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c cOLUMNSDo) First() (*entity.COLUMNS, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.COLUMNS), nil
	}
}

func (c cOLUMNSDo) Take() (*entity.COLUMNS, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.COLUMNS), nil
	}
}

func (c cOLUMNSDo) Last() (*entity.COLUMNS, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.COLUMNS), nil
	}
}

func (c cOLUMNSDo) Find() ([]*entity.COLUMNS, error) {
	result, err := c.DO.Find()
	return result.([]*entity.COLUMNS), err
}

func (c cOLUMNSDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.COLUMNS, err error) {
	buf := make([]*entity.COLUMNS, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c cOLUMNSDo) FindInBatches(result *[]*entity.COLUMNS, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c cOLUMNSDo) Attrs(attrs ...field.AssignExpr) ICOLUMNSDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c cOLUMNSDo) Assign(attrs ...field.AssignExpr) ICOLUMNSDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c cOLUMNSDo) Joins(fields ...field.RelationField) ICOLUMNSDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c cOLUMNSDo) Preload(fields ...field.RelationField) ICOLUMNSDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c cOLUMNSDo) FirstOrInit() (*entity.COLUMNS, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.COLUMNS), nil
	}
}

func (c cOLUMNSDo) FirstOrCreate() (*entity.COLUMNS, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.COLUMNS), nil
	}
}

func (c cOLUMNSDo) FindByPage(offset int, limit int) (result []*entity.COLUMNS, count int64, err error) {
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

func (c cOLUMNSDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c cOLUMNSDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c cOLUMNSDo) Delete(models ...*entity.COLUMNS) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *cOLUMNSDo) withDO(do gen.Dao) *cOLUMNSDo {
	c.DO = *do.(*gen.DO)
	return c
}
