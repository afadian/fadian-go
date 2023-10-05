package entry

import (
	"context"

	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/afadian/fadian-go/data"
	"github.com/afadian/fadian-go/internal/log"
)

func Initialize(ctx context.Context, params *data.BootstrapParams) *fx.App {
	opts := []fx.Option{
		fx.Supply(
			fx.Annotate(ctx, fx.As(new(context.Context))),
			fx.Annotate(params.Version, fx.ResultTags(`name:"version"`)),
			fx.Annotate(params.Commit, fx.ResultTags(`name:"commit"`)),
			fx.Annotate(params.Debug, fx.ResultTags(`name:"debug"`)),
		),
	}

	opts = append(opts, AppEntry()...)

	if !params.Debug {
		opts = append(opts, fx.NopLogger)
	} else {
		zap.L().Debug("using debug mode")
		opts = append(opts, fx.WithLogger(log.FX))
	}

	return fx.New(opts...)
}
