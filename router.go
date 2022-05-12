package main

import (
	"github.com/RaymondCode/simple-demo/controller"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initRouter(r *gin.Engine) {
	// public directory is used to serve static resources

	//加载模板 放在配置路由前面
	//r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./public")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) // 开启swag
	apiRouter := r.Group("/douyin")

	// basic apis
	apiRouter.GET("/feed/", controller.Feed)
	apiRouter.POST("/user/register/", controller.Register)
	apiRouter.POST("/user/login/", controller.Login)
	apiRouter.GET("/user/", controller.UserInfo)
	apiRouter.POST("/publish/action/", controller.Publish)
	apiRouter.GET("/publish/list/", controller.PublishList)
	//authed := apiRouter.Group("/") //需要登陆保护
	//authed.Use(middleware.JWT())
	//{
	//
	//}

	// extra apis - I
	apiRouter.POST("/favorite/action/", controller.FavoriteAction)
	apiRouter.GET("/favorite/list/", controller.FavoriteList)
	apiRouter.POST("/comment/action/", controller.CommentAction)
	apiRouter.GET("/comment/list/", controller.CommentList)

	// extra apis - II
	apiRouter.POST("/relation/action/", controller.RelationAction)
	apiRouter.GET("/relation/follow/list/", controller.FollowList)
	apiRouter.GET("/relation/follower/list/", controller.FollowerList)
}
