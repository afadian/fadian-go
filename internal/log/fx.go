package log

import (
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func FX(p DebugParam, logger *zap.Logger) fxevent.Logger {
	if !p.Debug {
		return &fxevent.NopLogger
	}

	return &fxevent.ZapLogger{
		Logger: logger,
	}
}
