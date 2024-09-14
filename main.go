package main

import (
	"fmt"

	"github.com/zhljt/webserver-go/log"
	"go.uber.org/zap"

	"github.com/zhljt/webserver-go/router"
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
	testlog.DPanic("AAIAIAIIAIAI")
	log.LogConfigImpl.SetLevel("console-output", zap.WarnLevel)
	testlog.Debug("test. debug --2")
	testlog.Info("test INFO --2")
	testlog.Warn("test WARNING --2")
	testlog.Error("test ERROR --2")
	testlog.DPanic("AAIAIAIIAIAI --2")
}

func main() {
	err := log.InitLogger()
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	// initialize.InitDB()
	// testlog()
	RunSimpleServer()

}
