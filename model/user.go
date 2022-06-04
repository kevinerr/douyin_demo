package model

import (
	"golang.org/x/crypto/bcrypt"
)

//User 用户模型
//type User struct {
//	gorm.Model
//	UserName       string `gorm:"unique"`
//	PasswordDigest string
//	FollowCount    int64 `gorm:"default:0"`
//	FollowerCount  int64 `gorm:"default:0"`
//	IsFollow       bool  `gorm:"default:false"`
//}
type User struct {
	Id            int64  `gorm:"column:id"`             //用户ID
	Username      string `gorm:"column:username"`       //登录账号
	Password      string `gorm:"column:password"`       //登录密码
	Nickname      string `gorm:"column:nickname"`       //昵称
	FollowCount   int64  `gorm:"column:follow_count"`   //关注总数
	FollowerCount int64  `gorm:"column:follower_count"` //粉丝总数
}

// TableName 会将 User 的表名重写为 `user`
func (User) TableName() string {
	return "user"
}

const (
	PassWordCost = 12 //密码加密难度
)

//SetPassword 设置密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

//CheckPassword 校验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
