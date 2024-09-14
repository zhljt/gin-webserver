package system

import "github.com/zhljt/webserver-go/model"

// User
type UserRole struct {
	// 用户ID
	UserID uint `json:"user_id"  gorm:"column:user_id;type:int unsigned not null;primary_key;comment:用户ID"`
	// 角色ID
	RoleID uint `json:"role_id"  gorm:"column:role_id;type:int unsigned not null;primary_key;comment:角色ID"`
	// 基础记录
	model.RowRecord
}

func (UserRole) TableName() string {
	return "sys_user_role"
}
