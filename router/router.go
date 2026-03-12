package router

import "github.com/gin-gonic/gin"

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 公共路由
	public := r.Group("/api")
	{
		public.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}

	// 私有路由（需要中间件）
	private := r.Group("/api/private")
	//private.Use(AuthMiddleware())
	{
		private.GET("/user")
	}

	return r
}
