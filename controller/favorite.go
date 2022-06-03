package controller

import (
	"github.com/RaymondCode/simple-demo/pkg/e"
	"github.com/RaymondCode/simple-demo/pkg/util"
	"github.com/RaymondCode/simple-demo/serializer"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

// FavoriteAction
func FavoriteAction(c *gin.Context) {
	var favoriteService service.FavoriteService

	//获取参数
	token := c.Query("token")
	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	actionType, _ := strconv.ParseInt(c.Query("action_type"), 10, 32)

	//身份判断
	claims, err := util.ParseToken(token)
	if err != nil {
		code := e.ErrorAuthCheckTokenTimeout
		c.JSON(http.StatusOK, serializer.Response{
			StatusCode: code,
			StatusMsg:  e.GetMsg(code),
		})
		return
	} else if time.Now().Unix() > claims.ExpiresAt {
		code := e.ErrorAuthCheckTokenTimeout
		c.JSON(http.StatusOK, serializer.Response{
			StatusCode: code,
			StatusMsg:  e.GetMsg(code),
		})
		return
	}
	userId = claims.Id

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
	res := favoriteService.DisposeFavorite(userId, videoId, int32(actionType), token)
	c.JSON(http.StatusOK, res)

}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	var favoriteService service.FavoriteService

	//获取参数
	token := c.Query("token")
	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)

	//身份判断
	claims, err := util.ParseToken(token)
	if err != nil {
		code := e.ErrorAuthCheckTokenTimeout
		c.JSON(http.StatusOK, serializer.Response{
			StatusCode: code,
			StatusMsg:  e.GetMsg(code),
		})
		return
	} else if time.Now().Unix() > claims.ExpiresAt {
		code := e.ErrorAuthCheckTokenTimeout
		c.JSON(http.StatusOK, serializer.Response{
			StatusCode: code,
			StatusMsg:  e.GetMsg(code),
		})
		return
	}
	userId = claims.Id

	// 获取点赞列表
	res := favoriteService.GetFavorites(userId)

	c.JSON(http.StatusOK, res)
}
