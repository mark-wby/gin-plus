package ginPlusCore

import (
	"context"
	"sync"
	"time"
)

//事件总线数据结构体
type EventData struct {
	Data interface{}

}

//定义传递事件数据的通道
type EventChannelData chan *EventData

//获取数据
func (this EventChannelData) Data(time time.Duration) interface{}  {
	ctx,cancel := context.WithTimeout(context.Background(),time)
	defer cancel()

	select {
		case <-ctx.Done():
			//超时返回
			return "timeout"
			break
		case data:=<-this:
			//正确返回
			return data
			break
	}
	return nil
}



//事件总线结构体
type EventBus struct {
	subscribes map[string]EventChannelData
	lock sync.RWMutex
}

func NewEventBus() *EventBus {
	return &EventBus{subscribes:make(map[string]EventChannelData)}
}

//订阅
func (this *EventBus)Sub(topic string)EventChannelData  {
	this.lock.Lock()
	defer this.lock.Unlock()
	if ec,found :=this.subscribes[topic];found{
		return ec
	}else {
		this.subscribes[topic] = make(EventChannelData)
		return  this.subscribes[topic]
	}
}

//发布
func (this *EventBus)Pub(topic string,data interface{})  {
	this.lock.Lock()
	defer this.lock.Unlock()
	if ec,found :=this.subscribes[topic];found{
		go func() {
			evenData := new(EventData)
			evenData.Data = data
			ec<-evenData
		}()
	}else {
		go func() {
			eventChannelData := make(EventChannelData)
			eventChannelData<-&EventData{Data:data}
			this.subscribes[topic] = eventChannelData
		}()
	}
}


