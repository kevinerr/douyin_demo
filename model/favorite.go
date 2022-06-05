package model

import "time"

type Favorite struct {
	Id         int64     `gorm:"column:id"`          //点赞记录ID
	UserId     int64     `gorm:"column:user_id"`     //用户的ID
	VideoId    int64     `gorm:"column:video_id"`    //点赞视频的ID
	CreateTime time.Time `gorm:"column:create_time"` //点赞时间
}

func (Favorite) TableName() string {
	return "favorite"
}
