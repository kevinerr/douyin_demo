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

//-------------------------------------------------------
//项目结构路径：/controller/favorite.go
//创建者：谢庭宇
//审查者：杭朋洁
//创建时间：2022/5/25
//描述：点赞功能相关的controller层
//Copyright2022
//--------------------------------------------------------

/*
	接口地址：/controller/favorite.FavoriteAction
	功能描述：登录用户对视频进行点赞/取消点赞操作
	参数：
		param token 登陆的用户
		param user_id 用户的id
		param video_id 视频的id
		param action_type 操作（1-点赞 2-取消点赞）
	请求方式：POST
	作者：谢庭宇
	创建时间：2022/5/25
	Copyright2022
*/
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
	res := favoriteService.DisposeFavorite(userId, videoId, int32(actionType), claims)
	c.JSON(http.StatusOK, res)

}

/*
	接口地址：/controller/favorite.FavoriteList
	功能描述：获取用户的所有点赞视频
	参数：
		param token 登陆的用户
		param user_id 用户的id
	请求方式：GET
	作者：谢庭宇
	创建时间：2022/5/25
	Copyright2022
*/
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

	// 获取点赞列表
	res := favoriteService.GetFavorites(userId)

	c.JSON(http.StatusOK, res)
}
