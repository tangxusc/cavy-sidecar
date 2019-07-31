package model

import (
	"context"
	"time"
)

type Command struct {
	CmdType       string
	AggregateId   string
	AggregateType string
	Data          []byte
}

type Aggregate interface {
	Listen(ctx context.Context, reset func())
	SendCommand(command *Command)
	GetKey() string
	SendEvent(event *Event)
}

type Sourcing struct {
	Key string
	//cmd命令
	CommandChan chan *Command
	//上下文
	Ctx context.Context
	//聚合对象
	Aggregate []byte
	//最后更新时间
	LastTime time.Time
	//eventHandler
	EventChan chan []*Event
}
