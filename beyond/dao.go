package beyond

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	//"reflect"
	//"strconv"
)

type Dao struct {
	db *sql.DB
	err error
}

func (dao *Dao) Connect() (*sql.DB, error) {
	dao.db, dao.err = sql.Open("mysql", "root:mysql@tcp(127.0.0.1:3306)/jserver?charset=utf8")

	return dao.db, dao.err
}

func (dao *Dao) Query(sqlSelect string, parseFunc func([]string, []interface{}) interface{}) []interface{} {
	//var err error
	countSql := "select count(*) from ( " + sqlSelect + " ) AS TT"
	fmt.Println(countSql)
	var count int64
	dao.db.QueryRow(countSql).Scan(&count)
	resultSet := make([]interface{}, count)

	rows, _ := dao.db.Query(sqlSelect)

	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	idx := 0
	for rows.Next() {
		rows.Scan(scanArgs...)
		//使用传入的函数，按照需要处理结果的形式
		obj := parseFunc(columns, values)
		resultSet[idx] = obj
		idx++
	}
	return resultSet
}