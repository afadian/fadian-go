package log

import (
	"github.com/afadian/fadian-go/data"
	"go.uber.org/zap"
)

func NewLogger(params *data.BootstrapParams) *zap.Logger {
	var logger *zap.Logger
	var err error
	if params.Debug {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = zap.NewProduction()
	}

	if err != nil {
		zap.L().Panic("initial logger failed", zap.Error(err))
	}

	defer logger.Sync()

	return logger
}
