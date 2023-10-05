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
			name := c.String("name")
			if c.Bool("interactive") {
				fmt.Print("对谁发癫？")
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
			if err := tpl.Execute(buf, FadianParams{
				Name: name,
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
