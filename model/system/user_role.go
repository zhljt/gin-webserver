/*
 * @Author: Lin Jin Ting
 * @LastEditors: Lin Jin Ting
 * @Email: ljt930@gmail.com
 * @Description:
 * @Date: 2024-08-18 10:59:59
 * @LastEditTime: 2024-09-22 00:08:38
 */
package system

import "github.com/zhljt/gin-webserver/model/common"

// User
type UserRole struct {
	// 用户ID
	UserID uint `json:"user_id"  gorm:"column:user_id;type:int unsigned not null;primary_key;comment:用户ID"`
	// 角色ID
	RoleID uint `json:"role_id"  gorm:"column:role_id;type:int unsigned not null;primary_key;comment:角色ID"`
	// 基础记录
	common.RowRecord
}

func (UserRole) TableName() string {
	return "sys_user_role"
}
