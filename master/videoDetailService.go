package master

import (
	"fmt"
	"github.com/singhbrijesh540/fampay-assignment/youtubeService"
	"time"
)

type videoDetailService struct {
	videoDetailRepo VideoDetailRepo
	youtubeService  youtubeService.YoutubeDataApiService
}

func (vdService videoDetailService) SearchVideoDetail(request SearchVideoDetailRequest) ([]*VideoDetail, error) {
	videoDetail, err := vdService.videoDetailRepo.SearchVideoDetail(request.Title, request.Description, request.Page, request.Size)
	if err != nil {
		fmt.Print("[FetchVideoDetail] Error in fetching videos detail data")
		return nil, err
	}
	if videoDetail == nil {
		return nil, nil
	}
	return videoDetail, err
}

func (vdService videoDetailService) FetchVideoDetail(pageInt int, sizeInt int) ([]*VideoDetail, error) {
	videosdetail, err := vdService.videoDetailRepo.GetVideosDetail(pageInt, sizeInt)
	if err != nil {
		fmt.Print("[FetchVideoDetail] Error in fetching videos detail data")
	}
	return videosdetail, err
}

//Get data from youtube data v3 api and save it to postgres database after precessing the data
func (vdService videoDetailService) GetVideoDetailAndSave() error {
	response, err := vdService.youtubeService.GetDataFromYoutubeDataApi()
	if err != nil {
		return err
	}
	var videoDetailList []*VideoDetail
	for _, item := range response.Items {
		publishedAt, err := time.Parse(time.RFC3339, item.Snippet.PublishedAt)
		if err != nil {
			fmt.Print("error in parsing Published time")
			return err
		}
		videoDetailList = append(videoDetailList, &VideoDetail{
			Title:        item.Snippet.Title,
			Description:  item.Snippet.Description,
			PublishedAt:  publishedAt,
			ThumbnailUrl: item.Snippet.Thumbnails.High.Url,
		})
	}
	if len(videoDetailList) > 0 {
		videoDetail, err := vdService.videoDetailRepo.SaveVideoDetail(videoDetailList)
		if err != nil {
			fmt.Print("Error in creating video detail")
		}
		fmt.Print(len(videoDetail))
	}
	return err
}

func NewVideoDetailService(vdRepo VideoDetailRepo, yService youtubeService.YoutubeDataApiService) VideoDetailService {
	return videoDetailService{
		videoDetailRepo: vdRepo,
		youtubeService:  yService,
	}
}
