package service

import (
	"strings"

	"github.com/zhljt/gin-webserver/tool"
)

type DXInfo struct {
	ID   string
	IP   string
	Port string
}

type DXService struct{}

func (dxs *DXService) GenerateConfig(info DXInfo) (string, error) {
	id := tool.NumToHex(info.ID)
	ip := tool.IPToHex(info.IP)
	port := tool.NumToHex(info.Port)
	mod := tool.NumToHex("0")
	other := "00000000"
	head := "5AA5"
	body := id + ip + port + mod + other
	verify, err := tool.VerifyCRC(body)
	if err != nil {
		return "", err
	}
	dxConfig := strings.ToUpper(head + body + verify)
	return dxConfig, nil
}
