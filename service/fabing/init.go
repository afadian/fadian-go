package fabing

import (
	"github.com/afadian/fadian-go/data/sections"
)

type BaseText string

func Initialize() BaseText {
	return BaseText(sections.FabingText)
}
