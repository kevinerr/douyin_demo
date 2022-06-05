package repository

import (
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/pkg/util"
	"gorm.io/gorm"
	"time"
)

type FollowRepository struct {
}

// GetFollowListByUId 获取用户关注列表
func (c FollowRepository) GetFollowListByUId(userId int64) ([]model.User, error) {
	var users []model.User
	err := model.DB.Table("follow").
		Select("user.id, user.nickname, user.follow_count, user.follower_count").
		Joins("join user on follow.follow_id = user.id").
		Where("follower_id = ?", userId).
		Scan(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

// GetFollowerListByUId 获取用户的粉丝列表
func (c FollowRepository) GetFollowerListByUId(userId int64) ([]model.User, error) {
	var users []model.User
	err := model.DB.Table("follow").
		Select("user.id, user.nickname, user.follow_count, user.follower_count").
		Joins("join user on follow.follower_id = user.id").
		Where("follow_id = ?", userId).
		Scan(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

// RelationAct 关注与取关
func (c FollowRepository) RelationAct(userId, toUserId int64, actionType int32) bool {
	if actionType == 1 {
		return ok(userId, toUserId)
	} else if actionType == 2 {
		return cancel(userId, toUserId)
	}
	return false
}

func cancel(userId, toUserId int64) bool {
	//开启事务
	tx := model.DB.Begin()

	//删除关联
	err := tx.Where("follower_id = ? and follow_id = ?", userId, toUserId).
		Delete(model.Follow{}).Error
	if err != nil {
		//回滚
		tx.Rollback()
		return false
	}
	//用户关注数处理
	err = tx.Model(&model.User{Id: userId}).Update("follow_count", gorm.Expr("follow_count - 1")).Error
	err = tx.Model(&model.User{Id: toUserId}).Update("follower_count", gorm.Expr("follower_count - 1")).Error
	if err != nil {
		//回滚
		tx.Rollback()
		return false
	}
	//提交事务
	tx.Commit()
	return true
}

func ok(userId, toUserId int64) bool {
	//开启事务
	tx := model.DB.Begin()
	//创建关系
	//雪花算法生成主键
	snowflake := util.Snowflake{}
	var follow = model.Follow{
		Id:         snowflake.Generate(),
		FollowerId: userId,
		FollowId:   toUserId,
		CreateTime: time.Now(),
	}
	//创建关系
	err := tx.Create(&follow).Error
	if err != nil {
		//回滚
		tx.Rollback()
		return false
	}
	//更新用户信息
	err = tx.Model(&model.User{Id: userId}).Update("follow_count", gorm.Expr("follow_count + 1")).Error
	err = tx.Model(&model.User{Id: toUserId}).Update("follower_count", gorm.Expr("follower_count + 1")).Error
	if err != nil {
		//回滚
		tx.Rollback()
		return false
	}
	//提交事务
	tx.Commit()
	return true
}
