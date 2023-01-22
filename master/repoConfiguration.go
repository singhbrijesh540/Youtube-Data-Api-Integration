package master

type VideoDetailRepo interface {
	SaveVideoDetail(videoDetail []*VideoDetail) ([]*VideoDetail, error)
	GetVideosDetail(page int, size int) ([]*VideoDetail, error)
	SearchVideoDetail(title string, description string, page int, size int) ([]*VideoDetail, error)
}
