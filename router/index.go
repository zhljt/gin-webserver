package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	apiv1 "github.com/zhljt/gin-webserver/api/v1"
	"github.com/zhljt/gin-webserver/middleware"
)

func InitRouters() *gin.Engine {

	Router := gin.New()
	Router.Use(middleware.GinLogger(zap.L().Named("gin-web")))

	PublicGroup := Router.Group("pub")
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}
	UserRouters(Router)
	ToolRouters(Router)
	return Router

}

func UserRouters(r *gin.Engine) {

	UserGroup := r.Group("user")
	{
		UserGroup.POST("/login", apiv1.APIGroupPtr.UserApi.Login)
	}

}

func ToolRouters(r *gin.Engine) {

	ToolGroup := r.Group("dx")
	{
		ToolGroup.GET("/config", apiv1.APIGroupPtr.ConfigDXApi.Config5GDX)
	}

}
