package db

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/tangxusc/cavy-sidecar/pkg/config"
	"time"
)

var dbConn *sqlx.DB

const DbType = "mysql"

func InitConn(ctx context.Context) error {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s?charset=utf8&parseTime=true", config.Instance.Db.Username, config.Instance.Db.Password,
		"tcp", config.Instance.Db.Address, config.Instance.Db.Port, config.Instance.Db.Database)
	var err error
	dbConn, err = sqlx.Open(DbType, dsn)
	if err != nil {
		logrus.Errorf("[db]连接数据库发生错误:%v", err)
		return err
	}
	dbConn.SetConnMaxLifetime(time.Duration(config.Instance.Db.LifeTime) * time.Second)
	dbConn.SetMaxOpenConns(config.Instance.Db.MaxOpen)
	dbConn.SetMaxIdleConns(config.Instance.Db.MaxIdle)
	return nil
}

func QueryRow(sqlString string, data interface{}, param ...interface{}) error {
	rowx := dbConn.QueryRowx(sqlString, param...)
	err := rowx.StructScan(data)
	switch {
	case err == sql.ErrNoRows:
		return nil
	case err != nil:
		logrus.Errorf("[db]执行sql出现错误,sql:%v,param:%v,错误:%v", sqlString, param, err)
		return err
	}
	return nil
}

func Query(sqlString string, f func() interface{}, param ...interface{}) error {
	rows, e := dbConn.Queryx(sqlString, param...)
	if e != nil {
		logrus.Errorf("[db]执行sql出现错误,sql:%v,param:%v,错误:%v", sqlString, param, e)
		return e
	}
	defer rows.Close()
	for rows.Next() {
		e := rows.StructScan(f())
		if e != nil {
			logrus.Errorf("[db]执行sql出现错误,sql:%v,param:%v,错误:%v", sqlString, param, e)
			return e
		}
	}
	return nil
}

func NameExec(sqlString string, data interface{}) (result sql.Result, e error) {
	tx, e := dbConn.Beginx()
	if e != nil {
		logrus.Errorf("[db]开始事务出现错误,sql:%v,参数:%v,错误:%v", sqlString, data, e)
		return
	}
	result, e = tx.NamedExec(sqlString, data)
	if e != nil {
		logrus.Errorf("[db]执行sql出现错误,sql:%v,参数:%v,错误:%v", sqlString, data, e)
		e = tx.Rollback()
		if e != nil {
			logrus.Errorf("[db]回滚事务出现错误,sql:%v,参数:%v,错误:%v", sqlString, data, e)
		}
		return
	}
	e = tx.Commit()
	if e != nil {
		logrus.Errorf("[db]提交事务出现错误,sql:%v,参数:%v,错误:%v", sqlString, data, e)
	}
	return
}

func Transaction(f func(tx *sqlx.Tx) error) {
	tx, e := dbConn.Beginx()
	if e != nil {
		logrus.Errorf("[db]开始事务出现错误,错误:%v", e)
		return
	}
	e = f(tx)
	if e != nil {
		logrus.Errorf("[db]执行函数f出现错误,f:%v,错误:%v", f, e)
		e = tx.Rollback()
		if e != nil {
			logrus.Errorf("[db]回滚事务出现错误,f:%v,错误:%v", f, e)
		}
		return
	}
	e = tx.Commit()
	if e != nil {
		logrus.Errorf("[db]提交事务出现错误,f:%v,错误:%v", f, e)
	}
}
