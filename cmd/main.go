package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/tangxusc/cavy-sidecar/pkg/command"
	"os"
	"os/signal"
)

func main() {
	signals := make(chan os.Signal, 1)
	ctx, cancel := context.WithCancel(context.Background())
	newCommand := command.NewCommand(ctx)
	go func() {
		signal.Notify(signals, os.Interrupt, os.Kill)
	}()
	if err := newCommand.Execute(); err != nil {
		logrus.Errorf("发生了错误,错误:%v", err.Error())
	}
	<-signals
	cancel()
}
