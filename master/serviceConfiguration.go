package master

type VideoDetailService interface {
	GetVideoDetailAndSave() error
	FetchVideoDetail(pageInt int, sizeInt int) ([]*VideoDetail, error)
	SearchVideoDetail(request SearchVideoDetailRequest) ([]*VideoDetail, error)
	SearchVideoDetailV2(query string) ([]*VideoDetail, error)
}
