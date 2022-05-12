package model

import (
	"time"
)

type Video struct {
	Id            int64     `gorm:"column:id"`             //视频ID
	AuthorId      int64     `gorm:"column:author_id"`      //作者的ID
	Description   string    `gorm:"column:description"`    //视频描述
	PlayUrl       string    `gorm:"column:play_url"`       //视频播放地址
	CoverUrl      string    `gorm:"column:cover_url"`      //视频封面地址
	FavoriteCount int64     `gorm:"column:favorite_count"` //视频的点赞总数
	CommentCount  int64     `gorm:"column:comment_count"`  //视频的评论总数
	CreateTime    time.Time `gorm:"column:create_time"`    //视频创建时间
}
