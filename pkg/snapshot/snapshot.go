package snapshot

import (
	"context"
	"github.com/sirupsen/logrus"
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

func FindLastSnapBy(id, aggregateType string) *Aggregate {
	//TODO:如何获取快照?
	return nil
}

var aggChan chan interface{}

func Listen(ctx context.Context) {
	aggChan = make(chan interface{})
	go func() {
		for {
			select {
			case <-ctx.Done():
				close(aggChan)
				return
			case agg := <-aggChan:
				logrus.Debugf("snapshot 收到agg:%v", agg)
				handler(agg)
			}
		}
	}()
}

func handler(agg interface{}) {
	//TODO:按照策略存储快照
}

func SendAggregate(agg interface{}) {
	go func() {
		aggChan <- agg
	}()
}
