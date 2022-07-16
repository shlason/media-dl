package main

import (
	"os"

	"github.com/shlason/media-dl/pkgs/cli"
)

func main() {
	app := cli.GetInstance()
	app.Run(os.Args)
}
