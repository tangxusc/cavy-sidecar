package snapshot

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/tangxusc/cavy-sidecar/pkg/db"
	"github.com/tangxusc/cavy-sidecar/pkg/event"
	"time"
)

type Aggregate struct {
	Id         string    `db:"id"`
	AggType    string    `db:"agg_type"`
	AggId      string    `db:"agg_id"`
	CreateTime time.Time `db:"create_time"`
	Data       []byte    `db:"data"`
}

/*
查找并返回最后一个快照,按照创建事件倒序
*/
func FindLastSnapBy(id, aggregateType string) (agg *Aggregate, err error) {
	agg = &Aggregate{}
	err = db.QueryRow(`select * from snapshot where agg_id=? and agg_type=? order by create_time desc`, agg, id, aggregateType)
	return
}

var aggChan chan *Aggregate

func Listen(ctx context.Context) {
	aggChan = make(chan *Aggregate)
	go func() {
		for {
			select {
			case <-ctx.Done():
				close(aggChan)
				return
			case agg := <-aggChan:
				logrus.Debugf("[snapshot]收到agg:%v", agg)
				go SaveSnapshot(agg)
			}
		}
	}()
}

/*
存储快照
存储出现错误,可以忽略;
*/
func SaveSnapshot(agg *Aggregate) {
	_, e := db.NameExec(`INSERT INTO snapshot(id,agg_type,agg_id,create_time,data) VALUES(:id,:agg_type,:agg_id,:create_time,:data)`, agg)
	if e != nil {
		logrus.Errorf("[snapshot]存储快照出现错误,聚合:%v,错误:%v", agg, e)
	} else {
		logrus.Debugf("[snapshot]存储快照完成")
	}
}

func SaveAggregate(Id string, aggType string, agg []byte, events []*event.Event) {
	if DefaultStrategy.Allow(Id, aggType, agg, events) {
		aggChan <- &Aggregate{
			AggType: aggType,
			AggId:   Id,
			Data:    agg,
		}
	}
}

type SaveStrategy interface {
	Allow(Id string, aggType string, agg []byte, events []*event.Event) bool
}

type CountStrategy struct {
	Max int
}

func (c *CountStrategy) Allow(Id string, aggType string, agg []byte, events []*event.Event) bool {
	if len(events) > c.Max {
		return true
	}
	return false
}

//默认提供基于事件数量的快照
var DefaultStrategy SaveStrategy = &CountStrategy{Max: 100}
