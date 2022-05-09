package service

import (
	"github.com/RaymondCode/simple-demo/pkg/e"
	"github.com/RaymondCode/simple-demo/pkg/util"
	"github.com/RaymondCode/simple-demo/repository"
	"github.com/RaymondCode/simple-demo/serializer"
	"github.com/jinzhu/gorm"
	logging "github.com/sirupsen/logrus"
	"sync/atomic"
)

var userIdSequence = int32(1)

//UserRegisterService 用户服务
type UserService struct {
	UserName string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

func (service *UserService) Register() serializer.Response {
	code := e.SUCCESS
	var user repository.User
	var count int
	repository.DB.Model(&repository.User{}).Where("username=?", service.UserName).First(&user).Count(&count)
	//表单验证
	if count == 1 {
		code = e.ErrorExistUser
		return serializer.Response{
			StatusCode: code,
			StatusMsg:  e.GetMsg(code),
		}
	}
	user.UserName = service.UserName
	//加密密码
	if err := user.SetPassword(service.Password); err != nil {
		logging.Info(err)
		code = e.ErrorFailEncryption
		return serializer.Response{
			StatusCode: code,
			StatusMsg:  e.GetMsg(code),
		}
	}
	//生成ID
	atomic.AddInt32(&userIdSequence, 1)
	user.ID = uint(userIdSequence)
	//生成token
	token, err := util.GenerateToken(user.ID, service.UserName, 0)
	if err != nil {
		logging.Info(err)
		code = e.ErrorAuthToken
		return serializer.Response{
			StatusCode: code,
			StatusMsg:  e.GetMsg(code),
		}
	}
	//创建用户
	if err := repository.DB.Create(&user).Error; err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			StatusCode: code,
			StatusMsg:  e.GetMsg(code),
		}
	}
	return serializer.Response{
		StatusCode: code,
		Data:       serializer.Register{UserId: int(user.ID), Token: token},
		StatusMsg:  e.GetMsg(code),
	}
}

//Login 用户登陆函数
func (service *UserService) Login() serializer.Response {
	var user repository.User
	code := e.SUCCESS
	if err := repository.DB.Where("user_name=?", service.UserName).First(&user).Error; err != nil {
		//如果查询不到，返回相应的错误
		if gorm.IsRecordNotFoundError(err) {
			logging.Info(err)
			code = e.ErrorNotExistUser
			return serializer.Response{
				StatusCode: code,
				StatusMsg:  e.GetMsg(code),
			}
		}
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			StatusCode: code,
			StatusMsg:  e.GetMsg(code),
		}
	}
	if user.CheckPassword(service.Password) == false {
		code = e.ErrorNotCompare
		return serializer.Response{
			StatusCode: code,
			StatusMsg:  e.GetMsg(code),
		}
	}
	token, err := util.GenerateToken(user.ID, service.UserName, 0)
	if err != nil {
		logging.Info(err)
		code = e.ErrorAuthToken
		return serializer.Response{
			StatusCode: code,
			StatusMsg:  e.GetMsg(code),
		}
	}
	return serializer.Response{
		StatusCode: code,
		Data:       serializer.Login{UserId: int(user.ID), Token: token},
		StatusMsg:  e.GetMsg(code),
	}
}
