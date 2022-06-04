package service

import (
	"github.com/RaymondCode/simple-demo/pkg/e"
	"github.com/RaymondCode/simple-demo/pkg/util"
	"github.com/RaymondCode/simple-demo/repository"
	"strconv"
	"time"
)

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}
type FollowService struct {
}

type User struct {
	Id            int64  `json:"id"`             //用户ID
	Nickname      string `json:"name"`           //昵称
	FollowCount   int64  `json:"follow_count"`   //关注总数
	FollowerCount int64  `json:"follower_count"` //粉丝总数
	IsFollow      bool   `json:"is_follow"`      // true为关注了，false为为关注,意思是是否彼此关注了
}

// RelationAction 关系操作
func (service *FollowService) RelationAction(toUserIdStr, actionTypeStr, token string) Response {
	//token验证
	claims, err := util.ParseToken(token) //token判断查询者是否登录
	code := 0
	if err != nil {
		code = e.ErrorAuthCheckTokenFail
		return Response{
			StatusCode: int32(code),
			StatusMsg:  "token错误",
		}
	} else if time.Now().Unix() > claims.ExpiresAt {
		code = e.ErrorAuthCheckTokenTimeout
		return Response{
			StatusCode: int32(code),
			StatusMsg:  "token超时",
		}
	}

	//参数解析
	toUserId, err := strconv.ParseInt(toUserIdStr, 10, 64)
	actionType, err := strconv.ParseInt(actionTypeStr, 10, 32)
	if err != nil {
		return Response{
			StatusCode: 400,
			StatusMsg:  "参数解析失败",
		}
	}
	//校验一下合理性
	if actionType != 1 && actionType != 2 {
		return Response{
			StatusCode: 402,
			StatusMsg:  "actionType参数有误",
		}
	}
	var userDao repository.UserRepository
	if user, err := userDao.SelectById(claims.Id); user == nil || err != nil {
		return Response{
			StatusCode: 403,
			StatusMsg:  "userId不存在",
		}
	}
	if user, err := userDao.SelectById(toUserId); user == nil || err != nil {
		return Response{
			StatusCode: 403,
			StatusMsg:  "toUserId不存在",
		}
	}

	//进行数据库相关操作
	var followDao repository.FollowRepository
	flag := followDao.RelationAct(claims.Id, toUserId, int32(actionType))
	if !flag {
		return Response{
			StatusCode: 401,
			StatusMsg:  "数据修改失败",
		}
	}
	return Response{
		StatusCode: 0,
		StatusMsg:  "数据修改成功",
	}
}

// GetFollowListByUId 用户列表
func (service *FollowService) GetFollowListByUId(token, userId string) ([]User, Response) {
	//token验证
	claims, err := util.ParseToken(token) //token判断查询者是否登录
	code := 0
	if err != nil {
		code = e.ErrorAuthCheckTokenFail
		return nil, Response{
			StatusCode: int32(code),
			StatusMsg:  "token错误",
		}
	} else if time.Now().Unix() > claims.ExpiresAt {
		code = e.ErrorAuthCheckTokenTimeout
		return nil, Response{
			StatusCode: int32(code),
			StatusMsg:  "token超时",
		}
	}
	//信息查询
	var followDao = repository.FollowRepository{}
	num, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		return nil, Response{
			StatusCode: 400,
			StatusMsg:  "参数解析失败",
		}
	}
	res, err := followDao.GetFollowListByUId(num)
	if err != nil {
		return nil, Response{
			StatusCode: 401,
			StatusMsg:  err.Error() + ",数据库查询数据失败!",
		}
	}
	//信息过滤
	var userDao repository.UserRepository
	var users = make([]User, 0, 10)
	for _, u := range res {
		var user = User{
			Id:            u.Id,
			Nickname:      u.Nickname,
			FollowerCount: u.FollowerCount,
			FollowCount:   u.FollowCount,
		}
		_, user.IsFollow = userDao.IsFollow(num, u.Id)
		users = append(users, user)
	}
	//返回
	return users, Response{
		StatusCode: 0,
		StatusMsg:  "获取成功",
	}
}

// GetFollowerListByUId 粉丝列表
func (service *FollowService) GetFollowerListByUId(token, userId string) ([]User, Response) {
	//token验证
	claims, err := util.ParseToken(token) //token判断查询者是否登录
	code := e.SuccessUpLoadFile
	if err != nil {
		code = e.ErrorAuthCheckTokenFail
		return nil, Response{
			StatusCode: int32(code),
			StatusMsg:  "token错误",
		}
	} else if time.Now().Unix() > claims.ExpiresAt {
		code = e.ErrorAuthCheckTokenTimeout
		return nil, Response{
			StatusCode: int32(code),
			StatusMsg:  "token超时",
		}
	}
	//信息查询
	var followDao = repository.FollowRepository{}
	num, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		return nil, Response{
			StatusCode: 400,
			StatusMsg:  "参数解析失败",
		}
	}
	res, err := followDao.GetFollowerListByUId(num)
	if err != nil {
		return nil, Response{
			StatusCode: 401,
			StatusMsg:  err.Error() + ",数据库查询数据失败!",
		}
	}
	//信息过滤
	var userDao repository.UserRepository
	var users = make([]User, 0, 10)
	for _, u := range res {
		var user = User{
			Id:            u.Id,
			Nickname:      u.Nickname,
			FollowerCount: u.FollowerCount,
			FollowCount:   u.FollowCount,
		}
		_, user.IsFollow = userDao.IsFollow(u.Id, num)
		users = append(users, user)
	}
	return users, Response{
		StatusCode: 0,
		StatusMsg:  "获取成功",
	}
}
