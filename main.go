package main

import (
	"fmt"
	"os"

	"github.com/m-mizutani/zlog"
	cli "github.com/urfave/cli/v2"
)

var logger = zlog.New()

type App struct {
	app *cli.App
}

func NewApp() *App {
	var cfg config
	return &App{
		app: &cli.App{
			Name: "gomodtrace",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:        "input",
					Aliases:     []string{"i"},
					Value:       "-",
					Destination: &cfg.Input,
					Usage:       "Input file (use '-' for stdin)",
				},
				&cli.StringFlag{
					Name:        "format",
					Aliases:     []string{"f"},
					Value:       "tree",
					Destination: &cfg.Output,
					Usage:       "Output format",
				},
			},
			Action: func(c *cli.Context) error {
				if err := cfg.Setup(); err != nil {
					return err
				}
				defer cfg.Teardown()

				deps, err := read(cfg.reader)
				if err != nil {
					return err
				}

				dependMap := NewDependMap(deps)

				for _, tgt := range c.Args().Slice() {
					if err := cfg.out(os.Stdout, dependMap.Trace(tgt)); err != nil {
						return err
					}
				}

				return nil
			},
		},
	}
}

func (x *App) Run(args []string) {
	if err := x.app.Run(args); err != nil {
		fmt.Println("Error:", err)
	}
}

func main() {
	NewApp().Run(os.Args)
}
