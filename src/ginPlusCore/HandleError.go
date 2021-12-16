package ginPlusCore

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//错误处理中间件
func ErrorHandle() gin.HandlerFunc  {
	return func(context *gin.Context) {
		defer func() {
			e := recover()
			if e!=nil{
				context.AbortWithStatusJSON(http.StatusOK,gin.H{
					"status":false,
					"msg":e,
				})
			}
		}()
		context.Next()
	}
}

//错误处理
func ThrowError(err error){
	if err!=nil{
		panic(err)
	}
}
