package controller

import (
	"github.com/RaymondCode/simple-demo/pkg/e"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
//// test data: username=zhanglei, password=douyin
var usersLoginInfo = map[string]User{
	"zhangleidouyin": {
		Id:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
}

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
	c.JSON(200, res)
}

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
	c.JSON(200, res)
}

func UserInfo(c *gin.Context) {
	var userInfoService service.UserService
	userId := c.Query("user_id")
	token := c.Query("token")
	res := userInfoService.UserInfo(userId, token)
	c.JSON(200, res)
}
