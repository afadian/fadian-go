package entry

import (
	"go.uber.org/fx"

	"github.com/afadian/fadian-go/cli"
	"github.com/afadian/fadian-go/internal/log"
)

func AppEntry() []fx.Option {
	return []fx.Option{
		log.Module(),
		cli.Module(),
	}
}
