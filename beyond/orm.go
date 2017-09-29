package beyond

import (
	"reflect"
	"fmt"
)

type IBaseDao interface {
	Init(model interface{})
}

type BaseDao struct {
	EntityType reflect.Type
	tableName     string            //表名
	pk            string            //主键
	columnToField map[string]string //字段名:属性名
	fieldToColumn map[string]string //属性名:字段名
}

func (dao *BaseDao) Init(model interface{})  {
	fmt.Println("init1")
	dao.columnToField = make(map[string]string)
	dao.fieldToColumn = make(map[string]string)

	dao.EntityType = reflect.TypeOf(model).Elem()
	types := dao.EntityType
	for i := 0; i < types.NumField(); i++ {
		typ := types.Field(i)
		tag := typ.Tag
		fmt.Println(tag)

		if len(tag) > 0 {
			column := tag.Get("column")
			name := typ.Name
			dao.columnToField[column] = name
			dao.fieldToColumn[name] = column

			if len(tag.Get("table")) > 0 {
				dao.tableName = tag.Get("table")
				dao.pk = column
			}
		}
	}
}

type Dao struct {
	BaseDao
}