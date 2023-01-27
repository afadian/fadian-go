package fadian

import (
	"github.com/afadian/fadian-go/data/sections"
)

type BaseText string

func Initialize() BaseText {
	return BaseText(sections.FadianText)
}
