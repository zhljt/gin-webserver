/*
 * @Author: Lin Jin Ting
 * @LastEditors: Lin Jin Ting
 * @Email: ljt930@gmail.com
 * @Description:
 * @Date: 2024-09-22 16:57:53
 * @LastEditTime: 2024-09-25 22:38:01
 */
package modules

import "github.com/zhljt/gin-webserver/model/modules/request"

type DXInfo struct {
	*request.DxConfigRequest `json:"dx_config_request"` // DX 配置请求信息
	DXHeader                 string                     `json:"dx_header"` // DX 头部信息
	DXCRC32                  string                     `json:"dx_crc32"`  // DX 32位CRC校验码
}
