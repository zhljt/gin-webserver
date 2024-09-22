package common

type RowRecord struct {
	UpdatedAt   string `json:"-" gorm:"column:updated_at;type:varchar(100) not null;default:'admin';comment:跟新人"`
	CreatedTime int64  `json:"created_time" gorm:"column:created_time;type:int(11) not null;default:0;comment:创建时间"`
	UpdatedTime int64  `json:"updated_time" gorm:"column:updated_time;type:int(11) not null;default:0;comment:更新时间"`
	DeletedTime int64  `json:"-" gorm:"column:deleted_time;type:int(11) not null;default:0;comment:删除时间"`
}
