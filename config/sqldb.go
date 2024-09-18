package config

type BaseDB struct {
	// 数据库路径
	Path string `json:"path" json:"path"`
	// 数据库主机名
	Host string `json:"host" json:"host"`
	// 数据库端口
	Port string `json:"port" json:"port"`
	// 数据库用户名
	User string `json:"user" json:"user"`
	//	数据库密码
	Password string `json:"password" json:"password"`
	// 数据库名称
	Database string `json:"database" json:"database"`
	// 最大连接数
	MaxOpenConns int `json:"maxOpenConns" json:"maxOpenConns"`
	//	最大空闲连接数
	MaxIdleConns int `json:"maxIdleConns" json:"maxIdleConns"`
	// 连接最大存活时间
	ConnMaxLifeTime int `json:"connMaxLifeTime" json:"connMaxLifeTime"`
	// 是否启用日志
	EnableLog bool `json:"enableLog" json:"enableLog"`
	// 是否启用缓存
	EnableCache bool `json:"enableCache" json:"enableCache"`
	// 是否开启全局禁用复数，true表示开启
	Singular bool `json:"singular" json:"singular"`
}
