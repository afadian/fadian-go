package log

import "go.uber.org/fx"

func Module() fx.Option {
	return fx.Module(
		"logger",
		fx.Provide(fx.Annotate(NewLogger, fx.ParamTags(`name:"debug"`))),
		fx.Invoke(ConfigureLogger),
	)
}
