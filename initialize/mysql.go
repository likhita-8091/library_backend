package initialize

import (
	"github.com/CodingJzy/library_backend/global"
	"github.com/CodingJzy/library_backend/model"
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
	global.DB = db
	return db
}

func CreateTables() {
	// 创建表
	err := global.DB.AutoMigrate(
		&model.User{},
		&model.BookKind{},
		&model.Book{},
		&model.T1{},
		&model.CreditCard{},
	)
	if err != nil {
		log.Fatal("create tables error：", err)
	}
	log.Println("create tables success")

	// 创建默认用户admin
	err = global.DB.FirstOrCreate(model.NewAdmin()).Debug().Error
	if err != nil {
		log.Fatal("create admin user error：", err.Error())
	}

	log.Println("create admin user success")
}
