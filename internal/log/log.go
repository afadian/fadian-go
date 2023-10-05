package log

import (
	"context"
	"github.com/samber/lo"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewLogger(debug bool, lc fx.Lifecycle) (*zap.Logger, error) {
	logger, err := lo.If(debug, zap.NewDevelopment).Else(zap.NewProduction)()
	if err != nil {
		zap.L().Panic("initial logger failed", zap.Error(err))
		return nil, err
	}

	lc.Append(fx.Hook{
		OnStop: func(context.Context) error {
			return logger.Sync()
		},
	})

	return logger, nil
}
