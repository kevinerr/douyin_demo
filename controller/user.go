package controller

import (
	"fmt"
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

//
//var userIdSequence = int64(1)
//
////type UserLoginResponse struct {
////	Response
////	UserId int64  `json:"user_id,omitempty"`
////	Token  string `json:"token"`
////}
////
////type UserResponse struct {
////	Response
////	User User `json:"user"`
////}

func Register(c *gin.Context) {
	var userRegisterService service.UserService
	username := c.Query("username")
	password := c.Query("password")
	fmt.Println(username, len(username))
	fmt.Println(password, len(password))
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

	if len(password) < 6 || len(username) < 3 {
		code := e.InvalidParams
		c.JSON(code, e.GetMsg(code))
		return
	}
	userLoginService.UserName = username
	userLoginService.Password = password
	res := userLoginService.Login()
	fmt.Println(res)
	c.JSON(200, res)
}

func UserInfo(c *gin.Context) {
	token := c.Query("token")
	fmt.Println(token)

	//if user, exist := usersLoginInfo[token]; exist {
	//	c.JSON(http.StatusOK, UserResponse{
	//		Response: Response{StatusCode: 0},
	//		User:     user,
	//	})
	//} else {
	//	c.JSON(http.StatusOK, UserResponse{
	//		Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
	//	})
	//}
}
