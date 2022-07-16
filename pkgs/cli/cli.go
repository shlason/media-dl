package cli

import (
	"fmt"

	"github.com/shlason/media-dl/pkgs/configs"
	"github.com/urfave/cli"
)

var app *cli.App

func init() {
	app = cli.NewApp()
	app.Name = configs.Cli.Name
	app.Usage = configs.Cli.Usage
	app.Flags = configs.Cli.Flags
	app.Action = action
}

func action(c *cli.Context) error {
	url := c.GlobalString("url")
	fmt.Printf("Hello %s!\n", url)
	return nil
}

func GetInstance() *cli.App {
	return app
}
