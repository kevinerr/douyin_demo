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
//项目结构路径：/service/comment.go
//创建者：祁盼
//审查者：杭朋洁
//创建时间：2022/5/25
//描述：评论功能相关的service层
//Copyright2022
//--------------------------------------------------------

type CommentService struct {
}

// 时间格式
var commentTimeLayoutStr = "01-02"

func (service *CommentService) CreateAction(videoId int64, token string, commentText string) serializer.CommentActionResponse {

	var commentRepository repository.CommentRepository
	var userRepository repository.UserRepository
	var videoRepository repository.VideoRepository
	code := e.SUCCESS

	//身份判断
	claims, err := util.ParseToken(token)
	if err != nil {
		code = e.ErrorAuthCheckTokenFail
		return serializer.CommentActionResponse{
			Response: serializer.Response{StatusCode: code, StatusMsg: e.GetMsg(code)},
		}
	} else if time.Now().Unix() > claims.ExpiresAt {
		code = e.ErrorAuthCheckTokenTimeout
		return serializer.CommentActionResponse{
			Response: serializer.Response{StatusCode: code, StatusMsg: e.GetMsg(code)},
		}
	}
	userId := claims.Id

	//组合comment对象
	nowTime := time.Now()
	snow := util.Snowflake{}
	newCommentId := snow.Generate()
	comment := &model.Comment{
		Id:         newCommentId,
		VideoId:    videoId,
		UserId:     userId,
		Content:    commentText,
		CreateTime: nowTime,
	}

	//插库操作
	if err := commentRepository.CreatComment(comment); err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.CommentActionResponse{
			Response: serializer.Response{StatusCode: code, StatusMsg: e.GetMsg(code)},
		}
	}

	//修改vedio信息
	videoRepository.AddVideoComment(videoId, 1)

	//ifFollow判断
	user, _ := userRepository.SelectById(userId)        //评论者
	authorId, _ := videoRepository.GetAuthorId(videoId) //作者
	var _, isFollow = userRepository.IsFollow(user.Id, authorId)

	//组合回调信息
	userContent := serializer.User{
		Id:            user.Id,
		Name:          user.Username,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
		IsFollow:      isFollow,
	}
	commentCallback := serializer.Comment{
		Id:         newCommentId,
		User:       userContent,
		Content:    commentText,
		CreateDate: nowTime.Format(commentTimeLayoutStr),
	}

	//返回
	return serializer.CommentActionResponse{
		Response: serializer.Response{StatusCode: code, StatusMsg: e.GetMsg(code)},
		Comment:  commentCallback,
	}

}

func (service *CommentService) DeleteAction(commentId int64, token string) serializer.CommentActionResponse {
	var commentRepository repository.CommentRepository
	var videoRepository repository.VideoRepository
	code := e.SUCCESS

	//身份判断
	claims, err := util.ParseToken(token)
	if err != nil {
		code = e.ErrorAuthCheckTokenFail
		return serializer.CommentActionResponse{
			Response: serializer.Response{StatusCode: code, StatusMsg: e.GetMsg(code)},
		}
	} else if time.Now().Unix() > claims.ExpiresAt {
		code = e.ErrorAuthCheckTokenTimeout
		return serializer.CommentActionResponse{
			Response: serializer.Response{StatusCode: code, StatusMsg: e.GetMsg(code)},
		}
	}
	userId := claims.Id

	//权限判断
	if commentAuthorId, err := commentRepository.GetAuthorId(commentId); err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.CommentActionResponse{
			Response: serializer.Response{StatusCode: code, StatusMsg: e.GetMsg(code)},
		}
	} else {
		if commentAuthorId != userId {
			code = e.OutOfUserPermission
			return serializer.CommentActionResponse{
				Response: serializer.Response{StatusCode: code, StatusMsg: e.GetMsg(code)},
			}
		}
	}

	videoId, _ := commentRepository.GetVideoId(commentId)

	//删除操作
	if err := commentRepository.DeleteComment(commentId); err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.CommentActionResponse{
			Response: serializer.Response{StatusCode: code, StatusMsg: e.GetMsg(code)},
		}
	}

	//修改vedio信息
	videoRepository.AddVideoComment(videoId, 2)

	//返回
	return serializer.CommentActionResponse{
		Response: serializer.Response{StatusCode: code, StatusMsg: e.GetMsg(code)},
	}
}

func (service *CommentService) CommentList(videoId int64, token string) serializer.CommentListResponse {
	var commentRepository repository.CommentRepository
	var userRepository repository.UserRepository
	var videoRepository repository.VideoRepository
	code := e.SUCCESS

	//判断身份
	claims, err := util.ParseToken(token)
	if err != nil {
		code = e.ErrorAuthCheckTokenFail
		return serializer.CommentListResponse{
			Response: serializer.Response{StatusCode: code, StatusMsg: e.GetMsg(code)},
		}
	} else if time.Now().Unix() > claims.ExpiresAt {
		code = e.ErrorAuthCheckTokenTimeout
		return serializer.CommentListResponse{
			Response: serializer.Response{StatusCode: code, StatusMsg: e.GetMsg(code)},
		}
	}

	//查询操作
	var comments []model.Comment
	if err := commentRepository.SelectComments(videoId, &comments); err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.CommentListResponse{
			Response: serializer.Response{StatusCode: code, StatusMsg: e.GetMsg(code)},
		}
	}

	//格式化返回内容
	var commentsInfoList []serializer.Comment
	for _, comment := range comments {
		//ifFollow判断
		user, _ := userRepository.SelectById(comment.UserId) //评论者
		authorId, _ := videoRepository.GetAuthorId(videoId)  //作者
		var _, isFollow = userRepository.IsFollow(user.Id, authorId)

		userInfo := serializer.User{
			Id:            user.Id,
			Name:          user.Username,
			FollowCount:   user.FollowCount,
			FollowerCount: user.FollowerCount,
			IsFollow:      isFollow,
		}
		commentInfo := serializer.Comment{
			Id:         comment.Id,
			User:       userInfo,
			Content:    comment.Content,
			CreateDate: comment.CreateTime.Format(commentTimeLayoutStr),
		}

		//拼接
		commentsInfoList = append(commentsInfoList, commentInfo)
	}

	//返回
	return serializer.CommentListResponse{
		Response:    serializer.Response{StatusCode: code, StatusMsg: e.GetMsg(code)},
		CommentList: commentsInfoList,
	}

}
