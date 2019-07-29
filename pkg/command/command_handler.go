package command

import (
	"context"
	"fmt"
	"github.com/tangxusc/cavy-sidecar/pkg/aggregate"
)

type Command struct {
	AggregateId   string
	AggregateType string
	Data          *interface{}
}

func handler(ctx context.Context, command *Command) {
	identKey := getIdentKey(command.AggregateId, command.AggregateType)
	aggregateSourcing := getAggregateEventSourcing(ctx, identKey)
	aggregateSourcing.CommandChan <- command
}

func getAggregateEventSourcing(ctx context.Context, key string) *aggregate.Sourcing {
	//创建,启动,放入aggregateEventCollection
	agg := instance(ctx, key)
	actual, loaded := aggregateEventCollection.LoadOrStore(key, agg)
	if loaded {
		return actual.(*aggregate.Sourcing)
	}
	//启动
	agg.Listen(ctx)
	return agg
}

func instance(ctx context.Context, key string) *aggregate.Sourcing {
	return &aggregate.Sourcing{
		Key:         key,
		CommandChan: make(chan *Command),
		Ctx:         ctx,
	}
}

func getIdentKey(id string, typeString string) string {
	return fmt.Sprintf("%s-%s", id, typeString)
}
