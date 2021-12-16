package main

import (
	"github.com/mark-wby/gin-plus/src/controller"
	"github.com/mark-wby/gin-plus/src/ginPlusCore"
	"github.com/mark-wby/gin-plus/src/middlewares"
)


// @title gin基础项目
// @version 1.0
// @host      localhost:8080
// @BasePath  /v1
func main(){
	ginPlusCore.NewGinPlusCore().
		DB(ginPlusCore.NewGormAdapter()).
		InitRedis().
		InitMq().
		Attach(middlewares.NewRequestLogMiddleware()).
		Mount("v1", controller.NewIndexController()).
		Launch()
}
