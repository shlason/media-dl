package parser

import (
	"log"
	"net/url"
)

type YoutubeParser struct {
	Url string
	Vid string
}

func (yp *YoutubeParser) Parse() {
	url, err := url.Parse(yp.Url)
	if err != nil {
		log.Fatalln(err)
	}
	yp.Vid = url.Query()["v"][0]
}
