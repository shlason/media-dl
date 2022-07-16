package configs

import (
	ucli "github.com/urfave/cli"
)

type cli struct {
	Name  string
	Usage string
	Flags []ucli.Flag
}

var Cli = cli{
	Name:  "media-dl",
	Usage: "Download any kinds of media from social media website.",
	Flags: []ucli.Flag{
		ucli.StringFlag{
			Name:  "url, u",
			Value: "",
			Usage: "Target's download URL",
		},
	},
}
