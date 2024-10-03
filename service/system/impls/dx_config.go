/*
 * @Author: Lin Jin Ting
 * @LastEditors: Lin Jin Ting
 * @Email: ljt930@gmail.com
 * @Description: 5G DX 配置参数 生成服务
 * @Date: 2024-09-14 22:14:10
 * @LastEditTime: 2024-09-22 20:59:10
 */
package impls

import (
	"strings"

	"github.com/zhljt/gin-webserver/global"
	"github.com/zhljt/gin-webserver/model/modules"
	"github.com/zhljt/gin-webserver/tool"
)

type DX5GServiceImpl struct{}

/**
 * @description: 通过本地部分16进制编码获取地线信息
 * @param {*modules.DXInfo} info
 * @return {*modules.DXInfo} info
 * @return {error} *
 */
func (dxs *DX5GServiceImpl) GetDXInfoByByteCode(info *modules.DXInfo) (*modules.DXInfo, error) {
	lg := global.G_ZapLogger.Named("DX5GService")
	lg.Debug("GetDXInfoByByteCode: " + info.ByteCode)
	dx_crc32, err := tool.VerifyCRC(info.ByteCode)
	lg.Debug("VerifyCRC: " + dx_crc32)
	if err != nil {
		return nil, err
	}
	info.DXHeader = "A55A"
	info.DXCRC32 = dx_crc32
	return info, nil
}

func (dxs *DX5GServiceImpl) GetDXInfoByConfig(info *modules.DXInfo) (*modules.DXInfo, error) {
	lg := global.G_ZapLogger.Named("DX5GService")
	lg.Debug("GetDXInfoByConfig, ID: " + info.DXID + " IP: " + info.DXIP + " Port: " + info.DXPort)
	byteCode := dxs.GetByteCode(info.DXID, info.DXIP, info.DXPort)
	info.ByteCode = byteCode
	return dxs.GetDXInfoByByteCode(info)
}

func (dxs *DX5GServiceImpl) GetByteCode(id, ip, port string) string {
	idHex := tool.NumToHex(id)
	ipHex := tool.IPToHex(ip)
	portHex := tool.NumToHex(port)
	moduleHex := tool.NumToHex("0")
	otherHex := "00000000"
	bodyCode := strings.ToUpper(idHex + ipHex + portHex + moduleHex + otherHex)
	return bodyCode
}
