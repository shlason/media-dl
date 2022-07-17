package cli

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/shlason/media-dl/pkgs/configs"
	"github.com/shlason/media-dl/pkgs/parser"
	"github.com/shlason/media-dl/pkgs/utils"
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
	mediaType := utils.Url.GetYoutubeType(url)
	if mediaType == configs.MediaTypes.YoutubeVideo {
		p := parser.YoutubeParser{Url: url}
		p.Parse()
		fmt.Println(p)
		res, err := http.Get("https://www.youtube.com/get_video_info?video_id=" + p.Vid)
		if err != nil {
			log.Fatalln("Get youtube video info failed")
		}
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatalln("Read youtube response body failed")
		}
		fmt.Println(res.StatusCode, body)
	}
	return nil
}

func GetInstance() *cli.App {
	return app
}
