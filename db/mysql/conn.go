package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	cfg "carrotCloud/config"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// init 数据库初始化
func init() {
	db, _ = sql.Open("mysql", cfg.MySQLSource)
	db.SetMaxOpenConns(1000)
	err := db.Ping()
	if err != nil {
		fmt.Println("Failed to connect mysql, err:" + err.Error())
		os.Exit(1)
	}
}

// DBConn：得到一个数据库连接对象
func DBConn() *sql.DB {
	return db
}

// ParseRows：解析每行数据
func ParseRows(rows *sql.Rows) []map[string]interface{} {
	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for j := range values {
		scanArgs[j] = &values[j]
	}

	record := make(map[string]interface{})
	records := make([]map[string]interface{}, 0)
	// 如果下一行存在则不断遍历
	for rows.Next() {
		err := rows.Scan(scanArgs...)
		checkErr(err)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = col
			}
		}
		// 将每条record都插入records
		records = append(records, record)
	}
	return records
}

// checkErr：检查错误，如果存在错误则打印
func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}
