package repository

import (
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/serializer"
	"github.com/jinzhu/gorm"
)

type VideoRepository struct {
}

func (c VideoRepository) CreateVideo(video *model.Video) error {
	err := model.DB.Create(video).Error
	return err
}

func (c VideoRepository) VideoList(videos *[]serializer.Video) {
	var res []model.Video
	model.DB.Find(&res).Scan(videos)
}

func (c VideoRepository) GetAuthorId(videoId int64) (authorId int64, error error) {
	var video model.Video
	err := model.DB.Where("id=?", videoId).First(&video).Error
	return video.AuthorId, err
}

func (c VideoRepository) AddVideoFavorite(videoId int64, actionType int32) {
	var video model.Video
	video.Id = videoId
	var expression string
	if actionType == 1 {
		expression = "favorite_count + ?"
	} else if actionType == 2 {
		expression = "favorite_count - ?"
	}
	model.DB.Model(&video).UpdateColumn("favorite_count", gorm.Expr(expression, 1))
}
