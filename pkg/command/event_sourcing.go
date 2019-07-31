package command

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/tangxusc/cavy-sidecar/pkg/event"
	"github.com/tangxusc/cavy-sidecar/pkg/rpc"
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
	CommandChan chan *Command
	//上下文
	Ctx context.Context
	//聚合对象
	aggregate []byte
	//最后更新时间
	lastTime time.Time
	//eventHandler
	EventChan chan []*event.Event
}

func (agg *Sourcing) SendEvent(e *event.Event) {
	go func() {
		agg.EventChan <- []*event.Event{e}
	}()
}

func (agg *Sourcing) GetKey() string {
	return agg.Key
}

func (agg *Sourcing) SendCommand(command *Command) {
	agg.CommandChan <- command
}

func Instance(ctx context.Context, key string) Aggregate {
	return &Sourcing{
		Key:         key,
		CommandChan: make(chan *Command),
		Ctx:         ctx,
		EventChan:   make(chan []*event.Event),
	}
}

func (agg *Sourcing) Listen(ctx context.Context, reset func()) {
	defer RemoveAggregateSourcing(agg)
	//handlerCommand
	go func() {
		for {
			select {
			case <-ctx.Done():
				close(agg.CommandChan)
				return
			case cmd := <-agg.CommandChan:
				logrus.Debugf("[aggregate]收到command:%v", cmd)
				reset()
				agg.handlerCommand(ctx, cmd)
			}
		}
	}()
	//callEventHandler
	go func() {
		for {
			select {
			case <-ctx.Done():
				close(agg.EventChan)
				return
			case ev := <-agg.EventChan:
				logrus.Debugf("[aggregate]收到command:%v", ev)
				go event.CallHandler(ctx, ev)
			}
		}
	}()
}

//1,溯源聚合
//2,处理command
//3,接受到event
func (agg *Sourcing) handlerCommand(ctx context.Context, cmd *Command) {
	//1,溯源聚合
	//1.1如果之前没有初始化聚合lastTime,那整个聚合从未被溯源,则需要找到快照和快照发生后的events进行溯源
	var events []*event.Event
	var err error
	if agg.lastTime.IsZero() {
		snap, err := snapshot.FindLastSnapBy(cmd.AggregateId, cmd.AggregateType)
		if err != nil {
			logrus.Errorf("[aggregate]查找快照出现错误,命令:%v,错误:%v", cmd, err)
			return
		}
		events, err = event.FindEventByTime(cmd.AggregateId, cmd.AggregateType, snap.CreateTime)
		if err != nil {
			logrus.Errorf("[aggregate]查找快照出现错误,命令:%v,错误:%v", cmd, err)
			return
		}
		agg.aggregate = snap.Data
	} else {
		events, err = event.FindEventByTime(cmd.AggregateId, cmd.AggregateType, agg.lastTime)
		if err != nil {
			logrus.Errorf("[aggregate]查找快照出现错误,命令:%v,错误:%v", cmd, err)
			return
		}
	}
	//1.2远程调用进行溯源
	agg.aggregate, err = CallEventSourcing(ctx, cmd.AggregateId, cmd.AggregateType, agg.aggregate, events)
	if err != nil {
		//TODO:错误返回处理
		return
	}
	//1.2.1溯源后更新时间为当前时间
	agg.lastTime = time.Now()
	//1.2.2发送溯源结果至快照存储,由快照存储策略决定是否存储,及如何存储
	go snapshot.SaveAggregate(cmd.AggregateId, cmd.AggregateType, agg.aggregate, events)

	//2.处理command
	//2.1发起rpc,获取events
	receiveEvents := make([]*event.Event, 0)
	receiveEvents, err = CallAggregate(ctx, cmd.AggregateId, cmd.AggregateType, agg.aggregate, cmd)
	if err != nil {
		logrus.Errorf("[aggregate]调用业务系统处理命令出现错误,聚合:%v,命令:%v,错误:%v", agg.aggregate, cmd, err)
		//TODO:错误返回处理
		return
	}
	//两种情况会导致events为0
	//1,出现错误
	//2,业务系统中并没有返回events,可能出现了业务不允许的情况
	if len(receiveEvents) == 0 {
		return
	}
	//3.1同步保存events到数据库
	err = event.Save(receiveEvents)
	if err != nil {
		//todo:保存到数据库出现问题,错误处理
		return
	}
	//3.3调用事件处理器,进行处理
	agg.EventChan <- receiveEvents
	//3.2发送到消息中间件,异步
	event.Send(receiveEvents)
}

//向业务系统发起rpc,开始事件溯源
func CallEventSourcing(ctx context.Context, id string, aggType string, aggregate []byte, events []*event.Event) ([]byte, error) {
	bytes, e := rpc.Sourcing(ctx, id, aggType, aggregate, events)
	if e != nil {
		return nil, e
	}
	return bytes, nil
}
