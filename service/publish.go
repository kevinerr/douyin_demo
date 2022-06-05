package service

import (
	"bytes"
	"fmt"
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/pkg/e"
	"github.com/RaymondCode/simple-demo/pkg/util"
	"github.com/RaymondCode/simple-demo/repository"
	"github.com/RaymondCode/simple-demo/serializer"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

//-------------------------------------------------------
//项目结构路径：/service/publish.go
//创建者：贺凯恒
//审查者：杭朋洁
//创建时间：2022/5/25
//描述：视频功能相关的service层
//Copyright2022
//--------------------------------------------------------

type PublishService struct {
}

func (service *PublishService) Publish(token string, data *multipart.FileHeader, title string, c *gin.Context) serializer.Response {
	var videoRepository repository.VideoRepository
	code := e.SUCCESS
	claims, err := util.ParseToken(token) //token判断查询者是否登录
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
	snow := util.Snowflake{}
	filename := strconv.Itoa(int(snow.Generate())) + "_" + filepath.Base(data.Filename)
	src, err := data.Open()
	if err != nil {
		code = e.InvalidParams
		return serializer.Response{
			StatusCode: code, StatusMsg: e.GetMsg(code),
		}
	}

	fileurl, err := OssUpload(filename, src)
	if err != nil {
		code = e.ErrorUpLoadFile
		return serializer.Response{
			StatusCode: code, StatusMsg: e.GetMsg(code),
		}
	}
	// 2022 - 06- 03 上传到oss

	/*
	* @Author: starine
	* @Date:   2022/6/5 16:19
	* @Description: 使用ffmpeg读取视频流中的第一帧作为封面。
	 */
	//先存入本地public文件夹
	saveFile := filepath.Join("./public/", filename) //将文件存储到public文件夹中
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		code = e.ErrorUpLoadFile
		return serializer.Response{
			StatusCode: code, StatusMsg: e.GetMsg(code),
		}
	}
	//使用ffmpeg从public文件中读取
	reader := bytes.NewBuffer(nil)
	err = ffmpeg.Input(saveFile).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", 1)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(reader, os.Stdout).
		Run()
	if err != nil {
		logging.Error(err)
	}
	//封面上传到oss
	covername := strconv.Itoa(int(snow.Generate())) + ".jpeg"
	coverurl, err := OssUpload(covername, reader)
	if err != nil {
		code = e.ErrorUpLoadFile
		return serializer.Response{
			StatusCode: code, StatusMsg: e.GetMsg(code),
		}
	}
	//删除本地public中的视频
	err = os.Remove(saveFile)
	if err != nil {
		logging.Info(err)
	}

	//雪花算法生成ID
	video := &model.Video{
		Id:         snow.Generate(),
		AuthorId:   claims.Id,
		Title:      title,
		PlayUrl:    fileurl,
		CoverUrl:   coverurl,
		CreateTime: time.Now(),
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
func (service *PublishService) PublishList(authorId string, token string) serializer.PublishResponse {
	//var userInfoRepository repository.UserRepository
	var publishRepository repository.UserRepository
	var favoriteRepository repository.FavoriteRepository
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

	authorIdInt64, _ := strconv.ParseInt(authorId, 10, 64)
	var results []serializer.Video
	user, _ := publishRepository.SelectById(authorIdInt64)
	_, isFollow := publishRepository.IsFollow(claims.Id, authorIdInt64)
	userResp := serializer.User{Id: authorIdInt64, Name: user.Username, FollowCount: user.FollowCount, FollowerCount: user.FollowerCount, IsFollow: isFollow}
	model.DB.Model(&model.Video{}).Select("id,cover_url,play_url,favorite_count, comment_count,title").Where("author_id = ?", authorId).Find(&results)
	fmt.Println("以上Error可以无视")
	for i := 0; i < len(results); i++ {
		results[i].Author = userResp
		results[i].IsFavorite = favoriteRepository.IsFavorite(results[i].Id, claims.Id)
	}
	return serializer.PublishResponse{
		Response:  serializer.Response{StatusCode: code, StatusMsg: e.GetMsg(code)},
		VideoList: results,
	}
}

func OssUpload(fileName string, file io.Reader) (string, error) {
	endpoint := "oss-cn-shanghai.aliyuncs.com"
	accessKey := "LTAI4GEi2cat7zLt37PSrixz"
	secretKey := "6iNRN9bdVJKC5gyRJruWnIHlWrApH2"
	client, err := oss.New(endpoint, accessKey, secretKey, oss.Timeout(10, 120))

	if err != nil {
		return "", err
	}
	// 获取存储空间
	bucket, err := client.Bucket("paper-boot")
	if err != nil {
		return "", err
	}
	// 上传文件。
	err = bucket.PutObject("test/"+fileName, file)
	if err != nil {
		return "", err
	}
	return "https://paper-boot.oss-cn-shanghai.aliyuncs.com/test/" + fileName, nil
}
