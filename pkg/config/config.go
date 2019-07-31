package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const debugArgName = "debug"

func InitLog() {
	if viper.GetBool(debugArgName) {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.SetReportCaller(true)
		logrus.Debug("已开启debug模式...")
	} else {
		logrus.SetLevel(logrus.WarnLevel)
	}

	Instance.Debug = viper.GetBool(debugArgName)
}

func BindParameter(cmd *cobra.Command) {
	viper.SetEnvPrefix("SIDECAR")
	viper.AutomaticEnv()

	cmd.PersistentFlags().BoolVarP(&Instance.Debug, debugArgName, "v", false, "debug mod")
	cmd.PersistentFlags().StringVarP(&Instance.Db.Address, "db-address", "dba", "", "数据库连接地址")
	cmd.PersistentFlags().StringVarP(&Instance.Db.Port, "db-port", "dbp", "3306", "数据库端口")
	cmd.PersistentFlags().StringVarP(&Instance.Db.Port, "db-port", "dbp", "3306", "数据库端口")
	cmd.PersistentFlags().StringVarP(&Instance.Db.Database, "db-Database", "dbd", "", "数据库实例")
	cmd.PersistentFlags().StringVarP(&Instance.Db.Username, "db-Username", "dbu", "", "数据库用户名")
	cmd.PersistentFlags().StringVarP(&Instance.Db.Password, "db-Password", "dbp", "", "数据库密码")
	cmd.PersistentFlags().IntVarP(&Instance.Db.LifeTime, "db-LifeTime", "dbl", 10, "数据库连接最大连接周期(秒)")
	cmd.PersistentFlags().IntVarP(&Instance.Db.MaxOpen, "db-MaxOpen", "dbo", 5, "数据库最大连接数")
	cmd.PersistentFlags().IntVarP(&Instance.Db.MaxIdle, "db-MaxIdle", "dbi", 5, "数据库最大等待数量")

	cmd.PersistentFlags().IntVarP(&Instance.Db.LifeTime, "agg-LifeTime", "aggl", 10, "聚合对象存活时间(分钟)")
	cmd.PersistentFlags().IntVarP(&Instance.Db.LifeTime, "agg-RecoverTime", "aggr", 5, "事件故障恢复时间")

	cmd.PersistentFlags().StringVarP(&Instance.Rpc.Port, "rpc-port", "rpcp", "9999", "rpc端口")
	_ = viper.BindPFlag(debugArgName, cmd.PersistentFlags().Lookup(debugArgName))
	_ = viper.BindPFlag("db-address", cmd.PersistentFlags().Lookup("db-address"))
	_ = viper.BindPFlag("db-port", cmd.PersistentFlags().Lookup("db-port"))
	_ = viper.BindPFlag("db-Database", cmd.PersistentFlags().Lookup("db-Database"))
	_ = viper.BindPFlag("db-Username", cmd.PersistentFlags().Lookup("db-Username"))
	_ = viper.BindPFlag("db-Password", cmd.PersistentFlags().Lookup("db-Password"))
	_ = viper.BindPFlag("db-LifeTime", cmd.PersistentFlags().Lookup("db-LifeTime"))
	_ = viper.BindPFlag("db-MaxOpen", cmd.PersistentFlags().Lookup("db-MaxOpen"))
	_ = viper.BindPFlag("db-MaxIdle", cmd.PersistentFlags().Lookup("db-MaxIdle"))

	_ = viper.BindPFlag("agg-LifeTime", cmd.PersistentFlags().Lookup("agg-LifeTime"))
	_ = viper.BindPFlag("agg-RecoverTime", cmd.PersistentFlags().Lookup("agg-RecoverTime"))

	_ = viper.BindPFlag("rpc-port", cmd.PersistentFlags().Lookup("rpc-port"))
}

type AggregateConfig struct {
	LifeTime    int
	RecoverTime int
}

type RpcConfig struct {
	Port string
}

type Config struct {
	Debug     bool
	Db        *DbConfig
	Aggregate *AggregateConfig
	Rpc       *RpcConfig
}

type DbConfig struct {
	Address  string
	Port     string
	Database string
	Username string
	Password string

	LifeTime int
	MaxOpen  int
	MaxIdle  int
}

var Instance = &Config{
	Db:        &DbConfig{},
	Aggregate: &AggregateConfig{},
	Rpc:       &RpcConfig{},
}
