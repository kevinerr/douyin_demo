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

type CommentListResponse struct {
	Response
	CommentList []Comment `json:"comment_list,omitempty"`
}

// CommentAction
func CommentAction(c *gin.Context) {
	var commentService service.CommentService

	//获取参数
	token := c.Query("token")
	commentText := c.DefaultQuery("comment_text", "")
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	commentId, _ := strconv.ParseInt(c.Query("comment_id"), 10, 64)
	actionType, _ := strconv.ParseInt(c.Query("action_type"), 10, 32)

	log.Println(userId) //这个参数就没啥用

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

// CommentList
func CommentList(c *gin.Context) {
	var commentService service.CommentService

	//获取参数
	token := c.Query("token")
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)

	//执行操作
	res := commentService.CommentList(videoId, token)
	c.JSON(http.StatusOK, res)
}
