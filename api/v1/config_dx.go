package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

	info := service.DXInfo{
		ID:   req.ID,
		IP:   req.IP,
		Port: req.Port,
	}

	data, err := service.ServiceGroupPtr.DXService.GenerateConfig(info)
	if err != nil {
		c.JSON(http.StatusOK, map[string]string{"msg": "计算失败", "code": "400102", "body": err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]string{"msg": "计算成功", "code": "0", "body": data})
}
