package controller

import (
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
	latestTime := c.DefaultQuery("latest_time", strconv.Itoa(int(time.Now().Unix())))
	token := c.Query("token")
	//latestTime := c.Query("latest_time")
	//t := time.Now()
	//fmt.Println(latestTime)
	//fmt.Println(t.Unix()) //1531293019
	//fmt.Println(t.Unix().Format("2006-01-02 15:04:05")) //2018-7-15 15:23:00

	//获取当前时间戳
	//fmt.Println(t.Unix()) //1531293019
	var feedService service.FeedService
	res := feedService.VideoList(latestTime, token)
	c.JSON(http.StatusOK, res)
	//service.VideoList(latestTime)
	//c.JSON(http.StatusOK, FeedResponse{
	//	Response:  Response{StatusCode: 0},
	//	VideoList: DemoVideos,
	//	NextTime:  time.Now().Unix(),
	//})
}
