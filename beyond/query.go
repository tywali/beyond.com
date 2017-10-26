package beyond

import (
	"fmt"
)

type Query struct {
	model *BaseModel

	sqlSelect string
	sqlFrom string
	sqlWhere string
}

func (q *Query) SetModel(model *BaseModel)  {
	q.model = model

}

func (q *Query) Where(condition string) *Query  {
	q.sqlWhere = condition
	return q
}

func (q *Query) All() interface{} {
	sql := "select * from " + q.model.tableName + " where " + q.sqlWhere
	fmt.Println(sql)
	return ""
}


