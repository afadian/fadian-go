package entry

import (
	"go.uber.org/fx"

	"github.com/afadian/fadian-go/internal/log"
	"github.com/afadian/fadian-go/service"
)

func AppEntry() []fx.Option {
	return []fx.Option{
		log.Module(),
		service.Module(),
	}
}
