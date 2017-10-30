package beyond

import (
	"reflect"
)

type BaseModel struct {
	EntityType reflect.Type
	tableName     string            //表名
	pk            string            //主键
	columnToField map[string]string //字段名:属性名
	fieldToColumn map[string]string //属性名:字段名
	Entity interface{}
}

