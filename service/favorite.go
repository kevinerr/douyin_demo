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

//-------------------------------------------------------
//项目结构路径：/service/favorite.go
//创建者：谢庭宇
//审查者：杭朋洁
//创建时间：2022/5/25
//描述：点赞功能相关的service层
//Copyright2022
//--------------------------------------------------------

type FavoriteService struct {
}

func (service *FavoriteService) DisposeFavorite(userId int64, videoId int64, actionType int32, claims *util.Claims) serializer.FavoriteActionResponse {

	var videoRepository repository.VideoRepository
	var favoriteRepository repository.FavoriteRepository
	code := e.SUCCESS

	//组合对象
	favorite := &model.Favorite{
		VideoId: videoId,
		UserId:  userId,
	}

	//获取身份
	userId = claims.Id

	// 判断是否已点赞或者已经取消点赞
	// 如果已经做了即不做处理
	if actionType == 1 {
		// 点赞操作
		if favoriteRepository.SelectFavorite(favorite) {
			code = e.UNDOSUCCESS
			return serializer.FavoriteActionResponse{
				Response: serializer.Response{StatusCode: code, StatusMsg: e.GetMsg(code)},
			}
		}
	} else {
		// 删除点赞操作
		if !favoriteRepository.SelectFavorite(favorite) {
			code = e.UNDOSUCCESS
			return serializer.FavoriteActionResponse{
				Response: serializer.Response{StatusCode: code, StatusMsg: e.GetMsg(code)},
			}
		}
	}
	//雪花算法生成ID
	snow := util.Snowflake{}
	favorite.Id = snow.Generate()
	favorite.CreateTime = time.Now()

	if !videoRepository.CheckVideoAvailable(videoId) {
		code = e.InvalidParams
		return serializer.FavoriteActionResponse{
			Response: serializer.Response{StatusCode: code, StatusMsg: e.GetMsg(code)},
		}
	}

	// 视频点赞表添加一条数据或者删除数据
	if err := favoriteRepository.FavoriteAct(favorite, actionType); err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
	} else {
		code = e.SUCCESS
	}

	// 视频点赞数++或--
	videoRepository.AddVideoFavorite(videoId, actionType)

	return serializer.FavoriteActionResponse{
		Response: serializer.Response{StatusCode: code, StatusMsg: e.GetMsg(code)},
	}
}

func (service *FavoriteService) GetFavorites(userId int64) interface{} {
	code := e.SUCCESS

	var videos []serializer.Video
	repository.VideoRepository{}.GetFavoriteVideoList(userId, &videos)
	return serializer.FavoriteListResponse{
		Response:  serializer.Response{StatusCode: code, StatusMsg: e.GetMsg(code)},
		VideoList: videos,
	}
}
