package router

import (
	"github.com/kainhuck/gold/middleware/jwt"
	"github.com/kainhuck/gold/pkg/config"
	"github.com/kainhuck/gold/router/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	gin.SetMode(config.Collection.Server.RunMode)

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("ping", v1.Ping)

		apiv1.POST("/register", v1.Register) // 注册
		apiv1.POST("/login", v1.Login) // 登录

		apiv1.GET("/users", v1.GetUsers)
		apiv1.GET("/user/:id", v1.GetUser)
		apiv1.DELETE("/user/:id", v1.DeleteUser)
	}
	apiv1.Use(jwt.JWT()) // 以下操作需要登录
	{
		apiv1.PUT("/edit_pass", v1.EditPassword)
	}

	return r
}
