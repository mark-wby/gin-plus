package customUtil

import (
	"fmt"
	"github.com/mark-wby/gin-plus/src/custom"
	"github.com/streadway/amqp"
	"log"
	"strconv"
)

//mq链接配置
type MqConfig struct {
	Mqserver struct{
		Host string `yaml:"host"`
		Port int	`yaml:"port"`
		UserName string	`yaml:"userName"`
		Password string	`yaml:"password"`
	}
}

type MqUtil struct {
	MqConn *amqp.Connection
}

func NewMqUtil() *MqUtil {

	port,_ := strconv.Atoi(custom.CustomConfig["mqServer"]["port"])

	dsn := fmt.Sprintf("amqp://%s:%s@%s:%d/",custom.CustomConfig["mqServer"]["userName"],custom.CustomConfig["mqServer"]["password"],
		custom.CustomConfig["mqServer"]["host"],port)
	fmt.Println(dsn)
	conn,err := amqp.Dial(dsn)
	if err !=nil{
		fmt.Println("mq链接失败")
		log.Fatal(err)
	}
	return &MqUtil{MqConn: conn}
}

//推送消息到mq
func (this *MqUtil) PushMsg(body string) error{
	channel,err := this.MqConn.Channel()
	if err!= nil{
		log.Fatal(err)
	}
	defer channel.Close()

	err = channel.Publish("logExchange","test",false,false,
		amqp.Publishing{
			ContentType:"text/plain",
			Body:[]byte(body),
		})
	return err
}



