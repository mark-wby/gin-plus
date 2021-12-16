package customUtil

import (
	"context"
	"ginPlus/src/customConfig"
	"github.com/go-redis/redis/v8"
	"strconv"
	"time"
)

type RedisConfig struct {
	RedisServer struct{
		Host string `yaml:"host"`
		Port string	`yaml:"port"`
		Password string	`yaml:"password"`
		Db int	`yaml:"db"`
	}
}

//redis工具类
type RedisUtil struct {
	rdb *redis.Client
	ctx context.Context
}

//实力redis
func NewRedisUtil() *RedisUtil {
	//将字符串转为数字
	db,_ := strconv.Atoi(customConfig.CustomConfig["redisServer"]["db"])
	rdb := redis.NewClient(&redis.Options{
		Addr: customConfig.CustomConfig["redisServer"]["host"]+":"+customConfig.CustomConfig["redisServer"]["port"],
		Password: customConfig.CustomConfig["redisServer"]["password"],
		DB: db,
	})

	ctx := context.Background()
	return &RedisUtil{rdb:rdb,ctx:ctx}
}

//获取缓存数据
func (this *RedisUtil) Get(key string) string{
	v,err := this.rdb.Get(this.ctx,key).Result()
	if err == redis.Nil{//数据不存在
		return ""
	}else if err != nil{
		panic(err)
	}
	return v
}

//设置缓存数据
func (this *RedisUtil) Set(key string,value interface{},ttl int){
	exp := ttl*1000000000
	err := this.rdb.Set(this.ctx,key,value,time.Duration(exp)).Err()
	if err != nil{
		panic(err)
	}
}

//设置锁
func (this *RedisUtil) SetLock(key string,value interface{},ttl int){
	exp := ttl*1000000000
	flag,err :=this.rdb.SetNX(this.ctx,key,value,time.Duration(exp)).Result()
	if err!=nil{
		panic(err)
	}
	if !flag {
		panic("加锁失败,请重试")
	}
}




