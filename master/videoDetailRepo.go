package master

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type videoDetailDao struct {
	db *gorm.DB
}

func (vDao videoDetailDao) SearchVideoDetail(title string, description string, page int, size int) ([]*VideoDetail, error) {
	offset := page * size
	var videoDetail []*VideoDetail
	var tx *gorm.DB
	if len(title) > 0 && len(description) > 0 {
		tx = vDao.db.
			Where("title like '%" + title + "%' and description like '%" + description + "%'").
			Offset(offset).
			Limit(size).
			Find(&videoDetail)
	} else if len(title) > 0 {
		tx = vDao.db.
			Where("title like '%" + title + "%'").
			Offset(offset).
			Limit(size).
			Find(&videoDetail)
	} else if len(description) > 0 {
		tx = vDao.db.
			Where("title like '%" + title + "%'").
			Offset(offset).
			Limit(size).
			Find(&videoDetail)
	}

	if tx == nil {
		return nil, nil
	}
	if tx.Error != nil {
		return nil, tx.Error
	}
	return videoDetail, nil
}

func (vDao videoDetailDao) GetVideosDetail(page int, size int) ([]*VideoDetail, error) {
	offset := page * size
	var videoDetail []*VideoDetail
	tx := vDao.db.Offset(offset).Limit(size).Order("published_at desc").Find(&videoDetail)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return videoDetail, nil
}

func (vDao videoDetailDao) SaveVideoDetail(videoDetail []*VideoDetail) ([]*VideoDetail, error) {
	tx := vDao.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "thumbnail_url"}},
		DoNothing: true,
	}).Create(&videoDetail)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return videoDetail, nil
}

func NewVideoDetailRepo(db *gorm.DB) VideoDetailRepo {
	return videoDetailDao{
		db: db,
	}
}
