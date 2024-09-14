package tool

import (
	"encoding/hex"
	"fmt"
	"net"
	"strconv"
)

// 把字符串形式ip地址 转换成16进制字符串
func IPToHex(ipStr string) string {
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return ""
	}
	ipv4 := ip.To4()
	if ipv4 == nil {
		return ""
	}
	// 使用hex.Encode将字节切片编码为16进制字符串, 字符串为[]byte类型，还需要使用string() 转换成string类型
	hexStr := make([]byte, hex.EncodedLen(len(ipv4)))
	hex.Encode(hexStr, ipv4)
	// 如果你只需要转换少量的IP地址，或者需要更直接的控制格式化过程，可能会选择使用fmt.Sprintf
	// hexStr := fmt.Sprintf("%02x%02x%02x%02x", ipv4[0], ipv4[1], ipv4[2], ipv4[3])
	return string(hexStr)
}

func NumToHex(numStr string) string {
	// 解析字符串为int
	num, err := strconv.ParseInt(numStr, 10, 64)
	if err != nil {
		return ""
	}

	// Determine the length of the hexadecimal representation
	// 格式化为两位的16进制字符串，使用大写字母和前导零
	var hexStr string
	switch {
	case num < 0xFF:
		hexStr = fmt.Sprintf("%02X", uint8(num))
	case num < 0xFFFF:
		hexStr = fmt.Sprintf("%04X", uint16(num))
	case num < 0xFFFFFF:
		hexStr = fmt.Sprintf("%06X", uint32(num))
	default:
		hexStr = fmt.Sprintf("%016X", num)
	}

	return hexStr
}
