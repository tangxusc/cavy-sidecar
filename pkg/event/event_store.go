package event

import "time"

type Event struct {
	Id      string
	AggId   string
	AggType string
	Create  time.Time
	data    interface{}
}

func FindEventByTime(id string, AggregateType string, t time.Time) []*Event {
	//TODO:查找事件
	return make([]*Event, 10)
}

func Save(events []*Event) {
	//TODO:保存event,要在同一个事务中,保存后状态是[未发送到消息中间件]
}
