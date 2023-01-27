package log

import (
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func FX(logger *zap.Logger) fxevent.Logger {
	return &fxevent.ZapLogger{
		Logger: logger,
	}
}
