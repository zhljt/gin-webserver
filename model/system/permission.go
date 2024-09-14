package system

// Permission
type Permission struct {
	// 权限ID
	ID uint `json:"id"  gorm:"column:id;type:int unsigned not null;primary_key;comment:权限ID"`
	// 权限代码
	Code string `json:"code"  gorm:"column:code;type:varchar(32) not null;default:'';comment:权限代码"`
	// 权限名称
	Name string `json:"name"  gorm:"column:name;type:varchar(128) not null;default:'';comment:权限名称"`
	// 权限类型 0菜单,1api,2data
	Type uint `json:"type"  gorm:"column:type;type:int unsigned not null;comment:权限类型,0菜单,1api,2data"`
}

func (Permission) TableName() string {
	return "sys_permissions"
}
