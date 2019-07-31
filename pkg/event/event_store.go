package event

import (
	"github.com/jmoiron/sqlx"
	"github.com/tangxusc/cavy-sidecar/pkg/db"
	"github.com/tangxusc/cavy-sidecar/pkg/model"
	"time"
)

/**
查找时间t之后的事件
*/
func FindEventByTime(id string, AggregateType string, t time.Time) ([]*model.Event, error) {
	events := make([]*model.Event, 0)
	e := db.Query("select * from events where agg_id=? and agg_type=? and create_time> ?", func() interface{} {
		result := &model.Event{}
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
func Save(events []*model.Event) error {
	return db.Transaction(func(tx *sqlx.Tx) error {
		for _, value := range events {
			value.HandlerStatus = model.Untreated
			value.MqStatus = model.MqNotSend
			_, e := tx.NamedExec(`INSERT INTO events(id,event_type,agg_type,agg_id,create_time,data,handler_status,mq_status)
 VALUES(:id,:event_type,:agg_type,:agg_id,:create_time,:data,:handler_status,:mq_status)`, value)
			if e != nil {
				return e
			}
		}
		return nil
	})
}

func UpdateEventHandlerStatus(ids []string) error {
	if len(ids) == 0 {
		return nil
	}
	return db.Transaction(func(tx *sqlx.Tx) error {
		query, args, err := sqlx.In(`update events set handler_status=1 where id in (?)`, ids)
		if err != nil {
			return err
		}
		query = tx.Rebind(query)
		_, err = tx.Exec(query, args...)
		return err
	})
}
