package initialize

import (
	"github.com/CodingJzy/library_system/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

// 连接mysql
func Mysql() *gorm.DB {
	// 读取配置
	m := global.Config.Mysql
	mysqlConfig := mysql.Config{
		DSN:                       m.DSN(),
		SkipInitializeWithVersion: false,
	}

	// 连接
	db, err := gorm.Open(mysql.New(mysqlConfig))
	if err != nil {
		log.Fatal("mysql connect error：", err)
	}

	// 设置连接池
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	log.Println("connect mysql success")
	return db
}
