package utils

import (
	"regexp"

	"github.com/shlason/media-dl/pkgs/configs"
)

type url struct {
	GetYoutubeType func(string) string
}

var validYoutubeVideoUrl = regexp.MustCompile(`http(?:s?):\/\/(?:www\.)?youtu(?:be\.com\/watch\?v=|\.be\/)([\w\-\_]*)(&(amp;)?‌​[\w\?‌​=]*)?`)

func getYoutubeType(url string) string {
	if b := validYoutubeVideoUrl.MatchString(url); b {
		return configs.MediaTypes.YoutubeVideo
	}
	return ""
}

var Url = url{
	GetYoutubeType: getYoutubeType,
}
