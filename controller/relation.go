package controller

import (
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type FollowListResponse struct {
	service.Response
	UserList []service.User `json:"user_list"`
}

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	token := c.Query("token")
	userId := c.Query("user_id")
	toUserId := c.Query("to_user_id")
	actionType := c.Query("action_type")
	var followService service.FollowService
	var res service.Response
	res = followService.RelationAction(userId, toUserId, actionType, token)
	//返回有结构体没有数组 BUG 字段没有首字母大写导致的
	c.JSON(http.StatusOK, res)
}

// GetFollowListByUId FollowList all users have same follow list
func GetFollowListByUId(c *gin.Context) {
	token := c.Query("token")
	userId := c.Query("user_id")
	var res FollowListResponse
	var followService service.FollowService
	res.UserList, res.Response = followService.GetFollowListByUId(token, userId)
	c.JSON(http.StatusOK, res)
}

// GetFollowerListByUId FollowerList all users have same follower list
func GetFollowerListByUId(c *gin.Context) {
	token := c.Query("token")
	userId := c.Query("user_id")
	var res FollowListResponse
	var followService service.FollowService
	res.UserList, res.Response = followService.GetFollowerListByUId(token, userId)
	c.JSON(http.StatusOK, res)
}
