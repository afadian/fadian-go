package cli

import (
	"github.com/afadian/fadian-go/cli/commands"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Options(
		commands.Module(),
		fx.Provide(NewApp),
		fx.Invoke(InvokeApp),
	)
}
