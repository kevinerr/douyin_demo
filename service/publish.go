package service

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/pkg/e"
	"github.com/RaymondCode/simple-demo/pkg/util"
	"github.com/RaymondCode/simple-demo/repository"
	"github.com/RaymondCode/simple-demo/serializer"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
	"io"
	"mime/multipart"
	"path/filepath"
	"strconv"
	"time"
)

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
		userId := claims.Id
		finalName := fmt.Sprintf("%d_%s", userId, filename)
		saveFile := filepath.Join("./public/", finalName) //将文件存储到public文件夹中
		if err := c.SaveUploadedFile(data, saveFile); err != nil {
			code = e.ErrorUpLoadFile
			return serializer.Response{
				StatusCode: code, StatusMsg: e.GetMsg(code),
			}
		}
	*/
	//雪花算法生成ID
	video := &model.Video{
		Id:         snow.Generate(),
		AuthorId:   claims.Id,
		Title:      title,
		PlayUrl:    fileurl,
		CoverUrl:   "https://api.kdcc.cn/img/rand.php",
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
func (service *PublishService) PublishList(userId string, token string) serializer.PublishResponse {
	//var userInfoRepository repository.UserRepository
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
	userIdInt64, _ := strconv.ParseInt(userId, 10, 64)
	//_, flag := userInfoRepository.IsFollow(claims.Id, userIdInt64) //查询A是否关注B
	//isFollow := flag
	var results []serializer.Video
	user, _ := publishRepository.SelectById(userIdInt64) //TODO is_favorite
	userResp := serializer.User{Id: userIdInt64, Name: user.Username, FollowCount: user.FollowCount, FollowerCount: user.FollowerCount}
	model.DB.Model(&model.Video{}).Select("id,cover_url,play_url,favorite_count, comment_count,title").Where("author_id=?", userId).Find(&results)
	fmt.Println(results)
	for i := 0; i < len(results); i++ {
		results[i].Author = userResp
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
