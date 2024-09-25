/*
 * @Author: Lin Jin Ting
 * @LastEditors: Lin Jin Ting
 * @Email: ljt930@gmail.com
 * @Description: 用户服务 接口
 * @Date: 2024-09-14 22:14:10
 * @LastEditTime: 2024-09-22 23:26:23
 */
package system

import (
	"github.com/zhljt/gin-webserver/model/system"
	"github.com/zhljt/gin-webserver/service/system/impls"
)

type UserService interface {
	Login(info system.User) (userInfo system.User, err error)
	Register(info system.User) (userInfo system.User, err error)
	ChangePassword(id uint, opassword string, npassword string) (err error)
	ResetPassword(ids []uint) (err error)
	// GetUserInfo(id uint) (userInfo system.User, err error)
}

func NewUserService() UserService {
	return &impls.UserServiceImpl{}
}
