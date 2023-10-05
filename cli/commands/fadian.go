package commands

import (
	"bytes"
	"fmt"
	"github.com/afadian/fadian-go/data/sections"
	"github.com/samber/lo"
	"github.com/urfave/cli/v2"
	"github.com/vinta/pangu"
	"go.uber.org/zap"
	"strings"
	"text/template"
)

func HandleFadian() (*cli.Command, error) {
	tpl, err := template.New("fadian").Parse(sections.FadianText)
	if err != nil {
		zap.L().Error("failed to parse fadian template", zap.Error(err))
		return nil, err
	}

	type FadianParams struct {
		Name string
	}

	return &cli.Command{
		Name:  "fadian",
		Usage: "发癫",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "name",
				Aliases:  []string{"n"},
				Usage:    "对谁发癫",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			buf := bytes.NewBuffer([]byte{})
			if err := tpl.Execute(buf, FadianParams{
				Name: c.String("name"),
			}); err != nil {
				zap.L().Error("failed to execute fadian template", zap.Error(err))
				return err
			}

			str := pangu.SpacingText(buf.String())

			fmt.Println(strings.Join(lo.RepeatBy(c.Int("num"), func(_ int) string {
				return str
			}), "\n"))

			return nil
		},
	}, nil
}
