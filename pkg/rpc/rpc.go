package rpc

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/sirupsen/logrus"
	"github.com/tangxusc/cavy-sidecar/pkg/command"
	"github.com/tangxusc/cavy-sidecar/pkg/config"
	"github.com/tangxusc/cavy-sidecar/pkg/event"
	"google.golang.org/grpc"
)

//go:generate protoc --go_out=plugins=grpc:. sourcing.proto
//go:generate protoc --go_out=plugins=grpc:. call_aggregate.proto
//go:generate protoc --go_out=plugins=grpc:. call_event_handler.proto

var conn *grpc.ClientConn
var sourceClient SourcingClient
var aggregateClient CallAggregateClient
var eventHandlerClient CallEventHandlerClient

func Start(ctx context.Context) error {
	var e error
	go func() {
		conn, e = grpc.Dial(fmt.Sprintf("localhost:%s", config.Instance.Rpc.Port), grpc.WithInsecure())
		if e != nil {
			logrus.Errorf("[rpc]连接业务系统失败,错误:%v", e.Error())
			return
		}
		sourceClient = NewSourcingClient(conn)
		aggregateClient = NewCallAggregateClient(conn)
		eventHandlerClient = NewCallEventHandlerClient(conn)
	}()
	defer func() {
		if conn != nil {
			e = conn.Close()
			if e != nil {
				logrus.Errorf("[rpc]关闭conn失败,错误:%v", e.Error())
			}
		}
	}()
	return e
}

func Sourcing(ctx context.Context, id string, aggType string, agg []byte, events []*event.Event) ([]byte, error) {
	sourceEvents := make([]*SourcingRequestEvent, len(events))

	for key, value := range events {
		//TODO:检查时间是否正确
		create := &timestamp.Timestamp{
			Seconds: int64(value.Create.Second()),
			Nanos:   int32(value.Create.UnixNano()),
		}
		sourceEvents[key] = &SourcingRequestEvent{
			Id:        value.Id,
			EventType: value.EventType,
			Create:    create,
			Data:      value.Data,
		}
	}
	response, e := sourceClient.Sourcing(ctx, &SourcingRequest{
		AggregateId:   id,
		AggregateType: aggType,
		Aggregate:     agg,
		Events:        sourceEvents,
	})
	if e != nil {
		logrus.Errorf("[rpc]方法[Sourcing]调用发生错误,错误:%v", e.Error())
		return nil, e
	}
	return response.Aggregate.Value, nil
}

func CallAggregate(ctx context.Context, aggId string, aggType string, agg []byte, cmd *command.Command) ([]*CallAggregateResponseEvent, error) {
	callCmd := &Command{
		CmdType: cmd.CmdType,
		Data:    cmd.Data,
	}
	response, e := aggregateClient.CallAggregate(ctx, &CallAggregateRequest{
		AggregateId:   aggId,
		AggregateType: aggType,
		Aggregate:     agg,
		Command:       callCmd,
	})
	if e != nil {
		logrus.Errorf("[rpc]方法[CallAggregate]调用发生错误,错误:%v", e.Error())
		return nil, e
	}
	return response.Events, nil
}

func CallEventHandler(ctx context.Context, events []*event.Event) ([]*CallEventHandlerResponseMark, error) {
	requestEvents := make([]*CallEventHandlerRequestEvent, len(events))

	for key, value := range events {
		create := &timestamp.Timestamp{
			Seconds: int64(value.Create.Second()),
			Nanos:   int32(value.Create.UnixNano()),
		}
		requestEvents[key] = &CallEventHandlerRequestEvent{
			Id:        value.Id,
			EventType: value.EventType,
			AggId:     value.AggId,
			AggType:   value.AggType,
			Create:    create,
			Data:      value.Data,
		}
	}
	response, e := eventHandlerClient.CallEventHandler(ctx, &CallEventHandlerRequest{
		Events: requestEvents,
	})
	if e != nil {
		logrus.Errorf("[rpc]方法[CallEventHandler]调用发生错误,错误:%v", e.Error())
		return nil, e
	}
	return response.Data, nil
}
