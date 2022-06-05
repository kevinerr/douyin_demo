package service

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/pkg/e"
	"github.com/RaymondCode/simple-demo/pkg/util"
	"github.com/RaymondCode/simple-demo/repository"
	"github.com/RaymondCode/simple-demo/serializer"
	"strconv"
	"time"
)

//-------------------------------------------------------
//项目结构路径：/service/feed.go
//创建者：贺凯恒
//审查者：杭朋洁
//创建时间：2022/5/25
//描述：feed（视频流）功能相关的service层
//Copyright2022
//--------------------------------------------------------

type FeedService struct {
}

// 时间格式
var timeLayoutStr = "2006-01-02 15:04:05"

func (service *FeedService) VideoList(latestTime string, token string) serializer.FeedResponse {
	var userRepository repository.UserRepository
	var favoriteRepository repository.FavoriteRepository
	code := e.SUCCESS
	var isLogin bool
	claims, err := util.ParseToken(token) //token判断查询者是否登录（可不登录）
	if err != nil {                       //未登录
		isLogin = false
	} else if time.Now().Unix() > claims.ExpiresAt { //token过期
		code = e.ErrorAuthCheckTokenTimeout
		return serializer.FeedResponse{
			Response: serializer.Response{StatusCode: code, StatusMsg: e.GetMsg(code)},
		}
	} else { //已登录
		isLogin = true
	}
	var videos = make([]serializer.Video, 2) //单次最多返回的视频个数
	var res []model.Video
	int64latestTime, _ := strconv.ParseInt(latestTime, 10, 64)                                                   //将string时间戳转化为int64时间戳
	timeStr := time.Unix(int64latestTime, 0).Format(timeLayoutStr)                                               //将int64时间戳装换成是string时间
	model.DB.Model(&model.Video{}).Where("create_time<?", timeStr).Limit(2).Order("create_time DESC").Find(&res) //返回按投稿时间小于timeStr的视频

	for i := 0; i < len(res); i++ {
		user, _ := userRepository.SelectById(res[i].AuthorId)
		videos[i].Id = res[i].Id
		videos[i].CoverUrl = res[i].CoverUrl
		videos[i].PlayUrl = res[i].PlayUrl
		videos[i].FavoriteCount = res[i].FavoriteCount
		videos[i].CommentCount = res[i].CommentCount
		videos[i].Title = res[i].Title
		var isFollow bool
		if isLogin == false { //未登录
			videos[i].IsFavorite = false
			isFollow = false
		} else { //已登陆
			videos[i].IsFavorite = favoriteRepository.IsFavorite(res[i].Id, claims.Id)
			_, isFollow = userRepository.IsFollow(claims.Id, res[i].AuthorId)
		}
		userResp := serializer.User{Id: user.Id, Name: user.Username, FollowCount: user.FollowCount, FollowerCount: user.FollowerCount, IsFollow: isFollow}
		videos[i].Author = userResp
	}
	var next_time time.Time
	if len(res)-1 >= 0 {
		next_time = res[len(res)-1].CreateTime
	} else {
		return serializer.FeedResponse{
			Response: serializer.Response{StatusCode: e.UNDOSUCCESS, StatusMsg: "视频都要被您刷完了鸭！~"},
		}
	}

	fmt.Println("返回的视频集最早时间：", (next_time.Unix()-3600*8)*1000, time.Unix(next_time.Unix()-3600*8, 0).Format(timeLayoutStr))

	return serializer.FeedResponse{
		Response:  serializer.Response{StatusCode: code, StatusMsg: e.GetMsg(code)},
		VideoList: videos,
		NextTime:  next_time.Unix() * 1000,
	}
}
