package routes

import (
	"blog/api/v1"
	"blog/middleware"
	"blog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	r.Use(gin.Recovery()) // 不用日志系统了
	r.Use(middleware.Logger())
	r.Use(middleware.Cors())
	// 路由接口
	auth := r.Group("api/v1")
	// 1. 需要鉴权路由
	auth.Use(middleware.JwtToken())
	// 这些接口都需要在登录的时候携带生成的token才有权限进行访问
	// 后续接口会自动的调用中间件JwtToken函数进行判断token
	{
		// 用户模块的路由接口
		auth.PUT("user/:id", v1.EditUser)
		auth.DELETE("user/:id", v1.DeleteUser)

		// 分类模块的路由接口
		auth.POST("category/add", v1.AddCategory)
		auth.PUT("category/:id", v1.EditCate)
		auth.DELETE("category/:id", v1.DeleteCate)

		// 文章模块的路由接口
		// 添加文章
		auth.POST("article/add", v1.AddArticle)

		// 编辑文章
		auth.PUT("article/:id", v1.EditArticle)
		// 删除文章
		auth.DELETE("article/:id", v1.DeleteArticle)

		// 上传文件
		auth.POST("upload", v1.Upload)
	}
	// 2. 不需要鉴权路由
	router := r.Group("api/v1")
	{
		router.POST("user/add", v1.AddUser)
		router.GET("users", v1.GetUsers)
		router.GET("category", v1.GetCates)
		// 查询所有的文章
		router.GET("article", v1.GetArticles)
		// 查询指定分类下的所有文章
		router.GET("article/list/:id", v1.GetCatArticleList)
		// 查询单个文章详情
		router.GET("article/info/:id", v1.GetArtInfo)
		// 登录
		router.POST("login", v1.Login)
	}

	r.Run(utils.HttpPort)
}
