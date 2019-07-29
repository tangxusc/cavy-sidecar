package aggregate

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/tangxusc/cavy-sidecar/pkg/command"
	"github.com/tangxusc/cavy-sidecar/pkg/event"
	"github.com/tangxusc/cavy-sidecar/pkg/snapshot"
	"time"
)

/*
事件溯源
接收到command
1,溯源聚合
2,处理command
3,接受到event
*/
type Sourcing struct {
	Key string
	//cmd命令
	CommandChan chan *command.Command
	//上下文
	Ctx context.Context
	//聚合对象
	aggregate *interface{}
	//最后更新时间
	lastTime time.Time
}

func (agg *Sourcing) Listen(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			close(agg.CommandChan)
			return
		case cmd := <-agg.CommandChan:
			logrus.Debugf("AggregateSourcing 收到command:%v", cmd)
			agg.handlerCommand(cmd)
		}
	}
}

//1,溯源聚合
//2,处理command
//3,接受到event
func (agg *Sourcing) handlerCommand(cmd *command.Command) {
	//1,溯源聚合
	//1.1如果之前没有初始化聚合lastTime,那整个聚合从未被溯源,则需要找到快照和快照发生后的events进行溯源
	var events []*event.Event
	if agg.lastTime.IsZero() {
		snap := snapshot.FindLastSnapBy(cmd.AggregateId, cmd.AggregateType)
		events = event.FindEventByTime(cmd.AggregateId, cmd.AggregateType, snap.Create)
		agg.aggregate = snap.Data
	} else {
		events = event.FindEventByTime(cmd.AggregateId, cmd.AggregateType, agg.lastTime)
	}
	//1.2远程调用进行溯源
	CallEventSourcing(agg.aggregate, events)
	//1.2.1溯源后更新时间为当前时间
	agg.lastTime = time.Now()
	//1.2.2发送溯源结果至快照存储,由快照存储策略决定是否存储,及如何存储
	snapshot.SendAggregate(agg.aggregate)

	//2.处理command
	//2.1发起rpc,获取events
	events = command.CallAggregate(agg.aggregate, cmd)
	//3.1同步保存events到数据库
	event.Save(events)
	//3.2发送到消息中间件,异步
	event.Send(events)
	//3.3调用事件处理器,进行处理
	event.CallHandler(events)
}

func CallEventSourcing(aggregate interface{}, events []*event.Event) *interface{} {
	//TODO:发起rpc,获得结果
	return nil
}
