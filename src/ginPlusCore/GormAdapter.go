package ginPlusCore

import (
	"fmt"
	"ginPlus/src/customConfig"
	"ginPlus/src/plugin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type DbConfig struct {
	Server struct{
		Port string `yaml:"port"`
		Host string `yaml:"host"`
		Database string	`yaml:"database"`
		Charset string `yaml:"charset"`
		UserName string `yaml:"userName"`
		Password string `yaml:"password"`
	}
}


type GormAdapter struct {
	*gorm.DB
}


//创建对象
func NewGormAdapter() *GormAdapter {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=Local",
		customConfig.CustomConfig["mysqlServer"]["userName"],
		customConfig.CustomConfig["mysqlServer"]["password"],
		customConfig.CustomConfig["mysqlServer"]["host"],
		customConfig.CustomConfig["mysqlServer"]["port"],
		customConfig.CustomConfig["mysqlServer"]["database"],
		customConfig.CustomConfig["mysqlServer"]["charset"],
		)

	//dsn :="root:WBY242436biao!@tcp("+dbConfig.Server.Host+":"+dbConfig.Server.Port+")/"+dbConfig.Server.Database+"?charset="+dbConfig.Server.Charset+"&parseTime=true&loc=Local"
	//fmt.Println(dsn)
	db,err := gorm.Open(mysql.Open(dsn), &gorm.Config{} )
	if err !=nil{
		fmt.Println("数据库链接失败")
		log.Fatal(err.Error())
	}

	//注入自定义插件,获取执行的sql
	db.Use(new(plugin.DbPlugin))

	return &GormAdapter{DB: db}
}


