package controller

import (
	"github.com/RaymondCode/simple-demo/pkg/e"
	"github.com/RaymondCode/simple-demo/serializer"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

//-------------------------------------------------------
//项目结构路径：/controller/publish.go
//创建者：贺凯恒
//审查者：杭朋洁
//创建时间：2022/5/25
//描述：视频功能相关的controller层
//Copyright2022
//--------------------------------------------------------

type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}

/*
	接口地址：/controller/publish.Publish
	功能描述：登录用户选择视频上传
	详细描述：检查 token 然后上传视频到 public 目录下
	参数：
		param token 登陆的用户
		param data 视频数据信息
		param title 视频标题
	请求方式：POST
	作者：贺凯恒
	创建时间：2022/5/25
	Copyright2022
*/
func Publish(c *gin.Context) {
	var publishService service.PublishService
	token := c.PostForm("token")
	data, err := c.FormFile("data")
	title := c.PostForm("title")
	if err != nil {
		code := e.ErrorUpLoadFile
		c.JSON(http.StatusOK, serializer.Response{
			StatusCode: code,
			StatusMsg:  e.GetMsg(code),
		})
		return
	}
	res := publishService.Publish(token, data, title, c)
	c.JSON(http.StatusOK, res)
}

/*
	接口地址：/controller/publish.PublishList
	功能描述：登录用户的视频发布列表，直接列出用户所有投稿过的视频
	详细描述：检查 token 然后上传视频到 public 目录下
	参数：
		param token 登陆的用户
		param user_id 用户id
	请求方式：GET
	作者：贺凯恒
	创建时间：2022/5/25
	Copyright2022
*/
func PublishList(c *gin.Context) {
	userId := c.Query("user_id")
	token := c.Query("token")
	var publishListService service.PublishService
	res := publishListService.PublishList(userId, token)
	c.JSON(http.StatusOK, res)
}
