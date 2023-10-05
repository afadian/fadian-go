package commands

import "go.uber.org/fx"

func Module() fx.Option {
	return fx.Module("cli.commands",
		fx.Provide(HandleNotFound),
		fx.Provide(fx.Annotate(HandleFabing, fx.ResultTags(`group:"root_commands"`))),
		fx.Provide(fx.Annotate(HandleFadian, fx.ResultTags(`group:"root_commands"`))),
	)
}
