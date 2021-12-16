package ginPlusCore

import (
	"bytes"
	"encoding/json"
	"fmt"
	"ginPlus/src/custom"
	"ginPlus/src/customConfig"
	"ginPlus/src/customUtil"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"reflect"
)

//设置全局goft实例
var GinPlusCoreInstance *GinPlusCore

type GinPlusCore struct {
	*gin.Engine                     //gin框架引擎
	Gp *gin.RouterGroup             //gin框架路由
	dba *GormAdapter                //数据库链接
	RedisUtil *customUtil.RedisUtil //redis链接
	MqUtil *customUtil.MqUtil       //mq链接
}

//实例构造函数
func NewGinPlusCore() *GinPlusCore {
	//gin.SetMode(gin.ReleaseMode)
	//初始化配置文件
	data,_:=ioutil.ReadFile("config/Config.yaml")
	res := make(map[string]map[string]string,0)
	yaml.Unmarshal(data,res)
	customConfig.CustomConfig = res

	return &GinPlusCore{Engine: gin.New()}
}


//设置数据库链接
func (this *GinPlusCore) DB(adapter *GormAdapter) *GinPlusCore{
	this.dba = adapter
	//设置goft实例
	GinPlusCoreInstance = this
	return this
}

//初始化redis链接
func (this *GinPlusCore) InitRedis() *GinPlusCore{
	this.RedisUtil = customUtil.NewRedisUtil()
	return this
}

//初始化mq链接
func (this *GinPlusCore) InitMq() *GinPlusCore{
	this.MqUtil = customUtil.NewMqUtil()
	return this
}

//重载路由函数
func (this *GinPlusCore) Handle(httpMethod, relativePath string, handlers interface{}) gin.IRoutes {
	if h:= Convert(handlers);h!=nil{
		return this.Gp.Handle(httpMethod,relativePath,h)
	}
	return nil
}

//最终gin的启动函数
func (this *GinPlusCore) Launch(){
	this.Run(":"+customConfig.CustomConfig["httpServer"]["port"])
}

//加载多个中间件
func (this *GinPlusCore) Attach(milleware ...IMilleware) *GinPlusCore  {
	this.Use(func(context *gin.Context) {
		//替换自定义的response(可以存储响应内容)
		blw := &custom.CustomResponseWrite{
			Body:           bytes.NewBufferString(""),
			ResponseWriter: context.Writer,
			LogUtil:        customUtil.NewLoggerUtil(),
		}

		//解析任何请求方式的请求参数,塞入结构体中
		//解析get请求和post请求的form参数
		//request.ParseForm()
		context.Request.ParseMultipartForm(128)
		// 获取请求实体长度
		contentLength := context.Request.ContentLength
		body := make([]byte, contentLength)
		//获取请求体数据
		context.Request.Body.Read(body)
		//定义map结构接受数据
		event := make(map[string]interface{},0)
		//将json数据解析成map
		json.Unmarshal(body, &event)
		//由于其他不是json请求获取到参数不能满足map,需要进行转化
		tmpData := make(map[string]interface{},0)
		for k,v := range context.Request.Form{
			if len(v)>1{
				tmpData[k] = v
			}else {
				tmpData[k] = v[0]
			}

		}

		blw.RequestParam = customUtil.MergeMap(tmpData,event)

		context.Writer = blw

		context.Set("custom",blw)

		//修改全局参数
		custom.RequestContext = context


		for _, value := range milleware {
			err :=value.OnRequest(context.Request)
			if err!=nil{
				context.AbortWithStatusJSON(http.StatusOK,gin.H{
					"status":false,
					"msg":err.Error(),
					"data":nil,
					"code":500,
				})
			}else {
				context.Next()
				value.OnRequestAfter(context,blw)
			}
		}
	})
	return this
}


//挂在函数,主要是路由和控制器
func (this *GinPlusCore) Mount(group string,classes ...IClass) *GinPlusCore{
	this.Gp = this.Group(group)

	//注入swagger文档地址
	this.Gp.Handle("GET","/swagger", func(context *gin.Context) {
		data,err :=ioutil.ReadFile("src/docs/swagger.json")
		if err!=nil{
			fmt.Println(err.Error())
		}
		context.String(http.StatusOK,string(data))
	})

	//注册404路由不存在处理
	this.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusOK,gin.H{
			"status":false,
			"msg":"请求路径不存在",
			"data":nil,
			"code":404,
		})
	})

	//注册请求方法不存在处理
	this.NoMethod(func(context *gin.Context) {
		context.JSON(http.StatusOK,gin.H{
			"status":false,
			"msg":"请求方式不存在",
			"data":nil,
			"code":405,
		})
	})


	for _,value := range classes {
		//注册路由
		value.Build(this)

		//判断控制器是否存在数据库链接属性
		vClass := reflect.ValueOf(value).Elem()
		if vClass.NumField()>0{
			if this.dba != nil{
				vClass.Field(0).Set(reflect.New(vClass.Field(0).Type().Elem()))//先申请一个地址进行赋值
				vClass.Field(0).Elem().Set(reflect.ValueOf(this.dba).Elem())//再将数据的指针复制到新申请的地址
			}
		}
	}
	return this
}
