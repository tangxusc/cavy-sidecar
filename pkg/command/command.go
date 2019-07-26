package command

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tangxusc/cavy-sidecar/pkg/config"
)

func NewCommand(ctx context.Context) *cobra.Command {
	var command = &cobra.Command{
		Use:   "start",
		Short: "start sidecar",
		RunE: func(cmd *cobra.Command, args []string) error {
			//0,日志
			config.InitLog()

			return nil
		},
	}
	logrus.SetFormatter(&logrus.TextFormatter{})
	config.BindParameter(command)

	return command
}
