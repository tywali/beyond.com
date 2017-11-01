package beyond

import (
	"fmt"
	"bytes"
	"strings"
	"reflect"
	"strconv"
)

type Query struct {
	model *BaseModel
	db *Dao

	asArray bool
	sqlSelect string
	sqlFrom string
	sqlWhere map[string]string
}

func (q *Query) SetModel(model *BaseModel)  {
	q.model = model
	q.asArray = false
}

func (q *Query) Where(condition map[string]string) *Query  {
	q.sqlWhere = condition
	return q
}

func (q *Query) AsArray() *Query {
	q.asArray = true
	return q
}

func (q *Query) All() []interface{} {
	var sql = ""
	sql = q.createSelect()
	sql += q.model.tableName
	sql += q.createWhere()

	fmt.Println(sql)

	var result []interface{}
	if q.asArray {
		result = q.db.Query(sql, q.dataToArray)
	} else {
		result = q.db.Query(sql, q.dataToModel)
	}
	return result
}

func (q *Query) dataToModel(columns []string, values []interface{}) interface{} {
	entityType := q.model.EntityType
	obj := reflect.New(entityType).Interface()
	typ := reflect.ValueOf(obj).Elem()

	for i, col := range values {
		if col != nil {
			field := typ.FieldByName(columns[i])
			sVal := string(col.([]byte))
			switch field.Kind() {
			case reflect.String:
				field.SetString(sVal)
			case reflect.Int:
				v, _ := strconv.ParseInt(sVal, 10, 0)
				field.SetInt(v)
			}
		}
	}
	return obj
}

func (q *Query) dataToArray(columns []string, values []interface{}) interface{} {
	item := make(map[string]string)
	for i, col := range values {
		item[columns[i]] = string(col.([]byte))
	}
	return item
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

	if len(q.sqlWhere) > 0 {
		var buf bytes.Buffer
		for k, v := range q.sqlWhere {
			buf.WriteString(" " + k + " = '" + v + "' and ")
		}
		buf.WriteString(",")
		sql = " where " + buf.String()
		sql = strings.Replace(sql, "and ,", "", -1)
	}

	return sql
}
