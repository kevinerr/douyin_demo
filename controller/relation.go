package controller

import (
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

//-------------------------------------------------------
//项目结构路径：/controller/relation.go
//创建者：林叶润
//审查者：杭朋洁
//创建时间：2022/5/25
//描述：用户关系功能相关的controller层
//Copyright2022
//--------------------------------------------------------

type FollowListResponse struct {
	service.Response
	UserList []service.User `json:"user_list"`
}

/*
	接口地址：/controller/relation.RelationAction
	功能描述：登录用户对其他用户进行关注/取消关注操作
	参数：
		param token 登陆的用户
		param to_user_id 其他用户id
		param action_type 操作 1-关注，2-取消关注
	请求方式：POST
	作者：林叶润
	创建时间：2022/5/25
	Copyright2022
*/
func RelationAction(c *gin.Context) {
	token := c.Query("token")
	//userId := c.Query("user_id")
	toUserId := c.Query("to_user_id")
	actionType := c.Query("action_type")
	var followService service.FollowService
	var res service.Response
	res = followService.RelationAction(toUserId, actionType, token)
	//返回有结构体没有数组 BUG 字段没有首字母大写导致的
	c.JSON(http.StatusOK, res)
}

/*
	接口地址：/controller/relation.GetFollowListByUId
	功能描述：用户关注的所有用户列表
	参数：
		param token 登陆的用户
		param user_id 用户id
	请求方式：GET
	作者：林叶润
	创建时间：2022/5/25
	Copyright2022
*/
func GetFollowListByUId(c *gin.Context) {
	token := c.Query("token")
	userId := c.Query("user_id")
	var res FollowListResponse
	var followService service.FollowService
	res.UserList, res.Response = followService.GetFollowListByUId(token, userId)
	c.JSON(http.StatusOK, res)
}

/*
	接口地址：/controller/relation.GetFollowerListByUId
	功能描述：用户所有的粉丝列表
	参数：
		param token 登陆的用户
		param user_id 用户id
	请求方式：GET
	作者：林叶润
	创建时间：2022/5/25
	Copyright2022
*/
// GetFollowerListByUId FollowerList all users have same follower list
func GetFollowerListByUId(c *gin.Context) {
	token := c.Query("token")
	userId := c.Query("user_id")
	var res FollowListResponse
	var followService service.FollowService
	res.UserList, res.Response = followService.GetFollowerListByUId(token, userId)
	c.JSON(http.StatusOK, res)
}
