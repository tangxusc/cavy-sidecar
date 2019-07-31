package model

import "time"

type AggregateSnapshot struct {
	Id         string    `db:"id"`
	AggType    string    `db:"agg_type"`
	AggId      string    `db:"agg_id"`
	CreateTime time.Time `db:"create_time"`
	Data       []byte    `db:"data"`
}
