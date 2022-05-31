package repository

import (
	"github.com/RaymondCode/simple-demo/model"
)

type FavoriteRepository struct {
}

func (c FavoriteRepository) SelectFavorite(favorite *model.Favorite) bool {
	var count int
	model.DB.Debug().Where("video_id = ?", favorite.VideoId).Where("user_id = ?", favorite.UserId).Count(&count)
	return count != 0
}

func (c FavoriteRepository) CreatFavorite(favorite *model.Favorite) error {
	err := model.DB.Debug().Select("UserId", "VideoId").Create(favorite).Error
	return err
}

func (c FavoriteRepository) DeleteFavorite(favorite *model.Favorite) error {
	err := model.DB.Debug().Where("video_id = ?", favorite.VideoId).Where("user_id = ?", favorite.UserId).Delete(favorite).Error
	return err
}

func (c FavoriteRepository) FavoriteAct(favorite *model.Favorite, actionType int32) error {
	if actionType == 1 {
		return FavoriteRepository{}.CreatFavorite(favorite)
	} else {
		return FavoriteRepository{}.DeleteFavorite(favorite)
	}
}

func (c FavoriteRepository) FavoriteList(userId int64, favorites *[]model.Favorite) {
	model.DB.Where("user_id=?", userId).Order("create_time DESC").Find(&favorites)
}
