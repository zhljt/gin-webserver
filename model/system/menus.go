package system

// Menus
type Menus struct {
	// 菜单ID
	ID uint `json:"id"  gorm:"column:id;type:int unsigned not null;primary_key;comment:菜单ID"`
	// 菜单组件
	Component string `json:"component"  gorm:"column:component;type:varchar(32) not null;default:'';comment:菜单组件"`
	// 菜单名称
	Name string `json:"name"  gorm:"column:name;type:varchar(128) not null;default:'';comment:菜单名称"`
	// 父菜单ID
	ParentID uint `json:"parent_id"  gorm:"column:parent_id;type:int unsigned;comment:父菜单ID"`
	// 菜单路径
	Path string `json:"path"  gorm:"column:path;type:varchar(64) not null;default:'';comment:菜单路径"`
	// 排序
	Sort uint `json:"sort"  gorm:"column:sort;type:int unsigned;default:0;comment:排序"`
	// 类型,1:菜单,2:按钮
	Type uint `json:"type"  gorm:"column:type;type:int unsigned not null;comment:类型,1:菜单,2:按钮"`
	// 是否可见 0不可见,1可见
	Visiable bool `json:"visiable"  gorm:"column:visiable;type:tinyint(1) unsigned not null;default:1;comment:是否可见 0不可见,1可见"`
	// 子菜单
	Children []*Menus `json:"children"  gorm:"-"`
	// 菜单元数据
	Meta
}

// 菜单元数据
type Meta struct {
	// 图标
	Icon string `json:"icon"  gorm:"column:icon;type:varchar(255);comment:图标"`
	// 菜单是否缓存
	KeepAlive bool `json:"keepAlive"  gorm:"column:keepAlive;type:tinyint(1) unsigned;comment:是否缓存"`
	// 菜单标志
	Tag uint `json:"tag"  gorm:"column:tag;type:int unsigned ;comment:菜单标志"`
	// 菜单Title
	Title string `json:"title"  gorm:"column:title;type:varchar(128) ;comment:菜单Title"`
}

func (Menus) TableName() string {
	return "sys_menus"
}
