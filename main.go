package main

import (
	"fmt"
	"github.com/smtdfc/photon-cli/commands"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:  "photon-cli",
		Usage: "CLI for Photon Framework",
		Commands: []*cli.Command{
			{
				Name:  "version",
				Usage: "Display current version",
				Action: func(c *cli.Context) error {
					fmt.Println("v1.0.1")
					return nil
				},
			},
			{
				Name:  "dev",
				Usage: "Start app in development mode",
				Flags: []cli.Flag{
					&cli.IntFlag{},
				},
				Action: commands.Dev,
			},
			{
				Name:  "init",
				Usage: "Create new project",
				Flags: []cli.Flag{
					&cli.IntFlag{},
				},
				Action: commands.Init,
			},
			{
				Name:  "gen",
				Usage: "Generate module/components for project",
				Flags: []cli.Flag{
					&cli.IntFlag{},
				},
				Action: commands.Gen,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
