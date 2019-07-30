package event

import (
	"github.com/jmoiron/sqlx"
	"github.com/tangxusc/cavy-sidecar/pkg/db"
	"time"
)

type Status int

const (
	//已保存
	Saved Status = iota
	//已发送
	Sent
	//发送错误
	SendError
)

type Event struct {
	Id      string    `db:"id"`
	AggId   string    `db:"agg_id"`
	AggType string    `db:"agg_type"`
	Create  time.Time `db:"create_time"`
	Data    []byte    `db:"data"`
	Status  Status    `db:"status"`
}

/**
查找时间t之后的事件
*/
func FindEventByTime(id string, AggregateType string, t time.Time) ([]*Event, error) {
	events := make([]*Event, 0)
	e := db.Query("select * from events where agg_id=? and agg_type=? and create_time> ?", func() interface{} {
		result := &Event{}
		events = append(events, result)
		return result
	}, id, AggregateType, t)
	if e != nil {
		return events, nil
	}
	return events, nil
}

/*
保存event到数据库
event必须在同一个事务中成功或者失败
保存event,要在同一个事务中,保存后状态是[未发送到消息中间件]
*/
func Save(events []*Event) {
	db.Transaction(func(tx *sqlx.Tx) error {
		for _, value := range events {
			value.Status = Saved
			_, e := tx.NamedExec("INSERT INTO events(id,agg_type,agg_id,create_time,data,status) VALUES(:id,:agg_type,:agg_id,:create_time,:data,:status)", value)
			if e != nil {
				return e
			}
		}
		return nil
	})
}
