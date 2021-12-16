##gin-plus是基于gin框架做的脚手架,功能如下
1. 敏捷开发,代码层次分明,目前有中间件层、控制器、工具层、插件层、异常层
2. 使用gorm作为orm,目前只支持mysql
3. 工具集合了redis、mq
4. 请求参数和返回参数进行了统一处理,可以从指定的地方获取
5. 添加gorm插件,能够获取每次执行的sql



###用法

```
    ginPlusCore.NewGinPlusCore().
    		DB(ginPlusCore.NewGormAdapter()).
    		InitRedis().
    		InitMq().
    		Attach(middlewares.NewRequestLogMiddleware()).
    		Mount("v1", controller.NewIndexController()).
    		Launch()
```

___
1. 核心启动结构体ginPlusCore,Launch()方法对应de是启动方法
2. NewGinPlusCore()实例化
3. DB()初始化数据库
4. InitRedis()初始化redis工具
5. InitMq()初始化Mq工具
6. Attach()绑定中间件,中间件需要实现方法OnRequest(request *http.Request)和OnRequestAfter(ctx *gin.Context,write *custom.CustomResponseWrite)
7. Mount()绑定控制器,控制器需实现方法Build(goft *GinPlusCore)
