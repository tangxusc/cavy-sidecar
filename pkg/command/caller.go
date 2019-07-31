package command

import (
	"context"
	"github.com/tangxusc/cavy-sidecar/pkg/event"
	"github.com/tangxusc/cavy-sidecar/pkg/rpc"
	"time"
)

func CallAggregate(ctx context.Context, aggId string, aggType string, agg []byte, cmd *Command) ([]*event.Event, error) {
	events, e := rpc.CallAggregate(ctx, aggId, aggType, agg, cmd)
	if e != nil {
		return nil, e
	}
	result := make([]*event.Event, len(events))
	for key, value := range events {
		any := value.Data
		result[key] = &event.Event{
			Id:        value.Id,
			EventType: any.TypeUrl,
			AggId:     aggId,
			AggType:   aggType,
			Create:    time.Now(),
			Data:      any.Value,
		}
	}
	return result, nil
}
