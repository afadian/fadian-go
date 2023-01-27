package log

import "go.uber.org/zap"

func ConfigureLogger(logger *zap.Logger) {
	zap.ReplaceGlobals(logger)
}
