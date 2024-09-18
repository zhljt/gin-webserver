package system

import "github.com/zhljt/gin-webserver/model"

// User
type User struct {
	// 用户ID
	ID uint `json:"id"  gorm:"column:id;type:int unsigned not null;primary_key;comment:用户ID"`
	// 账号
	Account string `json:"account"  gorm:"column:account;type:varchar(128) not null;default:'';comment:账号"`
	// 用户名
	Name string `json:"name"  gorm:"column:name;type:varchar(200) not null;default:'';comment:用户名"`
	// 密码
	Password string `json:"password"  gorm:"column:password;type:varchar(128) not null;default:'';comment:密码"`
	// 手机号
	Phone string `json:"phone,omitempty"  gorm:"column:phone;type:varchar(20);comment:手机号"`
	// 状态，0正常，1锁定，2删除
	Status uint `json:"status"  gorm:"column:status;type:int unsigned not null;default:1;comment:用户状态 0锁定,1正常,2删除"`
	// 角色ID
	RoleID uint `json:"role_id"  gorm:"foreignKey:role_id;references:role_id;comment:角色ID"`

	Roles []Role `json:"roles" gorm:"many2many:sys_user_role;"`
	// 基础记录
	model.RowRecord
}

func (User) TableName() string {
	return "sys_users"
}

// 例子
// type Category struct {
//     Id          uint    `json:"id" gorm:"column:id;type:int(10) unsigned not null AUTO_INCREMENT;primary_key"`
//     Title       string `json:"title" gorm:"column:title;type:varchar(250) not null;default:''"`
//     Description string `json:"description" gorm:"column:description;type:varchar(250) not null;default:''"`
//     Content     string `json:"content" gorm:"column:content;type:longtext default null"`
//     ParentId    uint    `json:"parent_id" gorm:"column:parent_id;type:int(10) unsigned not null;default:0;index:idx_parent_id"`
//     Status      uint   `json:"status" gorm:"column:status;type:tinyint(1) unsigned not null;default:0;index:idx_status"`
//     CreatedTime int64   `json:"created_time" gorm:"column:created_time;type:int(11) not null;default:0;index:idx_created_time"`
//     UpdatedTime int64   `json:"updated_time" gorm:"column:updated_time;type:int(11) not null;default:0;index:idx_updated_time"`
//     DeletedTime int64   `json:"-" gorm:"column:deleted_time;type:int(11) not null;default:0"`
// }
