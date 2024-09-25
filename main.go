/*
 * @Author: Lin Jin Ting
 * @LastEditors: Lin Jin Ting
 * @Email: ljt930@gmail.com
 * @Description:
 * @Date: 2024-04-19 16:27:26
 * @LastEditTime: 2024-09-22 21:12:01
 */
package main

import (
	"github.com/zhljt/gin-webserver/global"
	"github.com/zhljt/gin-webserver/initialize"
	"github.com/zhljt/gin-webserver/router"
	"go.uber.org/zap"
)

func RunSimpleServer() {
	r := router.InitRouters()
	r.Run("0.0.0.0:8000")

}

func testlog() {
	testlog := zap.L().Named("test")
	testlog.Debug("test. debug ddd")
	testlog.Info("test INFO")
	testlog.Warn("test WARNING")
	testlog.Error("test ERROR")
	testlog.Info("get leve" + global.SystemConfig.ZapLog.ZapCores[0].Level.String())
	// testlog.DPanic("AAIAIAIIAIAI")
	// config.LogConfigImpl.SetLevel("console-output", zap.WarnLevel)
	// testlog.Debug("test. debug --2")
	// testlog.Info("test INFO --2")
	// testlog.Warn("test WARNING --2")
	// testlog.Error("test ERROR --2")
	// testlog.DPanic("AAIAIAIIAIAI --2")
}

func main() {
	global.Viper = initialize.InitViper()

	// 初始化日志
	initialize.InitLogger()

	// initialize.InitDB()
	testlog()
	RunSimpleServer()

}
