package initialize

import (
	"fmt"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/militch/exswap/global"
)

// MysqlDb mysql结构体
func Gorm() *gorm.DB {
	db, err := gorm.Open("mysql", global.GVA_CONFIG.Mysql.Dsn())
	// db.SetLogger(true)
	db.LogMode(false)
	if err != nil {
		fmt.Printf("gorm err:%v\n", err)
		os.Exit(1)
	}
	ConnMaxLifetime, err := time.ParseDuration(global.GVA_CONFIG.Mysql.ConnMaxLifetime)
	if err != nil {
		fmt.Printf("gorm ConnMaxLifetime err:%v\n", err)
		os.Exit(1)
	}
	db.DB().SetMaxIdleConns(global.GVA_CONFIG.Mysql.MaxIdleConns)
	db.DB().SetMaxOpenConns(global.GVA_CONFIG.Mysql.MaxOpenConns)
	db.DB().SetConnMaxLifetime(ConnMaxLifetime)
	// defer db.Close()
	db.SingularTable(true)
	// if global.GVA_GVA_CONFIG.Mysql.Debug {
	// 	return db.Debug()
	// }
	return db
}
