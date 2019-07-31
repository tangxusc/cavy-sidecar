package command

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/tangxusc/cavy-sidecar/pkg/config"
	"github.com/tangxusc/cavy-sidecar/pkg/event"
	"time"
)

/*
每x秒获取数据库中未被处理的event,发送到eventhandler中处理
每x秒获取数据库中未被发送的event,发送到event_sender中处理
*/
func StartEventRecover(ctx context.Context) {
	go dbRecover(ctx)
	go busRecover(ctx)
}

func busRecover(ctx context.Context) {
	funRecover(ctx, func() {
		//查询数据库
		events, e := event.LoadUnSendEvent()
		if e != nil {
			logrus.Errorf("[command]尝试恢复数据库中未发送事件出现错误,%v", e.Error())
			return
		}
		//发送
		event.Send(events)
	}, time.Duration(config.Instance.Aggregate.RecoverTime))
}

func dbRecover(ctx context.Context) {
	funRecover(ctx, func() {
		//查询数据库
		events, e := event.LoadUnHandEvent()
		if e != nil {
			logrus.Errorf("[command]尝试恢复数据库中未处理事件出现错误,%v", e.Error())
			return
		}
		//发送到聚合事件处理器
		for _, value := range events {
			agg := getAggregate(ctx, getIdentKey(value.AggId, value.AggType))
			agg.SendEvent(value)
		}
	}, time.Duration(config.Instance.Aggregate.RecoverTime))
}

func funRecover(ctx context.Context, f func(), duration time.Duration) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			time.Sleep(time.Second * duration)
			f()
		}
	}
}
