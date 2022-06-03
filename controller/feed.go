package controller

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type FeedResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

func Feed(c *gin.Context) {
	//可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
	fmt.Println("请求feed的携带时间", c.Query("latest_time"))
	latestTime := c.DefaultQuery("latest_time", strconv.Itoa(int(time.Now().Unix())))
	latestTime = latestTime[0:10]
	token := c.Query("token")
	var feedService service.FeedService
	res := feedService.VideoList(latestTime, token)
	c.JSON(http.StatusOK, res)
}
