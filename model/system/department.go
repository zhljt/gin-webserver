package system

// Department
type Department struct {
	// 部门ID
	ID uint `json:"id"  gorm:"column:id;type:int unsigned not null;primary_key;comment:部门ID"`
	// 部门编码
	Code string `json:"code"  gorm:"column:code;type:varchar(32) not null;default:'';comment:部门编码"`
	// 部门名称
	Name string `json:"name"  gorm:"column:name;type:varchar(200) not null;default:'';comment:部门名称"`
	// 父部门ID
	ParentID uint `json:"parent_id"  gorm:"column:parent_id;type:int unsigned;comment:父部门ID"`
	// 状态,0禁用,1启用,2删除
	Status uint `json:"status"  gorm:"column:status;type:int unsigned not null;default:0;comment:状态,0禁用,1启用,2删除"`
	// 类型,0单位,1部门
	Type uint `json:"type"  gorm:"column:type;type:int unsigned not null;default:0;comment:类型,0单位,1部门"`
}

func (Department) TableName() string {
	return "sys_department"
}
