package ginPlusCore

import (
	"ginPlus/src/custom"
	"github.com/gin-gonic/gin"
	"net/http"
)

//定义中间件接口
type IMilleware interface {
	OnRequest(request *http.Request)	error                                //定义中间件前置处理函数
	OnRequestAfter(ctx *gin.Context,write *custom.CustomResponseWrite) error //定义中间件后置处理函数
}

