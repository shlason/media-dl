package configs

type mediaTypes struct {
	YoutubeVideo     string
	YoutubeThumbnail string
}

var MediaTypes = mediaTypes{
	YoutubeVideo:     "youtube_video",
	YoutubeThumbnail: "youtube_thumbnail",
}
