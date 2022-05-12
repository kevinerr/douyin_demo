package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type FeedResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

func Feed(c *gin.Context) {
	latestTime := c.DefaultQuery("latest_time", string(time.Now().Unix()))
	//latestTime := c.Query("latest_time")
	t := time.Now()
	fmt.Println(latestTime)
	fmt.Println(t.Unix()) //1531293019
	//fmt.Println(t.Unix().Format("2006-01-02 15:04:05")) //2018-7-15 15:23:00

	//获取当前时间戳
	fmt.Println(t.Unix()) //1531293019

	//service.VideoList(latestTime)
	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: DemoVideos,
		NextTime:  time.Now().Unix(),
	})
}
