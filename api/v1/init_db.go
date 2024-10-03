package v1

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	g "github.com/zhljt/gin-webserver/global"
	model_comm "github.com/zhljt/gin-webserver/model/common"
	model_sys_req "github.com/zhljt/gin-webserver/model/system/request"
	service_sys "github.com/zhljt/gin-webserver/service/system"
)

type InitDBApi struct{}

func (InitDBApi) InitDB(c *gin.Context) {
	// TODO: implement
	var dbReq *model_sys_req.InitDBRequest
	if err := c.ShouldBindJSON(&dbReq); err != nil {
		c.JSON(http.StatusOK, model_comm.ErrorWithCodeAndMessage(g.ARG_BIND_ERROR, "绑定参数错误"))
		return
	}
	ctx := context.Background()
	log := g.G_ZapLogger.Named("InitDB")
	ctx = context.WithValue(ctx, g.LOG_KEY, log)
	// TODO: init db
	err := service_sys.InitDB(ctx, dbReq)
	if err != nil {
		c.JSON(http.StatusOK, model_comm.ErrorWithCodeAndMessage(g.DB_INIT_ERROR, "初始化数据库失败"))
		return
	}
	c.JSON(http.StatusOK, model_comm.SuccessWithMessage("初始化数据库成功"))
}
