package event

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/tangxusc/cavy-sidecar/pkg/db"
	"github.com/tangxusc/cavy-sidecar/pkg/rpc"
)

func CallHandler(ctx context.Context, events []*Event) {
	marks, e := rpc.CallEventHandler(ctx, events)
	if e != nil {
		logrus.Errorf("[event]调用事件处理器失败,事件不做更新.")
		return
	}
	ids := make([]string, 0)
	for _, value := range marks {
		if value.Success {
			ids = append(ids, value.Id)
		}
	}
	//TODO:失败如何处理?
	e = UpdateEventHandlerStatus(ids)
	if e != nil {
		logrus.Errorf("[event]事件状态更新到已调用失败,错误:%v", e.Error())
	}
}

/*
找到未处理的事件
*/
func LoadUnHandEvent() ([]*Event, error) {
	events := make([]*Event, 0)
	e := db.Query(`select * from events where handler_status=0 order by create_time limit 100`, func() interface{} {
		result := &Event{}
		events = append(events, result)
		return result
	})
	return events, e
}
