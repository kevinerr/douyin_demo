package controller

import (
	"github.com/RaymondCode/simple-demo/pkg/e"
	"github.com/RaymondCode/simple-demo/serializer"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

//-------------------------------------------------------
//项目结构路径：/controller/comment.go
//创建者：祁盼
//审查者：杭朋洁
//创建时间：2022/5/25
//描述：评论功能相关的controller层
//Copyright2022
//--------------------------------------------------------

type CommentListResponse struct {
	Response
	CommentList []Comment `json:"comment_list,omitempty"`
}

/*
	接口地址：/controller/comment.CommentAction
	功能描述：登录用户对视频进行评论
	参数：
		param token 登陆的用户
		param comment_text 用户填写的评论内容，当action_type=1的时候使用
		param video_id 视频的id
		param user_id 用户id
		param comment_id 要删除的评论id，在action_type=2的时候使用
		param action_type 1-发布评论，2-删除评论
	请求方式：POST
	作者：祁盼
	创建时间：2022/5/25
	Copyright2022
*/
func CommentAction(c *gin.Context) {
	var commentService service.CommentService

	//获取参数
	token := c.Query("token")
	commentText := c.DefaultQuery("comment_text", "")
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	commentId, _ := strconv.ParseInt(c.Query("comment_id"), 10, 64)
	actionType, _ := strconv.ParseInt(c.Query("action_type"), 10, 32)

	log.Println(userId) //无用参数

	//参数检查
	if actionType != 1 && actionType != 2 {
		code := e.InvalidParams
		c.JSON(http.StatusOK, serializer.Response{
			StatusCode: code,
			StatusMsg:  e.GetMsg(code),
		})
		return
	}

	//执行操作
	if actionType == 1 { //发布评论
		res := commentService.CreateAction(videoId, token, commentText)
		c.JSON(http.StatusOK, res)
	} else if actionType == 2 { //删除评论
		res := commentService.DeleteAction(commentId, token)
		c.JSON(http.StatusOK, res)
	}
}

/*
	接口地址：/controller/comment.CommentList
	功能描述：获取视频的评论信息，按发布时间倒叙排序
	参数：
		param token 登陆的用户
		param video_id 视频的id
	请求方式：GET
	作者：祁盼
	创建时间：2022/5/25
	Copyright2022
*/
func CommentList(c *gin.Context) {
	var commentService service.CommentService

	//获取参数
	token := c.Query("token")
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)

	//执行操作
	res := commentService.CommentList(videoId, token)
	c.JSON(http.StatusOK, res)
}
