package model

import "time"

type Follow struct {
	Id         int64     `gorm:"column:id"`          //关注记录ID
	FollowerId int64     `gorm:"column:follower_id"` //粉丝用户ID
	FollowId   int64     `gorm:"column:follow_id"`   //被关注用户ID
	CreateTime time.Time `gorm:"column:create_time"` //关注记录时间
}
