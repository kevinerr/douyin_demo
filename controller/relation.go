package controller

import (
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

//type UserListResponse struct {
//	Response
//	UserList []User `json:"user_list"`
//}
//
//// RelationAction no practical effect, just check if token is valid
//func RelationAction(c *gin.Context) {
//	token := c.Query("token")
//
//	if _, exist := usersLoginInfo[token]; exist {
//		c.JSON(http.StatusOK, Response{StatusCode: 0})
//	} else {
//		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
//	}
//}
//
//// FollowList all users have same follow list
//func FollowList(c *gin.Context) {
//	c.JSON(http.StatusOK, UserListResponse{
//		Response: Response{
//			StatusCode: 0,
//		},
//		UserList: []User{DemoUser},
//	})
//}
//
//// FollowerList all users have same follower list
//func FollowerList(c *gin.Context) {
//	c.JSON(http.StatusOK, UserListResponse{
//		Response: Response{
//			StatusCode: 0,
//		},
//		UserList: []User{DemoUser},
//	})
//}

type FollowListResponse struct {
	Response
	userList []service.User `json:"user_list"`
}

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	token := c.Query("token")
	userId := c.Query("user_id")
	toUserId := c.Query("to_user_id")
	actionType := c.Query("action_type")
	var followService service.FollowService
	var res Response
	res = followService.RelationAction(userId, toUserId, actionType, token)
	c.JSON(http.StatusOK, res)
}

// GetFollowListByUId FollowList all users have same follow list
func GetFollowListByUId(c *gin.Context) {
	token := c.Query("token")
	userId := c.Query("user_id")
	var res FollowListResponse
	var followService service.FollowService
	res.userList, res.Response = followService.GetFollowListByUId(token, userId)
	c.JSON(http.StatusOK, res)
}

// GetFollowerListByUId FollowerList all users have same follower list
func GetFollowerListByUId(c *gin.Context) {
	token := c.Query("token")
	userId := c.Query("user_id")
	var res FollowListResponse
	var followService service.FollowService
	res.userList, res.Response = followService.GetFollowerListByUId(token, userId)
	c.JSON(http.StatusOK, res)
}
