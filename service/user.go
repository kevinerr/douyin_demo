package service

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/pkg/e"
	"github.com/RaymondCode/simple-demo/pkg/util"
	"github.com/RaymondCode/simple-demo/repository"
	"github.com/RaymondCode/simple-demo/serializer"
	"github.com/jinzhu/gorm"
	logging "github.com/sirupsen/logrus"
	"sync/atomic"
	"time"
)

var userIdSequence = int32(1)

//UserRegisterService 用户服务
type UserService struct {
	UserName string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

func (service *UserService) Register() serializer.UserLoginResponse {
	code := e.SUCCESS
	var userRegisterRepository repository.UserRepository
	user, flag := userRegisterRepository.IsExistUser(service.UserName)
	//表单验证
	if flag {
		code = e.ErrorExistUser
		return serializer.UserLoginResponse{
			Response: serializer.Response{StatusCode: code, StatusMsg: e.GetMsg(code)},
		}
	}
	user.UserName = service.UserName
	//加密密码
	if err := user.SetPassword(service.Password); err != nil {
		logging.Info(err)
		code = e.ErrorFailEncryption
		return serializer.UserLoginResponse{
			Response: serializer.Response{StatusCode: code, StatusMsg: e.GetMsg(code)},
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
		return serializer.UserLoginResponse{
			Response: serializer.Response{StatusCode: code, StatusMsg: e.GetMsg(code)},
		}
	}
	//创建用户
	if err := userRegisterRepository.CreateUser(user); err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.UserLoginResponse{
			Response: serializer.Response{StatusCode: code, StatusMsg: e.GetMsg(code)},
		}
	}
	return serializer.UserLoginResponse{
		Response: serializer.Response{StatusCode: code, StatusMsg: e.GetMsg(code)},
		UserId:   int64(user.ID),
		Token:    token,
	}
}

//Login 用户登陆函数
func (service *UserService) Login() serializer.UserLoginResponse {
	//var user model.User
	code := e.SUCCESS
	var userLoginRepository repository.UserRepository
	//var flag bool
	user, flag := userLoginRepository.IsExistUser(service.UserName)
	if !flag {
		//如果查询不到，返回相应的错误
		code = e.ErrorNotExistUser
		return serializer.UserLoginResponse{
			Response: serializer.Response{StatusCode: code, StatusMsg: e.GetMsg(code)},
		}
	}
	if user.CheckPassword(service.Password) == false {
		code = e.ErrorNotCompare
		return serializer.UserLoginResponse{
			Response: serializer.Response{StatusCode: code, StatusMsg: e.GetMsg(code)},
		}
	}
	token, err := util.GenerateToken(user.ID, service.UserName, 0)
	if err != nil {
		logging.Info(err)
		code = e.ErrorAuthToken
		return serializer.UserLoginResponse{
			Response: serializer.Response{StatusCode: code, StatusMsg: e.GetMsg(code)},
		}
	}
	return serializer.UserLoginResponse{
		Response: serializer.Response{StatusCode: code, StatusMsg: e.GetMsg(code)},
		UserId:   int64(user.ID),
		Token:    token,
	}
}

func (service *UserService) UserInfo(userId string, token string) serializer.UserResponse {
	//var user model.User
	var userInfoRepository repository.UserRepository
	claims, err := util.ParseToken(token)
	code := e.SUCCESS
	if err != nil {
		code = e.ErrorAuthCheckTokenFail
		return serializer.UserResponse{
			Response: serializer.Response{StatusCode: code, StatusMsg: e.GetMsg(code)},
		}
	} else if time.Now().Unix() > claims.ExpiresAt {
		code = e.ErrorAuthCheckTokenTimeout
		return serializer.UserResponse{
			Response: serializer.Response{StatusCode: code, StatusMsg: e.GetMsg(code)},
		}
	}
	user, err := userInfoRepository.SelectById(userId)
	if err != nil {
		//如果查询不到，返回相应的错误
		if gorm.IsRecordNotFoundError(err) {
			logging.Info(err)
			code = e.ErrorNotExistUser
			return serializer.UserResponse{
				Response: serializer.Response{StatusCode: code, StatusMsg: e.GetMsg(code)},
			}
		}
	}
	fmt.Println(user.FollowerCount)
	return serializer.UserResponse{
		Response: serializer.Response{StatusCode: code, StatusMsg: e.GetMsg(code)},
		User:     serializer.User{Id: int64(user.ID), Name: user.UserName, FollowCount: user.FollowCount, FollowerCount: user.FollowerCount, IsFollow: user.IsFollow},
	}
}
