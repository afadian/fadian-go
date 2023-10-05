package commands

import (
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

func HandleNotFound() cli.CommandNotFoundFunc {
	return func(context *cli.Context, command string) {
		zap.L().Error("command not found", zap.String("command", command))
	}
}
