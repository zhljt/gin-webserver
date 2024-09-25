/*
 * @Author: Lin Jin Ting
 * @LastEditors: Lin Jin Ting
 * @Email: ljt930@gmail.com
 * @Description:
 * @Date: 2024-09-14 22:14:10
 * @LastEditTime: 2024-09-25 22:43:25
 */
package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	model_comm "github.com/zhljt/gin-webserver/model/common"
	model_modules "github.com/zhljt/gin-webserver/model/modules"
	"github.com/zhljt/gin-webserver/service"
)

type ConfigDXApi struct{}
type ConfigDXReq struct {
	ID   string `form:"id" binding:"required"`
	IP   string `form:"ip" binding:"required"`
	Port string `form:"port" binding:"required"`
}

func (dxapi *ConfigDXApi) Config5GDX(c *gin.Context) {
	req := &ConfigDXReq{}
	err := c.ShouldBindQuery(req)
	if err != nil {
		c.JSON(http.StatusOK, map[string]string{"msg": "parse json obj error", "code": "400101", "body": err.Error()})
		return
	}
	// ip, ok := c.GetQuery("ip")
	// if !ok {
	// 	c.JSON(http.StatusOK, map[string]string{"msg": "parse json obj error", "code": "400101", "body": "not found dx server ip "})
	// 	return
	// }
	// port, ok := c.GetQuery("port")
	// if !ok {
	// 	c.JSON(http.StatusOK, map[string]string{"msg": "parse json obj error", "code": "400101", "body": "not found dx server port "})
	// 	return
	// }

	info := &model_modules.DXInfo{}
	info.DXID = req.ID
	info.DXIP = req.IP
	info.DXPort = req.Port
	data, err := service.ServiceGroupIns.DXService.GetDXInfoByConfig(info)
	if err != nil {
		c.JSON(http.StatusOK, model_comm.ErrorWithCodeAndMessage(111111, "登录失败"))
		return
	}
	c.JSON(http.StatusOK, model_comm.SuccessWithComplete("登录成功", data))
}
