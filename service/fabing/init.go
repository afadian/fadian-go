package fabing

import (
	"os"

	"go.uber.org/zap"

	"github.com/afadian/fadian-go/internal/util"
)

type BaseText string

func Initialize() (BaseText, error) {
	f, err := os.ReadFile(util.AbsolutePath("data/sections/fabing.txt"))
	if err != nil {
		zap.L().Error("read fabing.txt failed", zap.Error(err))
		return "", err
	}

	return BaseText(f), nil
}
