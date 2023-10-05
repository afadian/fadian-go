package main

import (
	"context"
	"go.uber.org/fx"

	"github.com/afadian/fadian-go/data"
	"github.com/afadian/fadian-go/entry"
)

var (
	ctx = context.Background()
	app *fx.App

	version = "0.0.1"
	commit  = "0000000"
)

func init() {
	app = entry.Initialize(ctx, &data.BootstrapParams{
		Version: version,
		Commit:  commit,
	})
}

func main() {
	if err := app.Start(ctx); err != nil {
		panic(err)
	}

	app.Wait()
}
