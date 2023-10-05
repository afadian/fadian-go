package cli

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"os"
	"time"

	"github.com/urfave/cli/v2"
	"go.uber.org/fx"
)

type Params struct {
	fx.In
	Version string `name:"version"`
	Commit  string `name:"commit"`

	Commands       []*cli.Command `group:"root_commands"`
	HandleNotFound cli.CommandNotFoundFunc
}

func NewApp(params Params) *cli.App {
	app := &cli.App{
		Name:      "fadian-go",
		Usage:     "发癫",
		Version:   fmt.Sprintf("%s-%s", params.Version, params.Commit),
		Copyright: fmt.Sprintf("2022-%d © 爱发癫 All Rights Reserved", time.Now().Year()),
		Authors: []*cli.Author{
			{Name: "AH Dark", Email: "ahdark@outlook.com"},
			{Name: "Kevin Williams", Email: "admin@utermux.dev"},
			{Name: "Purofle", Email: "purofle@gmail.com"},
		},
		EnableBashCompletion: true,
		CommandNotFound:      params.HandleNotFound,
		Commands:             params.Commands,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "debug",
				Hidden:  true,
				Usage:   "是否开启调试模式",
				EnvVars: []string{"DEBUG"},
			},
			&cli.IntFlag{
				Name:  "num",
				Usage: "重复次数",
				Value: 1,
			},
		},
	}

	app.Setup()

	zap.L().Debug("app initialized")

	return app
}

func InvokeApp(app *cli.App, lc fx.Lifecycle) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			if err := app.RunContext(ctx, os.Args); err != nil {
				zap.L().Warn("app run failed", zap.Error(err))
			}

			return nil
		},
	})
}
