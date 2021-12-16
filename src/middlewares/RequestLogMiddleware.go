package middlewares

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"ginPlus/src/custom"
	"ginPlus/src/ginPlusCore"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)



type RequestLogMiddleware struct {

}

//请求日志中间件构造函数
func NewRequestLogMiddleware() *RequestLogMiddleware {
	return &RequestLogMiddleware{}
}

//中间件处理函数
func (this *RequestLogMiddleware) OnRequest(request *http.Request) error{
	//将日志链路ID注入请求对象(请求对象每次进来都是新的)
	//判断请求是否存在链路ID
	plusTraceId := request.Header.Get("plusTraceId");
	if plusTraceId == "" {
		//不存在则添加
		str := strconv.Itoa(int(time.Now().UnixNano()))
		request.Header.Set("plusTraceId",str)
	}

	//设置请求开始时间
	request.Header.Set("startTime",strconv.Itoa(int(time.Now().UnixNano())))

	return nil
}

func (this *RequestLogMiddleware) OnRequestAfter(context *gin.Context,write *custom.CustomResponseWrite) error {
	//获取请求中执行的sql和自定义数据

	//将请求返回值变成map
	responseData := make(map[string]interface{},10)

	//将返回值解析成map
	json.Unmarshal(write.Body.Bytes(),&responseData)

	//fmt.Println(string(write.Body.Bytes()))

	msg :=map[string]interface{}{
		"requestPath":context.Request.URL.Path,
		"requestParam":write.RequestParam,
		"requestResponse":responseData,
		"requestSqlLog":write.LogUtil.GetSqlLog(),
		"requestLog":write.LogUtil.GetCustomLog(),
		"requestProjectName":"ceshi",
		"createdAt":context.Request.Header.Get("startTime"),
		"updatedAt":time.Now().UnixNano(),
		"requestId":fmt.Sprintf("%x",md5.Sum([]byte(context.Request.Header.Get("plusTraceId")))),
		"uniqueTraceId":context.Request.Header.Get("plusTraceId"),
	}
	//fmt.Println(msg)
	res,_ := json.Marshal(msg)
	fmt.Println(string(res))
	ginPlusCore.GinPlusCoreInstance.MqUtil.PushMsg(string(res))
	//fmt.Println(err)
	return nil
}
