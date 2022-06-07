# <span id="head1">字节训练营 Go语言-抖音实战项目[抖声APP]</span>

## [目录](#toc)

<br/><br/><br/><br/>
## <span id="head2"> 项目基本信息</span>
### <span id="head3"> 小组队号：4096<br/></span>
### <span id="head4">小组队名：CTRL NO.1 队<br/></span>
### <span id="head5"> 成员信息：<br/></span>
| 成员 | 掘金首页 |
| ------------ | ------------ |
|  祁盼  |  https://juejin.cn/user/1130732166332030 |
|  杭朋洁 |  https://juejin.cn/user/2432560267532743 |
|  贺凯恒 |  https://juejin.cn/user/638151599853367 |
|  林叶润 |  https://juejin.cn/user/1645327785668654 |
|  王明星 |  https://juejin.cn/user/1201107293977687 |
|  谢庭宇 |  https://juejin.cn/user/1107921199439965 |
### <span id="head6"> 任务分工</span>
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

<br/>

## <span id="head7"> 运行效果展示</span>

### <span id="head8"> 未登录刷视频及看评论</span>
https://user-images.githubusercontent.com/42311991/172157241-3e7ce023-2362-4265-a443-63facbe14db5.mp4

### <span id="head9"> 注册及自动登录登录、退出账号重新登陆</span>
https://user-images.githubusercontent.com/42311991/172157612-b68fdaa2-4c6e-4fc0-938c-14559176c635.mp4

### <span id="head10"> 查看其他用户信息</span>
https://user-images.githubusercontent.com/42311991/172157797-25a4ad14-a41b-4b49-8068-08b6dbf201bf.mp4

### <span id="head11"> 关注及取消关注用户、用户的粉丝、用户的关注</span>
https://user-images.githubusercontent.com/42311991/172158140-88742b8d-ce25-4e4e-be19-77ba327d55e9.mp4

### <span id="head12"> 点赞及取消点赞</span>
https://user-images.githubusercontent.com/42311991/172158459-f39691f1-6e21-4bec-9530-446a106209c0.mp4

### <span id="head13"> 评论及删除评论</span>
https://user-images.githubusercontent.com/42311991/172158971-900b7f97-4ce8-4b67-8241-01ecfcac8f45.mp4

### <span id="head14"> 发布视频</span>
https://user-images.githubusercontent.com/42311991/172159310-a4c43541-eed5-4acb-a0e7-00eb5b96992a.mp4

### <span id="head15"> 视频封面展示</span>
https://user-images.githubusercontent.com/42311991/172159535-67378e5d-324a-4983-b365-7c9d117aec94.mp4

### <span id="head16"> 连续刷视频操作</span>
https://user-images.githubusercontent.com/42311991/172159818-9f5828cc-2602-47ee-99d2-e3824d7b929f.mp4

<br/>

