package custom

import (
	"github.com/gin-gonic/gin"
)

//设置请求上下文
var RequestContext *gin.Context

//设置全局config配置变量
var CustomConfig map[string]map[string]string

