package repository

import (
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/pkg/util"
	"time"
)

type FavoriteRepository struct {
}

func (c FavoriteRepository) SelectFavorite(favorite *model.Favorite) bool {
	var count int
	model.DB.Table("favorite").Where("video_id = ? and user_id = ?", favorite.VideoId, favorite.UserId).Count(&count)
	return count != 0
}

func (c FavoriteRepository) IsFavorite(vedioId int64, userId int64) bool {
	var count int
	model.DB.Table("favorite").Where("video_id = ? and user_id = ?", vedioId, userId).Count(&count)
	return count != 0
}

func (c FavoriteRepository) CreatFavorite(favorite *model.Favorite) error {
	err := model.DB.Debug().Create(favorite).Error
	return err
}

func (c FavoriteRepository) DeleteFavorite(favorite *model.Favorite) error {
	err := model.DB.Table("favorite").Where("video_id = ? and user_id = ?", favorite.VideoId, favorite.UserId).Delete(favorite).Error
	return err
}

func (c FavoriteRepository) FavoriteAct(favorite *model.Favorite, actionType int32) error {
	if actionType == 1 {
		snow := util.Snowflake{}
		favoriteId := snow.Generate()
		nowTime := time.Now()
		favorite.Id = favoriteId
		favorite.CreateTime = nowTime
		return FavoriteRepository{}.CreatFavorite(favorite)
	} else {
		return FavoriteRepository{}.DeleteFavorite(favorite)
	}
}

func (c FavoriteRepository) FavoriteList(userId int64, favorites *[]model.Favorite) {
	model.DB.Where("user_id=?", userId).Order("create_time DESC").Find(&favorites)
}
