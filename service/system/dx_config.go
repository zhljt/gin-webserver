/*
 * @Author: Lin Jin Ting
 * @LastEditors: Lin Jin Ting
 * @Email: ljt930@gmail.com
 * @Description: DX  服务接口
 * @Date: 2024-09-22 20:59:49
 * @LastEditTime: 2024-09-25 22:16:28
 */
package system

import (
	"github.com/zhljt/gin-webserver/model/modules"
	"github.com/zhljt/gin-webserver/service/system/impls"
)

type Dx5GService interface {
	GetDXInfoByByteCode(info *modules.DXInfo) (*modules.DXInfo, error)
	GetDXInfoByConfig(info *modules.DXInfo) (*modules.DXInfo, error)
	GetByteCode(id, ip, port string) string
}

func NewDxConfigService() Dx5GService {
	return &impls.DX5GServiceImpl{}
}
