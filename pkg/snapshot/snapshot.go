package snapshot

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/tangxusc/cavy-sidecar/pkg/db"
	"github.com/tangxusc/cavy-sidecar/pkg/event"
	"time"
)

type Aggregate struct {
	Id      string
	AggType string
	AggId   string
	Create  time.Time
	//这样怎么序列化>?
	Data *interface{}
}

//TODO:test
func FindLastSnapBy(id, aggregateType string) (agg *Aggregate, err error) {
	err = db.QueryRow(`select * from snapshot where AggId=? and AggType=? order by create desc`, agg, id, aggregateType)
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
				go handler(agg)
			}
		}
	}()
}

//TODO:test
/*
存储快照
存储出现错误,可以忽略;
*/
func handler(agg *Aggregate) {
	_, e := db.NameExec(`INSERT INTO snapshot(Id,AggType,AggId,Create,Data) VALUES(:Id,:AggType,:AggId,:Create,:Data)`, agg)
	if e != nil {
		logrus.Errorf("[snapshot]存储快照出现错误,聚合:%v,错误:%v", agg, e)
	} else {
		logrus.Debugf("[snapshot]存储快照完成")
	}
}

func SaveAggregate(Id string, aggType string, agg *interface{}, events []*event.Event) {
	if DefaultStrategy.Allow(Id, aggType, agg, events) {
		go func() {
			aggChan <- &Aggregate{
				AggType: aggType,
				AggId:   Id,
				Data:    agg,
			}
		}()
	}
}

type SaveStrategy interface {
	Allow(Id string, aggType string, agg *interface{}, events []*event.Event) bool
}

type CountStrategy struct {
	Max int
}

func (c *CountStrategy) Allow(Id string, aggType string, agg *interface{}, events []*event.Event) bool {
	if len(events) > c.Max {
		return true
	}
	return false
}

//默认提供基于事件数量的快照
var DefaultStrategy SaveStrategy = &CountStrategy{Max: 100}
