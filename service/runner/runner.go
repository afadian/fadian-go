package runner

import (
	"context"
	"fmt"

	"go.uber.org/fx"

	"github.com/afadian/fadian-go/data"
	"github.com/afadian/fadian-go/service/fabing"
	"github.com/afadian/fadian-go/service/fadian"
)

type Service interface {
	Run(ctx context.Context) error
}

type ServiceImpl struct {
	fx.In

	BootstrapParams *data.BootstrapParams
	FabingService   fabing.Service
	FadianService   fadian.Service
}

func NewService(svc ServiceImpl) Service {
	return &svc
}

func (svc *ServiceImpl) Run(_ context.Context) error {
	if svc.BootstrapParams.IsInteractive {
		fmt.Print("请输入姓名：")
		if _, err := fmt.Scanln(&svc.BootstrapParams.Name); err != nil {
			return err
		}
	}

	if svc.BootstrapParams.IsFabing {
		fmt.Println(svc.FabingService.Fabing(svc.BootstrapParams.Name, svc.BootstrapParams.Num))
	} else {
		fmt.Println(svc.FadianService.Fadian(svc.BootstrapParams.Name, svc.BootstrapParams.Num))
	}

	return nil
}
