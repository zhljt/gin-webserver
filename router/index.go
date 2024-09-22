/*
 * @Author: Lin Jin Ting
 * @Email: ljt930@gmail.com
 * @Date: 2024-09-19 08:54:10
 * @Last Modified by: linjinting@gs
 * @Last Modified time: 2024-09-19 17:02:57
 * @FilePath: \gin-web\router\index.go
 * @Description:
 * Copyright (c) 2024 by ljt930@gmail.com, All Rights Reserved.
 */
package router

import (
	"net/http"
	"time"

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
		// Test
		PublicGroup.GET("/test", func(c *gin.Context) {
			time.Sleep(3 * time.Second)
			c.JSON(http.StatusOK, "test")
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
		UserGroup.POST("/register", apiv1.APIGroupPtr.UserApi.Register)
		UserGroup.POST("/logout", apiv1.APIGroupPtr.UserApi.Logout)
	}

}

func ToolRouters(r *gin.Engine) {

	ToolGroup := r.Group("dx")
	{
		ToolGroup.GET("/config", apiv1.APIGroupPtr.ConfigDXApi.Config5GDX)
	}

}
