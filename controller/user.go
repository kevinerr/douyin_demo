package controller

import (
	"github.com/RaymondCode/simple-demo/pkg/e"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

//-------------------------------------------------------
//项目结构路径：/controller/user.go
//创建者：贺凯恒
//审查者：杭朋洁
//创建时间：2022/5/25
//描述：用户登录、注册功能相关的controller层
//Copyright2022
//--------------------------------------------------------

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
var usersLoginInfo = map[string]User{
	"zhangleidouyin": {
		Id:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
}

/*
	接口地址：/controller/user.Register
	功能描述：用户注册接口
	参数：
		param username 用户注册使用的用户名
		param password 用户注册使用的密码
	请求方式：POST
	作者：贺凯恒
	创建时间：2022/5/25
	Copyright2022
*/
func Register(c *gin.Context) {
	var userRegisterService service.UserService
	username := c.Query("username")
	password := c.Query("password")
	//检验用户名和密码格式
	if len(password) < 6 || len(username) < 3 {
		code := e.InvalidParams
		c.JSON(code, e.GetMsg(code))
		return
	}
	userRegisterService.UserName = username
	userRegisterService.Password = password
	res := userRegisterService.Register()
	c.JSON(http.StatusOK, res)
}

/*
	接口地址：/controller/user.Login
	功能描述：用户登录接口
	参数：
		param username 用户登录使用的用户名
		param password 用户登录使用的密码
	请求方式：POST
	作者：贺凯恒
	创建时间：2022/5/25
	Copyright2022
*/
func Login(c *gin.Context) {
	var userLoginService service.UserService
	username := c.Query("username")
	password := c.Query("password")
	//检验用户名和密码格式
	if len(password) < 6 || len(username) < 3 {
		code := e.InvalidParams
		c.JSON(code, e.GetMsg(code))
		return
	}
	userLoginService.UserName = username
	userLoginService.Password = password
	res := userLoginService.Login()
	c.JSON(http.StatusOK, res)
}

/*
	接口地址：/controller/user.UserInfo
	功能描述：用户信息接口
	参数：
		param user_id 用户的id
		param token 登陆的用户
	请求方式：GET
	作者：贺凯恒
	创建时间：2022/5/25
	Copyright2022
*/
func UserInfo(c *gin.Context) {
	var userInfoService service.UserService
	userId := c.Query("user_id")
	token := c.Query("token")
	res := userInfoService.UserInfo(userId, token)
	c.JSON(http.StatusOK, res)
}
