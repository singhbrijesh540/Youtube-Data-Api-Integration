package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/singhbrijesh540/fampay-assignment/master"
	"github.com/singhbrijesh540/fampay-assignment/youtubeService"
	"net/http"
	"strconv"
)

type MasterController struct {
	BasePath           string
	YoutubeService     youtubeService.YoutubeDataApiService
	VideoDetailService master.VideoDetailService
}

type Controller interface {
	MakeRoutes(engine *gin.Engine)
}

func (c *MasterController) MakeRoutes(engine *gin.Engine) {
	getVideoDetailAPIRoutes(engine, c.BasePath, c.YoutubeService, c.VideoDetailService)
}

func getVideoDetailAPIRoutes(engine *gin.Engine, basePath string, youtubeService youtubeService.YoutubeDataApiService, videoDetailService master.VideoDetailService) {
	engine.GET(basePath+"/video-detail", GetVideoDetailHandler(videoDetailService))
	engine.GET(basePath+"/search/video-detail", SearchVideoDetailHandler(videoDetailService))

	//Optimized Search Api to fetch Video Detail
	engine.GET(basePath+"/search/v2/video-detail", SearchVideoDetailHandlerV2(videoDetailService))
}

//Fetch Video Details with Pagination and sorted in Reverse order of their published datetime
func GetVideoDetailHandler(service master.VideoDetailService) gin.HandlerFunc {
	return func(context *gin.Context) {
		pageStr, ok := context.GetQuery("page")
		pageInt, err := strconv.Atoi(pageStr)
		if !ok || err != nil {
			pageInt = 0
		}
		sizeStr, ok := context.GetQuery("size")
		sizeInt, err := strconv.Atoi(sizeStr)
		if !ok || err != nil {
			sizeInt = 10
		}

		videosDetail, apiErr := service.FetchVideoDetail(pageInt, sizeInt)
		if apiErr != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": apiErr.Error()})
			return
		}
		context.JSON(http.StatusOK, videosDetail)
	}
}

//Search video details on the basis of given title and description
func SearchVideoDetailHandler(service master.VideoDetailService) gin.HandlerFunc {
	return func(context *gin.Context) {
		title, ok := context.GetQuery("title")
		if !ok {
			title = ""
		}
		description, ok := context.GetQuery("description")
		if !ok {
			description = ""
		}
		pageStr, ok := context.GetQuery("page")
		pageInt, err := strconv.Atoi(pageStr)
		if !ok || err != nil {
			pageInt = 0
		}
		sizeStr, ok := context.GetQuery("size")
		sizeInt, err := strconv.Atoi(sizeStr)
		if !ok || err != nil {
			sizeInt = 10
		}
		response, apiErr := service.SearchVideoDetail(master.SearchVideoDetailRequest{
			Title:       title,
			Description: description,
			Page:        pageInt,
			Size:        sizeInt,
		})
		if apiErr != nil || response == nil {
			context.JSON(http.StatusNotFound, gin.H{"error": apiErr})
			return
		}
		context.JSON(http.StatusOK, response)
	}
}

//Optimized Search Api to fetch Video Detail
func SearchVideoDetailHandlerV2(service master.VideoDetailService) gin.HandlerFunc {
	return func(context *gin.Context) {
		query, ok := context.GetQuery("query")
		if !ok {
			query = ""
		}
		response, apiErr := service.SearchVideoDetailV2(query)
		if apiErr != nil || response == nil {
			context.JSON(http.StatusNotFound, gin.H{"error": apiErr})
			return
		}
		context.JSON(http.StatusOK, response)
	}
}