## <span id="head17"> 项目介绍</span>
使用Go语言、常用框架、Mysql数据库、OSS对象存储等实现[极简版抖音APP（抖声）](https://bytedance.feishu.cn/docs/doccnM9KkBAdyDhg8qaeGlIz7S7# "极简版抖音APP（抖声）")<br/>
本项目存储借助Mysql数据库、OSS对象存储实现（无需在本机安装mysql、redis环境，存储服务由**其他主机提供**，请保持本后端所依附主机接入互联网以接入数据库服务）<br/>
本项目实现了官方APP抖声的接口，为其提供后端服务，接口部分实现详见  [后端接口信息](#head35) <br/>
视频封面借助ffmpeg-go实现由视频自动抽帧生成： https://github.com/u2takey/ffmpeg-go <br/><br/>
### <span id="head18">项目演示地址：http://47.98.251.199:8080/ (在抖声APP高级设置中填入以上地址可以直接接入后端) <br/></span>
### <span id="head19">Linux发布版：[Release下载链接](https://github.com/kevinerr/douyin_demo/releases/tag/1.0 "Release下载链接")<br/></span>

<br/>

## <span id="head20"> 项目安全策略</span>

### <span id="head21"> 1、SQL注入预防</span>
**SQL注入的预防策略有两点**：
#### <span id="head22">1. SQL预编译：</span>

项目中的SQL语句拼接为了实现防注入，全部实现了预编译，示例如下：


    /*SQL语句全部以预编译的方式编写：*/

	err := model.DB.Table("follow").
        Select("user.id, user.nickname, user.follow_count, user.follower_count").
        Joins("join user on follow.follower_id = user.id").
        Where("follow_id = ?", userId).
        Scan(&users).Error

#### <span id="head23">2. 参数类型格式化（format）</span>

项目中的SQL语句拼接为了实现防注入，在预编译的基础上进一步实现了参数format，示例如下：

    /*参数vedioId、userId以int64型参数注入，非string类型：*/

	func (c FavoriteRepository) IsFavorite(vedioId int64, userId int64) bool {
    	var count int64
    	model.DB.Table("favorite").Where("video_id = ? and user_id = ?", vedioId, userId).Count(&count)
    	return count != 0
    }

<br/>

### <span id="head24"> 2、token规则</span>
Token规则的实现借助[**jwt-go**](https://github.com/dgrijalva/jwt-go "https://github.com/dgrijalva/jwt-go") 框架实现
链接[jwt.io](http://jwt.io/ "jwt.io") 对 JSON Web Tokens有很好的介绍
简而言之，它是一个签名的 JSON 对象，可以做一些验证（例如，身份验证）
令牌由三个部分组成，由**.'s** 分隔。

前两部分是 JSON 对象，已经过base64url编码。
第一部分称为**标题**。它包含验证最后一部分签名的必要信息。
中间的部分是**声明**，包含您关心的实际内容。
最后一部分是**签名**，以同样的方式编码。

#### <span id="head25"> token签发</span>

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

#### <span id="head26"> token校验</span>

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

### <span id="head27"> 3、用户越权风险</span>
#### <span id="head28"> 水平越权风险</span>
由于本项目只存在一类用户（User），因此不存在多用户之间的水平越权问题
#### <span id="head29"> 垂直越权风险</span>
用户由**登录态**和**未登录态**两种调用状态，直接可能存在的越权问题由Token规则提供保障，实现方式详见[token规则](#head24)

<br/>

### <span id="head30"> 4、数据库索引检查</span>
数据库的索引作为重要的安全和性能保障工作，由我们团队完成索引的设计和维护以保证数据库高效运行
索引设计可见附件中的[数据库设计在线文档](#head38)或下载附件中的[数据库SQL文件](###- 数据库SQL文件（不附带演示数据）)以查看
 
<br/>

### <span id="head31"> 5、数据库敏感数据保护</span>
用户密码在数据库中以**加盐后的散列值**存储
验证过程由[**jwt-go**](https://github.com/dgrijalva/jwt-go"https://github.com/dgrijalva/jwt-go") 框架辅助实现，salt值不直接存储在Mysql数据库中从而实现数据库去敏感化

示例：

|  userName |  password |
| ------------ | ------------ |
|  starine |  $2a$12$YwGv6/esUXyGvNde9IZRIe8BUJhNFBujNbZKk2WWtNLbRHA6TneeK |

<br/>

## <span id="head32"> 项目部署方式</span>
### <span id="head33"> 编译</span>
在项目根目录下执行：

	SET GOARCH=amd64
	SET GOOS=linux
	go build -o dousheng

根目录下生成文件 dousheng，即编译后的linux的bash可执行文件
### <span id="head34"> 运行</span>
在Linux主机上，首先安装screen命令，然后在sousheng文件所在的目录下执行：

	screen -R dousheng
	./dousheng > ./dousheng.log

然后Ctrl+A+D退出screen即可
输出日志保存在./dousheng.log中

<br/>

## <span id="head35"> 后端接口信息</span>
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

<br/>

## <span id="head36"> 附件下载地址</span>
### <span id="head37"> 数据库SQL文件（不附带演示数据）</span>
https://paper-boot.oss-cn-shanghai.aliyuncs.com/test/douyin_db.sql
### <span id="head38"> 数据库设计在线文档</span>
https://hng36mno7d.feishu.cn/docs/doccnfKsXYpPrqNHdAT7EPM3Fid
### <span id="head39"> 汇报文档下载（PPT）</span>
https://paper-boot.oss-cn-shanghai.aliyuncs.com/test/%E6%8A%96%E9%9F%B3%E5%AE%9E%E6%88%98%E9%A1%B9%E7%9B%AE%E5%B1%95%E7%A4%BA%E6%B1%87%E6%8A%A5%E6%96%87%E6%A1%A3.pptx
### <span id="head40"> 官方APP下载地址</span>
https://bytedance.feishu.cn/docs/doccnM9KkBAdyDhg8qaeGlIz7S7#
### <span id="head41"> 官方在线接口文档</span>
https://www.apifox.cn/apidoc/shared-8cc50618-0da6-4d5e-a398-76f3b8f766c5/api-18345145

<br/>

## <span id="head42"> 依赖安装方式</span>
### <span id="head43"> Linux下ffmpeg的安装</span>
参考链接：https://zhuanlan.zhihu.com/p/416620143

<br/><br/><br/><br/>
------------

<span id="toc">**README跳转目录**</span>

- [跳转目录](#head1)
	- [ 项目基本信息](#head2)
		- [ 小组队号](#head3)
		- [小组队名](#head4)
		- [ 成员信息](#head5)
		- [ 任务分工](#head6)
	- [ 运行效果展示](#head7)
		- [ 未登录刷视频及看评论](#head8)
		- [ 注册及自动登录登录、退出账号重新登陆](#head9)
		- [ 查看其他用户信息](#head10)
		- [ 关注及取消关注用户、用户的粉丝、用户的关注](#head11)
		- [ 点赞及取消点赞](#head12)
		- [ 评论及删除评论](#head13)
		- [ 发布视频](#head14)
		- [ 视频封面展示](#head15)
		- [ 连续刷视频操作](#head16)
	- [ 项目介绍](#head17)
		- [项目演示地址](#head18)
		- [Linux发布版](#head19)
	- [ 项目安全策略](#head20)
		- [ SQL注入预防](#head21)
			- [1. SQL预编译：](#head22)
			- [2. 参数类型格式化（format）](#head23)
		- [ token规则](#head24)
			- [ token签发](#head25)
			- [ token校验](#head26)
		- [ 用户越权风险](#head27)
			- [ 水平越权风险](#head28)
			- [ 垂直越权风险](#head29)
		- [ 数据库索引检查](#head30)
		- [ 数据库敏感数据保护](#head31)
	- [ 项目部署方式](#head32)
		- [ 编译](#head33)
		- [ 运行](#head34)
	- [ 后端接口信息](#head35)
	- [ 附件下载地址](#head36)
		- [ 数据库SQL文件](#head37)
		- [ 数据库设计在线文档](#head38)
		- [ 汇报文档下载（PPT）](#head39)
		- [ 官方APP下载地址](#head40)
		- [ 官方在线接口文档](#head41)
	- [ 依赖安装方式](#head42)
		- [ Linux下ffmpeg的安装](#head43)
