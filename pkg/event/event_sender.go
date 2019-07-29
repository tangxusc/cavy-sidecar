package event

import (
	"context"
	"github.com/sirupsen/logrus"
)

//发送到消息中间件,异步
func Send(events []*Event) {
	go func() {
		for _, value := range events {
			eventChan <- value
		}
	}()
}

var eventChan = make(chan *Event)

func Listen(ctx context.Context) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				close(eventChan)
				return
			case ev := <-eventChan:
				logrus.Debugf("event 收到event:%v", ev)
				handler(ev)
			}
		}
	}()
}

func handler(event *Event) {
	//TODO:发送event到消息中间件
}
