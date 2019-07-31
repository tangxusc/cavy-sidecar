package event

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/tangxusc/cavy-sidecar/pkg/db"
	"github.com/tangxusc/cavy-sidecar/pkg/model"
)

//TODO:收到消息中间件发送的消息后,在本地消息表中记录,并rpc调用业务系统处理消息
//处理消息后,对消息进行记录

//发送到消息中间件,异步
func Send(events []*model.Event) {
	go func() {
		for _, value := range events {
			eventChan <- value
		}
	}()
}

var eventChan = make(chan *model.Event)

func Listen(ctx context.Context) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				close(eventChan)
				return
			case ev := <-eventChan:
				logrus.Debugf("[event]收到event:%v", ev)
				handler(ev)
			}
		}
	}()
}

/*
发送event到消息中间件
其中数据来源: eventChan 和 数据库中发送失败的消息
数据库中发送失败的消息,需要重新发送
*/
func handler(event *model.Event) {
	//TODO:发送event到消息中间件
}

func LoadUnSendEvent() ([]*model.Event, error) {
	events := make([]*model.Event, 0)
	e := db.Query(`select * from events where mq_status=0 order by create_time limit 100`, func() interface{} {
		result := &model.Event{}
		events = append(events, result)
		return result
	})
	return events, e
}
