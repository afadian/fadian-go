package fadian

import (
	"strings"

	"github.com/samber/lo"
	"github.com/vinta/pangu"
	"go.uber.org/fx"

	"github.com/afadian/fadian-go/internal/util"
)

type Service interface {
	Fadian(name string, repeat int) string
}

type ServiceImpl struct {
	fx.In

	BaseText BaseText
}

var _ Service = (*ServiceImpl)(nil)

func NewService(svc ServiceImpl) Service {
	return &svc
}

func (s *ServiceImpl) Fadian(name string, num int) string {
	str := util.ReplaceString(string(s.BaseText), map[string]string{
		"$name": name,
	})

	str = pangu.SpacingText(str)

	return strings.Join(lo.RepeatBy(num, func(_ int) string {
		return str
	}), "\n")
}
