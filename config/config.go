package config

import "fmt"

type Config struct {
	Mysql Mysql `json:"mysql" yaml:"mysql"`
	System
}

type System struct {
	Addr string
}

type Mysql struct {
	Addr     string
	UserName string
	Password string
	DbName   string
}

//  dsn："user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
func (m Mysql) DSN() string {
	return fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", m.UserName, m.Password, m.Addr, m.DbName)
}
