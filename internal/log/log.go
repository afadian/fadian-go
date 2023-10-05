package log

import (
	"context"

	"github.com/samber/lo"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type DebugParam struct {
	fx.In
	Debug bool `name:"debug"`
}

func NewLogger(p DebugParam, lc fx.Lifecycle) (*zap.Logger, error) {
	logger, err := lo.If(p.Debug, zap.NewDevelopment).Else(zap.NewProduction)()
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
