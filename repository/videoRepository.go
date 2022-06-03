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

func (c VideoRepository) AddVideoComment(videoId int64, actionType int32) {
	var video model.Video
	video.Id = videoId
	var expression string
	if actionType == 1 {
		expression = "comment_count + ?"
	} else if actionType == 2 {
		expression = "comment_count - ?"
	}
	model.DB.Model(&video).Where("id = ?", videoId).UpdateColumn("comment_count", gorm.Expr(expression, 1))
}

func (c VideoRepository) GetFavoriteVideoList(userId int64, videos *[]serializer.Video) {
	var favoriteRepo FavoriteRepository
	var videoRepo VideoRepository
	// 查favorite list
	var favorites []model.Favorite
	favoriteRepo.FavoriteList(userId, &favorites)
	for i := range favorites {
		// 分别查每个video
		user := serializer.User{}
		video := serializer.Video{
			Author: user,
		}
		videoRepo.GetVideoById(favorites[i].VideoId, &video)
		*videos = append(*videos, video)
	}
}

func (c VideoRepository) GetVideoById(videoId int64, video *serializer.Video) {
	var userRepo UserRepository
	var temp model.Video
	model.DB.Where("id = ?", videoId).First(&temp)
	video.Id = temp.Id
	video.CommentCount = temp.CommentCount
	video.FavoriteCount = temp.FavoriteCount
	video.Title = temp.Title
	video.PlayUrl = temp.PlayUrl
	video.CoverUrl = temp.CoverUrl
	user, _ := userRepo.SelectById(temp.AuthorId)
	if user != nil {
		video.Author.Id = user.Id
		video.Author.FollowerCount = user.FollowCount
		video.Author.FollowCount = user.FollowCount
	}
}

func (c VideoRepository) CheckVideoAvailable(videoId int64) bool {
	var count int
	model.DB.Table("video").Where("id = ?", videoId).Count(&count)
	return count != 0
}
