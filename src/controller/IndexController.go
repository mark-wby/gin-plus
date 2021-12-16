package controller

import (
	"github.com/mark-wby/gin-plus/src/ginPlusCore"
	"github.com/gin-gonic/gin"
)

type IndexController struct {
	*ginPlusCore.GormAdapter
}

//构建路由
func (this *IndexController) Build(goft *ginPlusCore.GinPlusCore){
	goft.Handle("GET","/",this.GetIndex)
}

//构造函数
func NewIndexController() *IndexController {
	//return &IndexClass{GormAdapter:ginPlusCore.GoftInstance.Dba}
	return &IndexController{}
}

// @summary 服务管理
// @Description 服务初始连接测试
// @Accept json
// @Produce json
// @Header 200 {string} Token "qwerty"
// Success 200 {object} Response
// Failure 400 {object} ResponseError
// Failure 404 {object} ResponseError
// Failure 500 {object} ResponseError
// @Router /check [get]
func (this *IndexController) GetIndex(ctx *gin.Context) interface{}{
	//res,_ :=ctx.Get("custom")
	//
	//val, err :=res.(*customUtil.CustomResponseWrite)
	//if err{
	//	//断言成功
	//
	//}
	return "str";
}


