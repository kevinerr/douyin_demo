package service

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/pkg/e"
	"github.com/RaymondCode/simple-demo/pkg/util"
	"github.com/RaymondCode/simple-demo/repository"
	"github.com/RaymondCode/simple-demo/serializer"
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
	"mime/multipart"
	"path/filepath"
	"time"
)

type PublishService struct {
}

func (service *PublishService) Publish(token string, data *multipart.FileHeader, c *gin.Context) serializer.Response {
	var videoRepository repository.VideoRepository
	claims, err := util.ParseToken(token) //token判断查询者是否登录
	code := e.SuccessUpLoadFile
	if err != nil {
		code = e.ErrorAuthCheckTokenFail
		return serializer.Response{
			StatusCode: code, StatusMsg: e.GetMsg(code),
		}
	} else if time.Now().Unix() > claims.ExpiresAt {
		code = e.ErrorAuthCheckTokenTimeout
		return serializer.Response{
			StatusCode: code, StatusMsg: e.GetMsg(code),
		}
	}
	filename := filepath.Base(data.Filename)
	userId := claims.Id
	finalName := fmt.Sprintf("%d_%s", userId, filename)
	saveFile := filepath.Join("./public/", finalName) //将文件存储到public文件夹中
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		code = e.ErrorUpLoadFile
		return serializer.Response{
			StatusCode: code, StatusMsg: e.GetMsg(code),
		}
	}
	//雪花算法生成ID
	snow := util.Snowflake{}
	video := &model.Video{
		Id:          snow.Generate(),
		AuthorId:    claims.Id,
		Description: "testDescription",
		PlayUrl:     "http://localhost:8080/static/" + filename,
		CoverUrl:    "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		CreateTime:  time.Now(),
	}
	//创建video
	if err := videoRepository.CreateVideo(video); err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			StatusCode: code,
			StatusMsg:  e.GetMsg(code),
		}
	}
	return serializer.Response{
		StatusCode: code,
		StatusMsg:  e.GetMsg(code),
	}
}
func (service *PublishService) PublishList(token string) serializer.PublishResponse {
	var publishRepository repository.UserRepository
	code := e.SUCCESS
	claims, err := util.ParseToken(token)
	if err != nil {
		code = e.ErrorAuthCheckTokenFail
		return serializer.PublishResponse{
			Response: serializer.Response{StatusCode: code, StatusMsg: e.GetMsg(code)},
		}
	} else if time.Now().Unix() > claims.ExpiresAt {
		code = e.ErrorAuthCheckTokenTimeout
		return serializer.PublishResponse{
			Response: serializer.Response{StatusCode: code, StatusMsg: e.GetMsg(code)},
		}
	}
	userId := claims.Id
	var results []serializer.Video
	user, _ := publishRepository.SelectById(userId) //TODO 好笨的方法
	userResp := serializer.User{Id: userId, Name: user.Username, FollowCount: user.FollowCount, FollowerCount: user.FollowerCount}
	model.DB.Table("video").Select("video.id,cover_url,play_url,favorite_count, comment_count").Joins("left join user on video.author_id = ?", userId).Scan(&results)
	fmt.Println(results)
	for i := 0; i < len(results); i++ {
		results[i].Author = userResp
	}
	return serializer.PublishResponse{
		Response:  serializer.Response{StatusCode: code, StatusMsg: e.GetMsg(code)},
		VideoList: results,
	}
}
