package main

import (
	"github.com/CodingJzy/library_backend/core"
	"github.com/CodingJzy/library_backend/initialize"
)

func main() {
	// 初始化配置
	initialize.LoadConfig()

	// 连接数据库
	initialize.Mysql()

	// 自动创建表
	initialize.CreateTables()

	// 启动server
	core.RunServer()
}
