package system

// Apis
type Apis struct {
	// apiID
	ID uint `json:"id"  gorm:"column:id;type:int unsigned not null;primary_key;comment:apiID"`
	// api组件
	Component string `json:"component"  gorm:"column:component;type:varchar(32) not null;default:'';comment:api组件"`
	// api描述
	Desc string `json:"desc"  gorm:"column:desc;type:varchar(255) not null;default:'';comment:api描述"`
	// api分组
	Group string `json:"group"  gorm:"column:group;type:varchar(32) not null;default:'system';comment:api分组"`
	// api路径
	Path string `json:"path"  gorm:"column:path;type:varchar(64) not null;default:'';comment:api路径"`
	// api版本
	Version string `json:"version"  gorm:"column:version;type:varchar(5) not null;default:'V1';comment:api版本"`
}

func (Apis) TableName() string {
	return "sys_apis"
}
