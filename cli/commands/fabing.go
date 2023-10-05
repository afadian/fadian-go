package commands

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
	"text/template"

	"github.com/samber/lo"
	"github.com/urfave/cli/v2"
	"github.com/vinta/pangu"
	"go.uber.org/zap"

	"github.com/afadian/fadian-go/data/sections"
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
		Name:  "fabing",
		Usage: "发病",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "name",
				Aliases: []string{"n"},
				Usage:   "对谁发病",
				Value:   "",
			},
			&cli.BoolFlag{
				Name:    "interactive",
				Aliases: []string{"i"},
				Usage:   "是否交互式发病",
				Value:   false,
			},
		},
		Action: func(c *cli.Context) error {
			name := c.String("name")
			if c.Bool("interactive") {
				fmt.Print("对谁发病？")
				if _, err := fmt.Scan(&name); err != nil {
					zap.L().Error("failed to scan name", zap.Error(err))
					return err
				}
			}

			if name == "" {
				zap.L().Error("name is empty")
				return errors.New("name is empty")
			}

			buf := bytes.NewBuffer([]byte{})
			if err := tpl.Execute(buf, FabingParams{
				Name: name,
			}); err != nil {
				zap.L().Error("failed to execute fabing template", zap.Error(err))
				return err
			}

			str := pangu.SpacingText(buf.String())
			zap.L().Debug("fabing text generated", zap.String("text", str))

			fmt.Println(strings.Join(lo.RepeatBy(c.Int("num"), func(_ int) string {
				return str
			}), "\n"))

			return nil
		},
	}, nil
}
