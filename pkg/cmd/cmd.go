package cmd

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tangxusc/cavy-sidecar/pkg/command"
	"github.com/tangxusc/cavy-sidecar/pkg/config"
	"github.com/tangxusc/cavy-sidecar/pkg/db"
	"github.com/tangxusc/cavy-sidecar/pkg/event"
	"github.com/tangxusc/cavy-sidecar/pkg/snapshot"
	"math/rand"
	"os"
	"os/signal"
	"time"
)

func NewCmd(ctx context.Context) *cobra.Command {
	var command = &cobra.Command{
		Use:   "start",
		Short: "start sidecar",
		RunE: func(cmd *cobra.Command, args []string) error {
			rand.Seed(time.Now().Unix())
			//0,日志
			config.InitLog()
			e := db.InitConn(ctx)
			if e != nil {
				return e
			}

			command.Listen(ctx)
			snapshot.Listen(ctx)
			event.Listen(ctx)

			return nil
		},
	}
	logrus.SetFormatter(&logrus.TextFormatter{})
	config.BindParameter(command)

	return command
}

func HandlerNotify(cancel context.CancelFunc) {
	go func() {
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, os.Interrupt, os.Kill)
		<-signals
		cancel()
	}()
}
