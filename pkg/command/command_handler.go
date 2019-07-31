package command

import (
	"context"
	"fmt"
	"github.com/tangxusc/cavy-sidecar/pkg/config"
	"github.com/tangxusc/cavy-sidecar/pkg/model"
	"time"
)

func handler(ctx context.Context, cmd *model.Command) {
	identKey := getIdentKey(cmd.AggregateId, cmd.AggregateType)
	agg := getAggregate(ctx, identKey)
	agg.SendCommand(cmd)
}

func getAggregate(ctx context.Context, key string) model.Aggregate {
	ctx, cancel := context.WithCancel(ctx)
	agg := Instance(ctx, key)
	actual, loaded := aggregateMap.LoadOrStore(key, agg)
	if loaded {
		return actual.(model.Aggregate)
	}
	//设置存活时间
	duration := time.Minute * time.Duration(config.Instance.Aggregate.LifeTime)
	afterFunc := time.AfterFunc(duration, func() {
		cancel()
	})
	//启动
	agg.Listen(ctx, func() {
		afterFunc.Reset(duration)
	})
	return agg
}

func getIdentKey(id string, typeString string) string {
	return fmt.Sprintf("%s-%s", id, typeString)
}
