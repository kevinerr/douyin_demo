package model

import "time"

type Comment struct {
	Id         int64     `gorm:"column:id"`          //评论ID
	VideoId    int64     `gorm:"column:video_id"`    //视频ID
	UserId     int64     `gorm:"column:user_id"`     //评论者ID
	Content    string    `gorm:"column:content"`     //评论内容
	CreateTime time.Time `gorm:"column:create_time"` //评论时间
}
