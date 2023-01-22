package youtubeService

import (
	"github.com/singhbrijesh540/fampay-assignment/config"
	"google.golang.org/api/youtube/v3"
)

type youtubeDataApiService struct {
	youtubeService   *youtube.Service
	YoutubeApiConfig config.YoutubeApiConfig
}

func (service youtubeDataApiService) GetDataFromYoutubeDataApi() (response *youtube.SearchListResponse, err error) {
	// Make the API call to YouTube Data v3 Api.
	call := service.youtubeService.Search.List([]string{"id", "snippet"}).
		Q(service.YoutubeApiConfig.Query).
		Type(service.YoutubeApiConfig.Type).
		PublishedAfter(service.YoutubeApiConfig.PublishedAfter).
		MaxResults(service.YoutubeApiConfig.MaxResults)
	response, err = call.Do()
	return response, err
}

func NewYoutubeDataApiService(service *youtube.Service, apiConfig config.YoutubeApiConfig) YoutubeDataApiService {
	return youtubeDataApiService{
		youtubeService:   service,
		YoutubeApiConfig: apiConfig,
	}
}
