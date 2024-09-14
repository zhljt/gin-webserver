package system

import "github.com/gin-web/model"

type UserRole struct {
	// 用户ID
	UserId uint `json:"user_id" gorm:"column:user_id;type int unsigned not null;primary_key;comment:用户ID"`
	// 角色ID
	RoleId uint `json:"role_id" gorm:"column:role_id;type int unsigned not null;primary_key;comment:角色ID"`
	// 基础记录
	model.RowRecord
}
