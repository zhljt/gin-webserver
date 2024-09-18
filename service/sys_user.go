package service

import (
	"fmt"

	"github.com/gofrs/uuid/v5"
)

type UserInfo struct {
	UUID     uuid.UUID `json:"uuid" gorm:"index;comment:用户UUID"`      // 用户UUID
	Account  string    `json:"account" gorm:"index;comment:用户登录名"`    // 用户登录名
	Password string    `json:"-"  gorm:"comment:用户登录密码"`              // 用户登录密码
	Name     string    `json:"Name" gorm:"default:系统用户;comment:用户名称"` // 用户名称
	RoleID   string    `json:"roleID" gorm:"default:1;coment:角色ID"`   // 角色ID
}
type UserService struct{}

func (US *UserService) Login(info *UserInfo) (userInfo UserInfo, err error) {
	fmt.Println(info)
	fmt.Println("getUserInfoForID&& save new userInfo")

	fmt.Println("checkpasswd")

	fmt.Println("checkpasswd")
	return UserInfo{}, nil
}
