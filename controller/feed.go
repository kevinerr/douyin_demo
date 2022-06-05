package controller

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

//-------------------------------------------------------
//项目结构路径：/controller/feed.go
//创建者：贺凯恒
//审查者：杭朋洁
//创建时间：2022/5/25
//描述：feed（视频流）功能相关的controller层
//Copyright2022
//--------------------------------------------------------

type FeedResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

/*
	接口地址：/controller/feed.Feed
	功能描述：视频流接口
	参数：
		param latest_time 可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
		param token 可选参数，登录用户设置
	请求方式：GET
	作者：贺凯恒
	创建时间：2022/5/25
	Copyright2022
*/
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
