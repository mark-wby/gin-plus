package ginPlusCore

import (
	"fmt"
	"github.com/mark-wby/gin-plus/src/customExecption"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

type Responder interface {
	RespondTo() gin.HandlerFunc
}

//定义一个返回任意类型的函数
type InterfaceResponder func(ctx *gin.Context) interface{}

func (this InterfaceResponder) RespondTo() gin.HandlerFunc  {
	return func(context *gin.Context) {
		//异常捕获处理
		defer func() {
			if e :=recover();e!=nil{
				//抛出异常,对异常进行断言
				execption,err := e.(customExecption.CustomExecption)
				if err {
					context.JSON(http.StatusOK,gin.H{
						"code":execption.Code,
						"status":false,
						"msg":execption.Msg,
						"data":nil,
					})
				}else {
					context.JSON(http.StatusOK,gin.H{
						"code":500,
						"status":false,
						"msg":"系统错误",
						"data":nil,
					})
				}
			}
		}()

		//正常流程
		res :=this(context)
		context.JSON(http.StatusOK,gin.H{
			"code":200,
			"status":true,
			"msg":"调用成功",
			"data":res,
		})
	}
}

//定义stringResponder为返回string函数
//type StringResponder func(ctx *gin.Context) string
//
//
//func (this StringResponder) RespondTo() gin.HandlerFunc{
//	return func(context *gin.Context) {
//		res :=this(context)
//		context.JSON(http.StatusOK,gin.H{
//			"status":true,
//			"name":res,
//		})
//	}
//}

////定义modelResponder为返回实体函数
//type ModelResponder func(ctx *gin.Context) Model
//
//func (this ModelResponder) RespondTo() gin.HandlerFunc{
//	return func(context *gin.Context) {
//		res :=this(context)
//		context.JSON(http.StatusOK,res)
//	}
//}
//
////定义ModelListResponder为返回实体类切片
//type ModelListResponder func(ctx *gin.Context) []Model
//
//func (this ModelListResponder) RespondTo() gin.HandlerFunc{
//	return func(context *gin.Context) {
//		res :=this(context)
//		context.JSON(http.StatusOK,res)
//	}
//}


var ResponderList []Responder

//初始化函数
func init(){
	ResponderList = []Responder{
		//new(StringResponder),
		//new(ModelResponder),
		//new(ModelListResponder),
		new(InterfaceResponder),
	}
}

//转化函数
func Convert(handler interface{}) gin.HandlerFunc{
	v1:=reflect.ValueOf(handler)
	fmt.Println(v1.Type())
	for _,r:=range ResponderList{
		fr:=reflect.ValueOf(r).Elem()
		fmt.Println(fr.Type())
		if v1.Type().ConvertibleTo(fr.Type()){
			fr.Set(v1)//将对应要执行的方法设置成Responder对象
			return  fr.Interface().(Responder).RespondTo()//断言成responder对象并调用方法
		}
	}
	return nil
}



