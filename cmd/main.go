package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/tangxusc/cavy-sidecar/pkg/cmd"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	newCommand := cmd.NewCmd(ctx)
	cmd.HandlerNotify(cancel)

	if err := newCommand.Execute(); err != nil {
		logrus.Errorf("发生了错误,错误:%v", err.Error())
	}
}
