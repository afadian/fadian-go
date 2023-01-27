package main

import (
	"context"
	"flag"
	"github.com/afadian/fadian-go/data"

	"github.com/afadian/fadian-go/entry"
	"go.uber.org/fx"
)

var (
	app *fx.App
	ctx = context.Background()
)

func init() {
	var params data.BootstrapParams
	flag.StringVar(&params.Name, "name", "", "发癫对象")
	flag.BoolVar(&params.IsInteractive, "i", false, "是否进入交互模式")
	flag.BoolVar(&params.IsFabing, "f", false, "是否进入发病模式")
	flag.IntVar(&params.Num, "num", 1, "发病/发癫次数")
	flag.BoolVar(&params.Debug, "d", false, "是否开启调试模式")
	flag.Parse()

	app = entry.Initialize(ctx, &params)
}

func main() {
	app.Run()
}
