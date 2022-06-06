# 字节训练营 Go语言-抖音实战项目[抖声]
------------
## 运行效果展示

### 未登录刷视频及看评论
https://paper-boot.oss-cn-shanghai.aliyuncs.com/test/%E6%9C%AA%E7%99%BB%E5%BD%95%E5%88%B7%E8%A7%86%E9%A2%91%E5%8F%8A%E7%9C%8B%E8%AF%84%E8%AE%BA.mp4

### 注册及自动登录登录、退出账号重新登陆
https://paper-boot.oss-cn-shanghai.aliyuncs.com/test/%E6%B3%A8%E5%86%8C%E5%8F%8A%E8%87%AA%E5%8A%A8%E7%99%BB%E5%BD%95%E7%99%BB%E5%BD%95%E3%80%81%E9%80%80%E5%87%BA%E8%B4%A6%E5%8F%B7%E9%87%8D%E6%96%B0%E7%99%BB%E9%99%86.mp4

### 查看其他用户信息
https://paper-boot.oss-cn-shanghai.aliyuncs.com/test/%E6%9F%A5%E7%9C%8B%E5%85%B6%E4%BB%96%E7%94%A8%E6%88%B7%E4%BF%A1%E6%81%AF.mp4

### 关注及取消关注用户、用户的粉丝、用户的关注
https://paper-boot.oss-cn-shanghai.aliyuncs.com/test/%E5%85%B3%E6%B3%A8%E5%8F%8A%E5%8F%96%E6%B6%88%E5%85%B3%E6%B3%A8%E7%94%A8%E6%88%B7%E3%80%81%E7%94%A8%E6%88%B7%E7%9A%84%E7%B2%89%E4%B8%9D%E3%80%81%E7%94%A8%E6%88%B7%E7%9A%84%E5%85%B3%E6%B3%A8.mp4

### 点赞及取消点赞
https://paper-boot.oss-cn-shanghai.aliyuncs.com/test/%E7%82%B9%E8%B5%9E%E5%8F%8A%E5%8F%96%E6%B6%88%E7%82%B9%E8%B5%9E.mp4

### 评论及删除评论
https://paper-boot.oss-cn-shanghai.aliyuncs.com/test/%E8%AF%84%E8%AE%BA%E5%8F%8A%E5%88%A0%E9%99%A4%E8%AF%84%E8%AE%BA.mp4

### 发布视频
https://paper-boot.oss-cn-shanghai.aliyuncs.com/test/%E5%8F%91%E5%B8%83%E8%A7%86%E9%A2%91.mp4

### 视频封面展示
https://paper-boot.oss-cn-shanghai.aliyuncs.com/test/%E8%A7%86%E9%A2%91%E5%B0%81%E9%9D%A2%E5%B1%95%E7%A4%BA.mp4

### 连续刷视频操作
https://paper-boot.oss-cn-shanghai.aliyuncs.com/test/%E8%BF%9E%E7%BB%AD%E5%88%B7%E8%A7%86%E9%A2%91%E6%93%8D%E4%BD%9C.mp4

------------
## 项目基本信息
### 小组队号：4096<br/>
### 小组队名：CTRL NO.1 队<br/>
### 成员信息：<br/>
| 成员 | 掘金首页 |
| ------------ | ------------ |
|  祁盼  |  https://juejin.cn/user/1130732166332030 |
|  杭朋洁 |  https://juejin.cn/user/2432560267532743 |
|  贺凯恒 |  https://juejin.cn/user/638151599853367 |
|  林叶润 |  https://juejin.cn/user/1645327785668654 |
|  王明星 |  https://juejin.cn/user/1201107293977687 |
|  谢庭宇 |  https://juejin.cn/user/1107921199439965 |
### 任务分工
	一、数据库设计维护：@王明星 
	二、接口开发及调试：@贺凯恒 @林叶润 @祁盼 @谢庭宇 
	三、测试和代码审查：@杭朋洁 
	四、安全设计：
		1、讨论设计方案：@杭朋洁 @王明星  @贺凯恒 @林叶润 @祁盼 @谢庭宇 
		2、方案实施：
			[1]、SQL注入预防：@祁盼 @林叶润 
			[2]、token规则检查：@贺凯恒 @王明星 
			[3]、用户越权风险：@谢庭宇 @杭朋洁
			[4]、数据库索引检查：@林叶润 @王明星 
	五、README：@祁盼 
	六、演示文件：@祁盼 @谢庭宇 @杭朋洁 

