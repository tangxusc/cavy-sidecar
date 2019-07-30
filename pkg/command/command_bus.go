package command

import (
	"context"
	"github.com/sirupsen/logrus"
	"sync"
)

var aggregateMap = sync.Map{}
var CommandChan = make(chan *Command)

func Listen(ctx context.Context) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case cmd := <-CommandChan:
				logrus.Debugf("收到command:%v", cmd)
				go handler(ctx, cmd)
			}
		}
	}()

	//handlerCommand
	//1,commandhandler读取chan
	//2,根据聚合类型查找 该聚合的chan
	// 如果没有找到,那么直接初始化一个 aggregateEventSource,并将chan加入commandBus中 考虑怎么gc这个对象呢?
	// 如果找到了, 则直接将command发送到aggregateEventSource中.
	//3,aggregateEventSource 拿到command后,开启unit of work???

	//4,aggregateEventSource 开始溯源
	//5,aggregateEventSource 先查找快照
	// 如果找到,那么按照快照创建的时间寻找events
	// 加载events
	//6,eventSourcingCaller 发起rpc 溯源得到aggregate
	// 6.1 适当时机保存快照 发送 aggregate 到 snapshotStorage的chan中  做成chan模式

	//7,commandCaller 发起rpc 处理事件  是否做成读取chan的模式?

	//8,获取结果 发送到eventStorage chan中保存  是否做成写入chan的模式? 如果此模式那么如何确定保存成功?

	//8.1 发送时 EventSender chan中,发送到消息中间件  此处可以做成chan模式
	//9,保存后再eventHandlerCaller chan中发送rpc处理  此处可以做成chan模式
}

/*
从map中移除聚合
*/
func RemoveAggregateSourcing(agg Aggregate) {
	aggregateMap.Delete(agg.GetKey())
}
