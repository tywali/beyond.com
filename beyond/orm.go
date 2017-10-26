package beyond

import (
	"reflect"
	"fmt"
	"strings"
)

type BaseModel struct {
	EntityType reflect.Type
	tableName     string            //表名
	pk            string            //主键
	columnToField map[string]string //字段名:属性名
	fieldToColumn map[string]string //属性名:字段名
	Entity interface{}
}

type IBaseDao interface {
	Init()
	Register(model interface{})
}

type BaseDao struct {
	modelList map[string] *BaseModel

	curModel *BaseModel
}

func (dao *BaseDao) Init() {
	dao.modelList = make(map[string] *BaseModel)
}

func (dao *BaseDao) Register(model interface{})  {
	modelType := reflect.TypeOf(model)
	str := strings.Split(fmt.Sprintf("%s", modelType), ".")
	modelName := str[len(str) - 1]

	m := new(BaseModel)
	m.Entity = model
	m.columnToField = make(map[string]string)
	m.fieldToColumn = make(map[string]string)
	m.EntityType = reflect.TypeOf(model).Elem()
	types := m.EntityType
	for i := 0; i < types.NumField(); i++ {
		typ := types.Field(i)
		tag := typ.Tag
		fmt.Println(tag)

		if len(tag) > 0 {
			column := tag.Get("column")
			name := typ.Name
			m.columnToField[column] = name
			m.fieldToColumn[name] = column

			if len(tag.Get("table")) > 0 {
				m.tableName = tag.Get("table")
				m.pk = column
			}
		}
	}

	dao.modelList[strings.ToLower(modelName)] = m
}

func (dao *BaseDao) Model(name string) *BaseDao {
	dao.curModel = dao.modelList[strings.ToLower(name)]
	return dao
}

func (dao *BaseDao) Find() *Query {
	q := new(Query)
	q.model = dao.curModel
	return q
}

func (dao *BaseDao) NewModel(name string) interface{} {
	return dao.modelList[strings.ToLower(name)].Entity
}

func (dao *BaseDao) Save(model interface{})  {
	fmt.Println("save")
}