------------
## 项目介绍
使用Go语言、常用框架、Mysql数据库、OSS对象存储等实现[极简版抖音APP（抖声）](https://bytedance.feishu.cn/docs/doccnM9KkBAdyDhg8qaeGlIz7S7# "极简版抖音APP（抖声）")<br/>
本项目存储借助Mysql数据库、OSS对象存储实现（无需在本机安装mysql、redis环境，存储服务由**其他主机提供**，请保持本后端所依附主机接入互联网以接入数据库服务）<br/>
本项目实现了官方APP抖声的接口，为其提供后端服务，接口部分实现详见 * [后端接口信息](##-后端接口信息) <br/>
视频封面借助ffmpeg-go实现由视频自动抽帧生成： https://github.com/u2takey/ffmpeg-go <br/><br/>
**项目演示地址**：http://47.98.251.199:8080/ (在抖声APP高级设置中填入以上地址可以直接接入后端) <br/>

------------
## 项目安全策略

### 1、SQL注入预防
**SQL注入的预防策略有两点**：
#### 1. SQL预编译：

项目中的SQL语句拼接为了实现防注入，全部实现了预编译，示例如下：


    /*SQL语句全部以预编译的方式编写：*/

	err := model.DB.Table("follow").
        Select("user.id, user.nickname, user.follow_count, user.follower_count").
        Joins("join user on follow.follower_id = user.id").
        Where("follow_id = ?", userId).
        Scan(&users).Error

#### 2. 参数类型格式化（format）

项目中的SQL语句拼接为了实现防注入，在预编译的基础上进一步实现了参数format，示例如下：

    /*参数vedioId、userId以int64型参数注入，非string类型：*/

	func (c FavoriteRepository) IsFavorite(vedioId int64, userId int64) bool {
    	var count int64
    	model.DB.Table("favorite").Where("video_id = ? and user_id = ?", vedioId, userId).Count(&count)
    	return count != 0
    }

<br/>

### 2、token规则
Token规则的实现借助[**jwt-go**](https://github.com/dgrijalva/jwt-go "https://github.com/dgrijalva/jwt-go") 框架实现
链接[jwt.io](http://jwt.io/ "jwt.io") 对 JSON Web Tokens有很好的介绍
简而言之，它是一个签名的 JSON 对象，可以做一些验证（例如，身份验证）
令牌由三个部分组成，由**.'s** 分隔。

前两部分是 JSON 对象，已经过base64url编码。
第一部分称为**标题**。它包含验证最后一部分签名的必要信息。
中间的部分是**声明**，包含您关心的实际内容。
最后一部分是**签名**，以同样的方式编码。

#### token签发

	/*token签发操作*/

	nowTime := time.Now()
	expireTime := nowTime.Add(24000 * time.Hour)
	claims := Claims{
		Id:        id,
		Username:  username,
		Authority: authority,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "Douyin",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err

#### token校验

	/*token校验操作*/

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

<br/>

### 3、用户越权风险
#### 水平越权风险
由于本项目只存在一类用户（User），因此不存在多用户之间的水平越权问题
#### 垂直越权风险
用户由**登录态**和**未登录态**两种调用状态，直接可能存在的越权问题由Token规则提供保障，实现方式详见[token规则](###-2、token规则)

<br/>

### 4、数据库索引检查
数据库的索引作为重要的安全和性能保障工作，由我们团队完成索引的设计和维护以保证数据库高效运行
索引设计可见附件中的[数据库设计在线文档](###-数据库设计在线文档)或下载附件中的[数据库SQL文件](###- 数据库SQL文件（不附带演示数据）)以查看
 
<br/>

### 5、数据库敏感数据保护
用户密码在数据库中以**加盐后的散列值**存储
验证过程由[**jwt-go**](https://github.com/dgrijalva/jwt-go"https://github.com/dgrijalva/jwt-go") 框架辅助实现，salt值不直接存储在Mysql数据库中从而实现数据库去敏感化

示例：

|  userName |  password |
| ------------ | ------------ |
|  starine |  $2a$12$YwGv6/esUXyGvNde9IZRIe8BUJhNFBujNbZKk2WWtNLbRHA6TneeK |

------------
## 项目部署方式
### 编译
在项目根目录下执行：

	SET GOARCH=amd64
	SET GOOS=linux
	go build -o dousheng

根目录下生成文件 dousheng，即编译后的linux的bash可执行文件
### 运行
在Linux主机上，首先安装screen命令，然后在sousheng文件所在的目录下执行：

	screen -R dousheng
	./dousheng > ./dousheng.log

然后Ctrl+A+D退出screen即可
输出日志保存在./dousheng.log中

------------
## 后端接口信息
|  接口 |  地址 |  请求方式 | 作者  |  接口描述 |  水平权限限制 | 垂直权限限制 |
| ------------ | ------------ | ------------ | ------------ | ------------ | ------------ |------------ |
| **I接口** | ------------ | ---- | ---- | ------------ | ---- |---- |
| 视频流接口  | /douyin/feed  | GET  | 贺凯恒  |  不限制登录状态，返回按投稿时间倒序的视频列表，视频数由服务端控制，单次最多30个 | 登录/未登录  | 无  |
| 用户注册  | /douyin/user/register/  | POST  | 贺凯恒  |  新用户注册时提供用户名，密码，昵称即可，用户名需要保证唯一。创建成功后返回用户 id 和权限token | 未登录  | 无  |
| 用户登录  | /douyin/user/login/  | POST  | 贺凯恒  | 通过用户名和密码进行登录，登录成功后返回用户 id 和权限 token  | 未登录  | 无  |
| 用户信息  | /douyin/user/  | GET  | 贺凯恒  |  获取登录用户的 id、昵称，如果实现社交部分的功能，还会返回关注数和粉丝数 | 登录  | 无  |
| 投稿接口  | /douyin/publish/action/  | POST  | 贺凯恒  |  登录用户选择视频上传，自动获取封面 | 登录  | 无  |
| 发布列表  | /douyin/publish/list/  | GET  | 贺凯恒  | 登录用户的视频发布列表，直接列出用户所有投稿过的视频  | 登录  | 无  |
| **II接口** | ------------ | ---- | ---- | ------------ | ---- |---- |
| 点赞操作  | /douyin/favorite/action/  | POST  | 谢庭宇  | 登录用户对视频的点赞和取消点赞操作  | 登录  | 无  |
| 点赞列表  | /douyin/favorite/list/  | GET  | 谢庭宇  | 登录用户的所有点赞视频  | 登录  | 无  |
| 评论操作  | /douyin/comment/action/  | POST  | 祁盼  | 登录用户对视频进行评论 | 登录  | 无  |
| 评论列表  | /douyin/comment/list/  | GET  | 祁盼  |  查看视频的所有评论，按发布时间倒序 | 登录/未登录  | 无  |
| **III接口** | ------------ | ---- | ---- | ------------ | ---- |---- |
| 关注操作  | /douyin/relation/action/  | POST  | 林叶润  | 登录用户对其他用户进行关注/取消关注操作  | 登录  | 无  |
| 关注列表  | /douyin/relatioin/follow/list/  | GET  | 林叶润  | 用户关注的所有用户信息列表  | 登录  | 无  |
| 粉丝列表  | /douyin/relation/follower/list/  | GET  | 林叶润  | 用户所有的粉丝信息列表  | 登录  | 无  |

------------
## 附件下载地址
### 数据库SQL文件（不附带演示数据）
https://paper-boot.oss-cn-shanghai.aliyuncs.com/test/douyin_db.sql
### 数据库设计在线文档
https://hng36mno7d.feishu.cn/docs/doccnfKsXYpPrqNHdAT7EPM3Fid
### 汇报文档下载

### 官方APP下载地址
https://bytedance.feishu.cn/docs/doccnM9KkBAdyDhg8qaeGlIz7S7#
### 官方在线接口文档
https://www.apifox.cn/apidoc/shared-8cc50618-0da6-4d5e-a398-76f3b8f766c5/api-18345145

------------
## 依赖安装方式
### Linux下ffmpeg的安装
参考链接：https://zhuanlan.zhihu.com/p/416620143
