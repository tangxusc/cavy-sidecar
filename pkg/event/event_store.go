package event

import (
	"github.com/jmoiron/sqlx"
	"github.com/tangxusc/cavy-sidecar/pkg/db"
	"time"
)

type Event struct {
	Id      string
	AggId   string
	AggType string
	Create  time.Time
	data    interface{}
}

/**
查找时间t之后的事件
*/
func FindEventByTime(id string, AggregateType string, t time.Time) ([]*Event, error) {
	events := make([]*Event, 0)
	e := db.Query(`select * from events where AggId=? and AggType=? and Create> ?`, &events, id, AggregateType, t)
	if e != nil {
		return events, nil
	}
	return events, nil
}

//TODO:test
/*
保存event到数据库
event必须在同一个事务中成功或者失败
*/
func Save(events []*Event) {
	db.Transaction(func(tx *sqlx.Tx) error {
		for _, value := range events {
			_, e := tx.NamedExec(`INSERT INTO events(Id,AggType,AggId,Create,Data) VALUES(:Id,:AggType,:AggId,:Create,:Data)`, value)
			if e != nil {
				return e
			}
		}
		return nil
	})
	//TODO:保存event,要在同一个事务中,保存后状态是[未发送到消息中间件]
}
