package entry

import (
	"context"
	"github.com/afadian/fadian-go/service/runner"
	"os"

	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/afadian/fadian-go/data"
	"github.com/afadian/fadian-go/internal/log"
)

func Initialize(ctx context.Context, params *data.BootstrapParams) *fx.App {
	opts := []fx.Option{
		fx.Supply(
			fx.Annotate(ctx, fx.As(new(context.Context))),
			params,
		),
	}

	opts = append(opts, AppEntry()...)

	if !params.Debug {
		opts = append(opts, fx.NopLogger)
	} else {
		zap.L().Debug("using debug mode")
		opts = append(opts, fx.WithLogger(log.FX))
	}

	opts = append(opts, fx.Invoke(run))

	return fx.New(opts...)
}

func run(_ context.Context, runner runner.Service, lc fx.Lifecycle) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			if err := runner.Run(ctx); err != nil {
				zap.L().Error("run failed", zap.Error(err))
				return err
			}

			os.Exit(0)
			return nil
		},
	})
}
