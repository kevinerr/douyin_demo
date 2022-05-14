package service

import (
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/pkg/e"
	"github.com/RaymondCode/simple-demo/pkg/util"
	"github.com/RaymondCode/simple-demo/repository"
	"github.com/RaymondCode/simple-demo/serializer"
	"time"
)

type FeedService struct {
}

func (service *FeedService) VideoList(latestTime string, token string) serializer.FeedResponse {
	//var videoRepository repository.VideoRepository
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
	var videos []serializer.Video
	var res []model.Video
	model.DB.Find(&res).Scan(videos)
	for i := 0; i < len(videos); i++ {
		user, _ := userRepository.SelectById(res[i].AuthorId) //TODO 好笨的方法
		userResp := serializer.User{Id: user.Id, Name: user.Username, FollowCount: user.FollowCount, FollowerCount: user.FollowerCount}
		videos[i].Author = userResp
	}
	return serializer.FeedResponse{
		Response:  serializer.Response{StatusCode: code, StatusMsg: e.GetMsg(code)},
		VideoList: videos,
	}
}
