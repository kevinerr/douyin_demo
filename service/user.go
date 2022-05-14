package service

import (
	"github.com/RaymondCode/simple-demo/pkg/e"
	"github.com/RaymondCode/simple-demo/pkg/util"
	"github.com/RaymondCode/simple-demo/repository"
	"github.com/RaymondCode/simple-demo/serializer"
	"github.com/jinzhu/gorm"
	logging "github.com/sirupsen/logrus"
	"strconv"
	"time"
)

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
	user.Username = service.UserName
	//加密密码
	if err := user.SetPassword(service.Password); err != nil {
		logging.Info(err)
		code = e.ErrorFailEncryption
		return serializer.UserLoginResponse{
			Response: serializer.Response{StatusCode: code, StatusMsg: e.GetMsg(code)},
		}
	}
	//雪花算法生成ID
	snow := util.Snowflake{}
	user.Id = snow.Generate()
	//生成token
	token, err := util.GenerateToken(user.Id, service.UserName, 0)
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
		UserId:   user.Id,
		Token:    token,
	}
}

//Login 用户登陆函数
func (service *UserService) Login() serializer.UserLoginResponse {
	code := e.SUCCESS
	var userLoginRepository repository.UserRepository
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
	token, err := util.GenerateToken(user.Id, service.UserName, 0)
	if err != nil {
		logging.Info(err)
		code = e.ErrorAuthToken
		return serializer.UserLoginResponse{
			Response: serializer.Response{StatusCode: code, StatusMsg: e.GetMsg(code)},
		}
	}
	return serializer.UserLoginResponse{
		Response: serializer.Response{StatusCode: code, StatusMsg: e.GetMsg(code)},
		UserId:   user.Id,
		Token:    token,
	}
}

//useId是被查询者的ID，token判断查询者是否登录
func (service *UserService) UserInfo(userId string, token string) serializer.UserResponse {
	userIdInt64, _ := strconv.ParseInt(userId, 10, 64)
	var userInfoRepository repository.UserRepository
	code := e.SUCCESS
	claims, err := util.ParseToken(token) //token判断查询者是否登录
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
	user, err := userInfoRepository.SelectById(userIdInt64) //查询是否存在useId的用户
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

	_, flag := userInfoRepository.IsFollow(claims.Id, userIdInt64) //查询A是否关注B
	isFollow := flag
	return serializer.UserResponse{
		Response: serializer.Response{StatusCode: code, StatusMsg: e.GetMsg(code)},
		User:     serializer.User{Id: user.Id, Name: user.Username, FollowCount: user.FollowCount, FollowerCount: user.FollowerCount, IsFollow: isFollow},
	}
}
