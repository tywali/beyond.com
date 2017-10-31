package beyond

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"reflect"
	"strconv"
)

type Dao struct {
	db *sql.DB
	err error
}

func (dao *Dao) Connect() (*sql.DB, error) {
	dao.db, dao.err = sql.Open("mysql", "root:mysql@tcp(127.0.0.1:3306)/jserver?charset=utf8")

	return dao.db, dao.err
}

func (dao *Dao) Query(sqlSelect string, entityType reflect.Type) {
	//var err error
	rows, _ := dao.db.Query(sqlSelect)

	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		//将行数据保存到record字典
		rows.Scan(scanArgs...)
		obj := reflect.New(entityType).Interface()
		typ := reflect.ValueOf(obj).Elem()

		fmt.Println(obj, typ)
		//record := make(map[string]string)
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
				//record[columns[i]] = string(col.([]byte))
			}
		}
		fmt.Println(obj)
		//fmt.Println(record)
	}
}