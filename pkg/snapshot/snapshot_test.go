package snapshot

import (
	"context"
	"fmt"
	"github.com/tangxusc/cavy-sidecar/pkg/config"
	"github.com/tangxusc/cavy-sidecar/pkg/db"
	"testing"
	"time"
)

var aggregate = []byte{'t', 'e', 's', 't'}

func TestFindLastSnapBy(t *testing.T) {
	initTestDb()
	agg, err := FindLastSnapBy("1", "Test")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(agg)
}

func TestSaveSnapshot(t *testing.T) {
	initTestDb()
	agg := &Aggregate{
		Id:         "1",
		AggType:    "Test",
		AggId:      "1",
		CreateTime: time.Now(),
		Data:       aggregate,
	}
	SaveSnapshot(agg)
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
