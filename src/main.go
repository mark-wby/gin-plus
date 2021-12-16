package main

import (
	"ginPlus/src/controller"
	"ginPlus/src/ginPlusCore"
	"ginPlus/src/middlewares"
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
