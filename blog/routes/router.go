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
	r.Use(gin.Recovery()) // 不用日志中间件
	r.Use(middleware.Logger())
	r.Use(middleware.Cors())

	// 托管静态资源
	//r.LoadHTMLGlob("static/admin/index.html")
	//r.Static("admin/static", "static/admin/static")
	//r.StaticFile("admin/favicon.icon", "static/admin/favicon.icon")
	//r.GET("admin", func(c *gin.Context) {
	//	c.HTML(200, "index.html", nil)
	//})

	// 路由
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

		// 标签
		// 1. 添加标签
		//auth.POST("tag/add", v1.AddTag)
	}
	// 2. 不需要鉴权路由
	router := r.Group("api/v1")
	{
		//
		router.POST("user/add", v1.AddUser)
		router.GET("users", v1.GetUsers)
		// 查询单个参数，用于编辑用户接口
		router.GET("user/:id", v1.GetUserInfo)
		router.GET("category/:id", v1.GetCateInfo)
		router.GET("category", v1.GetCates)
		// 查询所有的文章
		router.GET("article", v1.GetArticles)
		// 查询指定分类下的所有文章
		router.GET("article/list/:id", v1.GetCatArticleList)
		// 查询单个文章详情
		router.GET("article/info/:id", v1.GetArtInfo)
		// 登录
		router.POST("login", v1.Login)

		// FIXME:需要权限
		router.POST("tag/add", v1.AddTag)
		router.GET("tags", v1.GetAllTags)
		router.DELETE("tag/:id", v1.DeleteTag)
		router.PUT("tag/:id", v1.EditTag)
	}

	r.Run(utils.HttpPort)
}
