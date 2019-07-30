package event

import (
	"context"
	"fmt"
	"github.com/tangxusc/cavy-sidecar/pkg/config"
	"github.com/tangxusc/cavy-sidecar/pkg/db"
	"testing"
	"time"
)

var aggregate = []byte{'t', 'e', 's', 't'}

func TestSave(t *testing.T) {
	initTestDb()

	events := make([]*Event, 0, 2)
	create := time.Now()
	events = append(events, &Event{
		Id:      "1",
		AggId:   "1",
		AggType: "Test",
		Create:  create,
		Data:    aggregate,
	})
	events = append(events, &Event{
		Id:      "2",
		AggId:   "1",
		AggType: "Test",
		Create:  create,
		Data:    aggregate,
	})
	Save(events)
}

func initTestDb() {
	config.Instance.Db = &config.DbConfig{
		Address:  "127.0.0.1",
		Port:     "3306",
		Database: "test",
		Username: "root",
		Password: "123456",
		LifeTime: 10,
		MaxOpen:  2,
		MaxIdle:  2,
	}
	err := db.InitConn(context.TODO())
	if err != nil {
		panic(err.Error())
	}
}

func TestFindEventByTime(t *testing.T) {
	initTestDb()
	create := time.Now().Add(time.Hour * -1)
	events, e := FindEventByTime("1", "Test", create)
	if e != nil {
		panic(e.Error())
	}
	fmt.Println(len(events))
}
