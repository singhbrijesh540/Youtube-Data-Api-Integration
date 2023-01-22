package youtubeService

import "google.golang.org/api/youtube/v3"

type YoutubeDataApiService interface {
	GetDataFromYoutubeDataApi() (response *youtube.SearchListResponse, err error)
}
