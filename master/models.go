package master

import (
	"gorm.io/gorm"
	"time"
)

type VideoDetail struct {
	gorm.Model
	Title        string
	Description  string
	PublishedAt  time.Time
	ThumbnailUrl string
}
