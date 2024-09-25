/*
 * @Author: Lin Jin Ting
 * @LastEditors: Lin Jin Ting
 * @Email: ljt930@gmail.com
 * @Description: server 总入口
 * @Date: 2024-09-14 22:14:10
 * @LastEditTime: 2024-09-22 23:44:15
 */
package service

import (
	"github.com/zhljt/gin-webserver/service/system"
)

type ServiceGroup struct {
	UserService system.UserService
	DXService   system.Dx5GService
}

var ServiceGroupIns = new(ServiceGroup)
