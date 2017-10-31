package beyond

import (
	"fmt"
	"bytes"
	"strings"
)

type Query struct {
	model *BaseModel
	db *Dao

	sqlSelect string
	sqlFrom string
	sqlWhere map[string]string
}

func (q *Query) SetModel(model *BaseModel)  {
	q.model = model

}

func (q *Query) Where(condition map[string]string) *Query  {
	q.sqlWhere = condition
	return q
}

func (q *Query) All() interface{} {
	var sql = ""
	sql = q.createSelect()
	sql += q.model.tableName
	sql += q.createWhere()

	fmt.Println(sql)

	q.db.Query(sql, q.model.EntityType)

	return ""
}

func (q *Query) createSelect() string {
	var sql = ""
	var buf bytes.Buffer
	for _, v := range q.model.columnToField {
		buf.WriteString(" " + v  + ", ")
	}
	sql = "select " + buf.String() + " from "
	sql = strings.Replace(sql, ",  from", " from", -1)
	return sql
}

func (q *Query) createWhere() string {
	var sql = ""

	var buf bytes.Buffer
	for k, v := range q.sqlWhere {
		buf.WriteString(" " + k + " = '" + v + "' and ")
	}
	buf.WriteString(",")
	sql = " where " + buf.String()
	sql = strings.Replace(sql, "and ,", "", -1)

	return sql
}
