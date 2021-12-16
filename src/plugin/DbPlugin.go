package plugin

import (
	"ginPlus/src/custom"
	"gorm.io/gorm"
)

type DbPlugin struct {


}

func (*DbPlugin) Name() string {
	return "CustomPlugin"
}

func (*DbPlugin)Initialize(db *gorm.DB) error {
	//db.Callback().Query().After("gorm:query").Register("querySql",querySql)
	// 所有其它 callback 之后
	// gorm:create 之后
	db.Callback().Create().After("gorm:create").Register("update_created_at", querySql)

	// gorm:query 之后
	db.Callback().Query().After("gorm:query").Register("my_plugin:after_query", querySql)

	// gorm:delete 之后
	db.Callback().Delete().After("gorm:delete").Register("my_plugin:after_delete", querySql)

	// gorm:update 之前
	db.Callback().Update().After("gorm:update").Register("my_plugin:before_update", querySql)
	return nil
}

func querySql(db *gorm.DB){
	sql :=db.Dialector.Explain(db.Statement.SQL.String(),db.Statement.Vars...)
	//fmt.Println("打印中-==========")
	//fmt.Println(sql)

	//获取全局请求上下文
	context := custom.RequestContext
	res,_ :=context.Get("custom")

	val, err :=res.(*custom.CustomResponseWrite)
	if err{
		//断言成功,写入日志数据
		val.LogUtil.InfoSqlLog(sql)
	}
}