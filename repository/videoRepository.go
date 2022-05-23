package repository

import (
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/serializer"
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
