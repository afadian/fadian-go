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

func HandleFabing() (*cli.Command, error) {
	tpl, err := template.New("fabing").Parse(sections.FabingText)
	if err != nil {
		zap.L().Error("failed to parse fabing template", zap.Error(err))
		return nil, err
	}

	zap.L().Debug("fabing template parsed", zap.String("template", sections.FabingText))

	type FabingParams struct {
		Name string
	}

	return &cli.Command{
		Name: "fabing",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "name",
				Aliases:  []string{"n"},
				Usage:    "name to be fabing",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			buf := bytes.NewBuffer([]byte{})
			if err := tpl.Execute(buf, FabingParams{
				Name: c.String("name"),
			}); err != nil {
				zap.L().Error("failed to execute fabing template", zap.Error(err))
				return err
			}

			str := pangu.SpacingText(buf.String())
			zap.L().Debug("fabing text generated", zap.String("text", str))

			fmt.Sprintln(strings.Join(lo.RepeatBy(c.Int("num"), func(_ int) string {
				return str
			}), "\n"))

			return nil
		},
	}, nil
}
