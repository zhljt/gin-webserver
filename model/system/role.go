package system

// Role
type Role struct {
	// 角色ID
	ID uint `json:"id"  gorm:"column:id;type:int unsigned not null;primary_key;comment:角色ID"`
	// 角色编码
	Code string `json:"code"  gorm:"column:code;type:varchar(32) not null;default:'';comment:角色编码"`
	// 角色名称
	Name string `json:"name"  gorm:"column:name;type:varchar(128) not null;default:'';comment:角色名称"`
	// 类型 0普通角色,1超级管理员,2禁用角色
	Type uint `json:"type"  gorm:"column:id;type:in unsigned not null;default:0;comment:0普通角色,1超级管理员,2禁用角色"`
}

func (Role) TableName() string {
	return "sys_roles"
}
