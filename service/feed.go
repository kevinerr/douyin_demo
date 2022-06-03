package service

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/conf"
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/pkg/e"
	"github.com/RaymondCode/simple-demo/pkg/util"
	"github.com/RaymondCode/simple-demo/repository"
	"github.com/RaymondCode/simple-demo/serializer"
	"strconv"
	"time"
)

type FeedService struct {
}

var timeLayoutStr = "2006-01-02 15:04:05"

func (service *FeedService) VideoList(latestTime0 string, token string) serializer.FeedResponse {
	latestTime := latestTime0[0:10]
	fmt.Println(latestTime)
	var userRepository repository.UserRepository
	code := e.SUCCESS
	claims, err := util.ParseToken(token) //token判断查询者是否登录
	if err != nil {
		code = e.ErrorAuthCheckTokenFail
		return serializer.FeedResponse{
			Response: serializer.Response{StatusCode: code, StatusMsg: e.GetMsg(code)},
		}
	} else if time.Now().Unix() > claims.ExpiresAt {
		code = e.ErrorAuthCheckTokenTimeout
		return serializer.FeedResponse{
			Response: serializer.Response{StatusCode: code, StatusMsg: e.GetMsg(code)},
		}
	}
	var videos = make([]serializer.Video, conf.MaxVideosNum) //TODO,单次最多返回的视频个数
	var res []model.Video
	int64latestTime, err := strconv.ParseInt(latestTime, 10, 64)                                                                 //将string时间戳转化为int64时间戳
	timeStr := time.Unix(int64latestTime, 0).Format(timeLayoutStr)                                                               //将int64时间戳装换成是string时间
	model.DB.Model(&model.Video{}).Where("create_time<?", timeStr).Limit(conf.MaxVideosNum).Order("create_time DESC").Find(&res) //返回按投稿时间小于timeStr的视频
	fmt.Println(res)
	for i := 0; i < len(res); i++ {
		user, _ := userRepository.SelectById(res[i].AuthorId) //TODO 好笨的方法
		videos[i].Id = res[i].Id
		videos[i].CoverUrl = res[i].CoverUrl
		videos[i].PlayUrl = res[i].PlayUrl
		videos[i].FavoriteCount = res[i].FavoriteCount
		videos[i].CommentCount = res[i].CommentCount
		_, flag := userRepository.IsFavorite(res[i].Id, claims.Id) //判断用户是否喜爱此视频
		videos[i].IsFavorite = flag
		videos[i].Title = res[i].Title
		_, isFollow := userRepository.IsFollow(claims.Id, user.Id) //判断用户是否关注此视频的作者
		userResp := serializer.User{Id: user.Id, Name: user.Username, FollowCount: user.FollowCount, FollowerCount: user.FollowerCount, IsFollow: isFollow}
		videos[i].Author = userResp
	}
	next_time := res[len(res)-1].CreateTime
	//fmt.Println(next_time)
	//fmt.Println(next_time.Unix() - 3600*8)
	return serializer.FeedResponse{
		Response:  serializer.Response{StatusCode: code, StatusMsg: e.GetMsg(code)},
		VideoList: videos,
		NextTime:  next_time.Unix() * 1000,
	}
}
