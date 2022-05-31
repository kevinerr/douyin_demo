package repository

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/model"
)

type FavoriteRepository struct {
}

func (c FavoriteRepository) CreatFavorite(favorite *model.Favorite) error {
	//err := model.DB.Select("UserId", "CreateTime").Create(favorite).Error
	fmt.Println(favorite)
	err := model.DB.Debug().Select("UserId", "VideoId").Create(favorite).Error
	return err
}

func (c FavoriteRepository) DeleteFavorite(favorite *model.Favorite) error {
	fmt.Println(favorite)
	err := model.DB.Debug().Where("video_id = ?", favorite.VideoId).Where("user_id = ?", favorite.UserId).Delete(favorite).Error
	return err
}

func (c FavoriteRepository) FavoriteAct(favorite *model.Favorite, actionType int32) error {
	if actionType == 1 {
		FavoriteRepository{}.DeleteFavorite(favorite)
		return FavoriteRepository{}.CreatFavorite(favorite)
	} else {
		return FavoriteRepository{}.DeleteFavorite(favorite)
	}
}
