package service

import (
	"go.uber.org/fx"

	"github.com/afadian/fadian-go/service/fabing"
	"github.com/afadian/fadian-go/service/fadian"
	"github.com/afadian/fadian-go/service/runner"
)

func Module() fx.Option {
	return fx.Module("service",
		fx.Provide(fadian.Initialize),
		fx.Provide(fadian.NewService),

		fx.Provide(fabing.Initialize),
		fx.Provide(fabing.NewService),

		fx.Provide(runner.NewService),
	)
}
