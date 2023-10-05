package entry

import (
	"context"
	"os"

	"github.com/afadian/fadian-go/data"
	"github.com/afadian/fadian-go/internal/log"
	"go.uber.org/fx"
)

func Initialize(ctx context.Context, params *data.BootstrapParams) *fx.App {
	opts := []fx.Option{
		fx.Supply(
			fx.Annotate(ctx, fx.As(new(context.Context))),
			fx.Annotate(params.Version, fx.ResultTags(`name:"version"`)),
			fx.Annotate(params.Commit, fx.ResultTags(`name:"commit"`)),
			fx.Annotate(os.Getenv("DEBUG") == "true", fx.ResultTags(`name:"debug"`)),
		),
	}

	opts = append(opts, AppEntry()...)
	opts = append(opts, fx.WithLogger(log.FX))

	return fx.New(opts...)
}
