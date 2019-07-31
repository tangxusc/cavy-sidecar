package model

import "time"

type HandlerStatus int

const (
	//未处理
	Untreated HandlerStatus = iota
	//已处理
	Processed
)

type MqStatus int

const (
	//未发送
	MqNotSend MqStatus = iota
	//已发送
	MqSent
)

type Event struct {
	Id            string        `db:"id"`
	EventType     string        `db:"event_type"`
	AggId         string        `db:"agg_id"`
	AggType       string        `db:"agg_type"`
	Create        time.Time     `db:"create_time"`
	Data          []byte        `db:"data"`
	HandlerStatus HandlerStatus `db:"handler_status"`
	MqStatus      MqStatus      `db:"mq_status"`
}
