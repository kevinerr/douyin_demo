package service

import (
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/pkg/e"
	"github.com/RaymondCode/simple-demo/pkg/util"
	"github.com/RaymondCode/simple-demo/repository"
	"github.com/RaymondCode/simple-demo/serializer"
	logging "github.com/sirupsen/logrus"
	"time"
)

type FavoriteService struct {
}

func (service *FavoriteService) CreateFavorite(userId int64, videoId int64, actionType int32, token string) serializer.FavoriteActionResponse {
	var videoRepository repository.VideoRepository
	var favoriteRepository repository.FavoriteRepository
	code := e.SUCCESS

	//组合对象
	favorite := &model.Favorite{
		VideoId: videoId,
		UserId:  userId,
	}

	//身份判断
	claims, err := util.ParseToken(token)
	if err != nil {
		code = e.ErrorAuthCheckTokenFail
		return serializer.FavoriteActionResponse{
			Response: serializer.Response{StatusCode: code, StatusMsg: e.GetMsg(code)},
		}
	} else if time.Now().Unix() > claims.ExpiresAt {
		code = e.ErrorAuthCheckTokenTimeout
		return serializer.FavoriteActionResponse{
			Response: serializer.Response{StatusCode: code, StatusMsg: e.GetMsg(code)},
		}
	}
	userId = claims.Id

	// 判断是否已点赞或者已经取消点赞
	// 如果已经做了即不做处理
	if actionType == 1 {
		// 需要不存在
		if favoriteRepository.SelectFavorite(favorite) {
			code = e.SUCCESS
			return serializer.FavoriteActionResponse{
				Response: serializer.Response{StatusCode: code, StatusMsg: e.GetMsg(code)},
			}
		}
	} else {
		// 需要存在
		if !favoriteRepository.SelectFavorite(favorite) {
			code = e.SUCCESS
			return serializer.FavoriteActionResponse{
				Response: serializer.Response{StatusCode: code, StatusMsg: e.GetMsg(code)},
			}
		}
	}

	// TODO xietingyu redis + 定时任务实现
	// TODO xietingyu 判断视频ID是否正常
	// 视频点赞数++或--
	videoRepository.AddVideoFavorite(videoId, actionType)
	// 视频点赞表添加一条数据或者删除数据
	if err := favoriteRepository.FavoriteAct(favorite, actionType); err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
	} else {
		code = e.SUCCESS
	}
	return serializer.FavoriteActionResponse{
		Response: serializer.Response{StatusCode: code, StatusMsg: e.GetMsg(code)},
	}
}

func (service *FavoriteService) GetFavorites(userId int64, token string) interface{} {
	code := e.SUCCESS
	//身份判断
	//claims, err := util.ParseToken(token)
	//if err != nil {
	//	code = e.ErrorAuthCheckTokenFail
	//	return serializer.FavoriteActionResponse{
	//		Response: serializer.Response{StatusCode: code, StatusMsg: e.GetMsg(code)},
	//	}
	//} else if time.Now().Unix() > claims.ExpiresAt {
	//	code = e.ErrorAuthCheckTokenTimeout
	//	return serializer.FavoriteActionResponse{
	//		Response: serializer.Response{StatusCode: code, StatusMsg: e.GetMsg(code)},
	//	}
	//}
	//userId = claims.Id
	var videos []serializer.Video
	repository.VideoRepository{}.GetFavoriteVideoList(userId, &videos)
	return serializer.FavoriteListResponse{
		Response:  serializer.Response{StatusCode: code, StatusMsg: e.GetMsg(code)},
		VideoList: videos,
	}
}
