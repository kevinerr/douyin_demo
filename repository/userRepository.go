package repository

import (
	"github.com/RaymondCode/simple-demo/model"
)

type UserRepository struct {
}

//创建一个用户
func (c UserRepository) CreateUser(user *model.User) error {
	err := model.DB.Create(user).Error
	return err
}

//判断用户名是否存在，存在flag返回true，否则返回false
func (c UserRepository) IsExistUser(username string) (*model.User, bool) {
	var user model.User
	var count int
	model.DB.Where("username=?", username).First(&user).Count(&count)
	if count == 1 {
		return &user, true
	}
	return &user, false
}

//根据用户ID查询一个用户
func (c UserRepository) SelectById(userId int64) (*model.User, error) {
	var user model.User
	if err := model.DB.Where("id=?", userId).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

//func (c UserRepository) SelectById2(userId int64, user *model.User) {
//	var user model.User
//	if err := model.DB.Where("id=?", userId).First(user)
//}

//follower_id--粉丝用户ID,follow_id--被关注用户ID，判断用户A是否关注用户B
func (c UserRepository) IsFollow(followerId int64, followId int64) (*model.Follow, bool) {
	var follow model.Follow
	var count int
	model.DB.Where("follower_id=? AND follow_id=?", followerId, followId).First(&follow).Count(&count)
	if count == 1 {
		return &follow, true
	}
	return &follow, false
}

//videoId--视频ID,userId--用户ID，判断用户是否喜爱视频
func (c UserRepository) IsFavorite(videoId int64, userId int64) (*model.Favorite, bool) {
	var favorite model.Favorite
	var count int
	model.DB.Where("video_id=? AND user_id=?", videoId, userId).First(&favorite).Count(&count)
	if count == 1 {
		return &favorite, true
	}
	return &favorite, false
}
