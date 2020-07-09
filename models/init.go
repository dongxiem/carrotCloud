package models

import (
	"carrotCloud/pkg/util"
	"github.com/jinzhu/gorm"
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB 数据库连接单例
var DB *gorm.DB

// Database 在中间件中初始化mysql链接
func Database(connString string) {
	// 1.打开数据库
	db, err := gorm.Open("mysql", connString)
	db.LogMode(true)
	if err != nil {
		util.Log().Panic("数据库连接不成功", err)
	}
	// 2.设置连接池
	// 2.1 空闲队列
	db.DB().SetMaxIdleConns(50)
	// 2.2 打开最大连接数
	db.DB().SetMaxOpenConns(100)
	// 2.3 超时设置
	db.DB().SetConnMaxLifetime(time.Second * 30)

	DB = db

	migration()
}
