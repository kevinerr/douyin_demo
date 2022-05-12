package repository

import "github.com/RaymondCode/simple-demo/model"

type PublishRepository struct {
}

func (c PublishRepository) CreateVideo(video *model.Video) error {
	err := model.DB.Create(video).Error
	return err
}
