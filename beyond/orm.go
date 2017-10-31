package beyond

import (
	"reflect"
	"fmt"
	"strings"
)

type BaseOrm struct {
	modelList map[string] *BaseModel
	db *Dao

	curModel *BaseModel
}

func (dao *BaseOrm) Init() {
	dao.modelList = make(map[string] *BaseModel)
	dao.db = new(Dao)
	dao.db.Connect()
}

func (dao *BaseOrm) Register(model interface{})  {
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

func (dao *BaseOrm) Model(name string) *BaseOrm {
	dao.curModel = dao.modelList[strings.ToLower(name)]
	return dao
}

func (dao *BaseOrm) Find() *Query {
	q := new(Query)
	q.model = dao.curModel
	q.db = dao.db
	return q
}

func (dao *BaseOrm) NewModel(name string) interface{} {
	return dao.modelList[strings.ToLower(name)].Entity
}

func (dao *BaseOrm) Save(model interface{})  {
	fmt.Println("save")
}
