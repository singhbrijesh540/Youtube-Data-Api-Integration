package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/singhbrijesh540/fampay-assignment/config"
	"github.com/singhbrijesh540/fampay-assignment/database"
	"github.com/singhbrijesh540/fampay-assignment/master"
	"github.com/singhbrijesh540/fampay-assignment/routes"
	"github.com/singhbrijesh540/fampay-assignment/youtubeService"
	"google.golang.org/api/youtube/v3"
	"log"
	"net/http"
	"os"
	"time"

	"google.golang.org/api/googleapi/transport"
)

func main() {

	os.Setenv("ENV", "local")
	db, err := database.NewDB()
	if err != nil {
		db = nil
	} else {
		fmt.Print(db.ConnPool)
	}
	env := config.GetEnv()

	client := &http.Client{
		Transport: &transport.APIKey{Key: env.YoutubeApiConfig.DeveloperKey},
	}

	service, err := youtube.New(client)
	if err != nil {
		log.Fatalf("Error creating new YouTube client: %v", err)
	}

	vdRepo := master.NewVideoDetailRepo(db)
	youtubeService := youtubeService.NewYoutubeDataApiService(service, env.YoutubeApiConfig)
	videoDetailService := master.NewVideoDetailService(vdRepo, youtubeService)

	if err != nil {
		fmt.Print("Error while Fetching data from youtube APi")
	}

	router := gin.Default()
	basePath := env.GoConfig.BasePath
	masterController := routes.MasterController{
		BasePath:           basePath,
		YoutubeService:     youtubeService,
		VideoDetailService: videoDetailService,
	}
	masterController.MakeRoutes(router)

	serverPort := env.GoConfig.ServerPort

	go GetYoutubeDataAndSaveCron(videoDetailService)
	err = router.Run(":" + serverPort)

	if err != nil {
		return
	}
}

//Function to fetch data from youtube data api and save in database every 10 sec
func GetYoutubeDataAndSaveCron(videoDetailService master.VideoDetailService) {
	for {
		time.Sleep(time.Second * 10)
		err := videoDetailService.GetVideoDetailAndSave()
		if err != nil {
			fmt.Print("Error while fetching data from youtube api")
		}
	}
}
