package controller

import (
	"github.com/RaymondCode/simple-demo/pkg/e"
	"github.com/RaymondCode/simple-demo/serializer"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	var favoriteService service.FavoriteService

	//获取参数
	token := c.Query("token")
	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	actionType, _ := strconv.ParseInt(c.Query("action_type"), 10, 32)

	//参数检查
	if actionType != 1 && actionType != 2 {
		code := e.InvalidParams
		c.JSON(http.StatusOK, serializer.Response{
			StatusCode: code,
			StatusMsg:  e.GetMsg(code),
		})
		return
	}

	/**
	if _, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
	*/

	// 点赞操作
	res := favoriteService.CreateFavorite(userId, videoId, int32(actionType), token)
	c.JSON(http.StatusOK, res)
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	//var favoriteService service.FavoriteService
	//
	////获取参数
	//token := c.Query("token")
	//userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)

	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: DemoVideos,
	})
}
