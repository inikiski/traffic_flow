package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Init() {
	username := "postgres"   //账号
	password := "abcd1234"   //密码
	host := "39.108.214.230" //数据库地址，可以是Ip或者域名
	port := 5432             //数据库端口
	Dbname := "traffic_flow"
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s", username, password, host, port, Dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	DB = db
}
