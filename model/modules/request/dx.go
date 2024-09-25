package request

type DxConfigRequest struct {
	DXID     string `json:"dxId"`     // DXID
	DXName   string `json:"dxName"`   // DX 名称
	DXIP     string `json:"dxIp"`     // DXIP
	DXPort   string `json:"dxPort"`   // DX 端口
	DXType   string `json:"dxType"`   // DX 类型
	ByteCode string `json:"byteCode"` // ByteCode 主体部分16进制字符串
}
