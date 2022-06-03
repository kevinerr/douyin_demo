package controller

import (
	"github.com/RaymondCode/simple-demo/pkg/e"
	"github.com/RaymondCode/simple-demo/serializer"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	var publishService service.PublishService
	token := c.Query("token")
	data, err := c.FormFile("data")
	title := c.Query("title")
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

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	userId := c.Query("user_id")
	token := c.Query("token")
	var publishListService service.PublishService
	res := publishListService.PublishList(userId, token)
	c.JSON(http.StatusOK, res)
}
