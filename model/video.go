package model

import "github.com/jinzhu/gorm"

type Video struct {
	gorm.Model
	Author        User `gorm:"ForeignKey:User;AssociationForeignKey:ID"`
	Uid           uint `gorm:"not null"`
	PlayUrl       string
	CoverUrl      string
	FavoriteCount int64 `gorm:"default:1"`
	CommentCount  int64 `gorm:"default:1"`
	IsFavorite    bool  `gorm:"default:true"`
}
