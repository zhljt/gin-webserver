package config

type BaseDB struct {
	// 数据库路径
	Path string `json:"path" yaml:"path"`
	// 数据库主机名
	Host string `json:"host" yaml:"host"`
	// 数据库端口
	Port string `json:"port" yaml:"port"`
	// 数据库用户名
	User string `json:"user" yaml:"user"`
	//	数据库密码
	Password string `json:"password" yaml:"password"`
	// 数据库名称
	Database string `json:"database" yaml:"database"`
	// 最大连接数
	MaxOpenConns int `json:"maxOpenConns" yaml:"maxOpenConns"`
	//	最大空闲连接数
	MaxIdleConns int `json:"maxIdleConns" yaml:"maxIdleConns"`
	// 连接最大存活时间
	ConnMaxLifeTime int `json:"connMaxLifeTime" yaml:"connMaxLifeTime"`
	// 是否启用日志
	EnableLog bool `json:"enableLog" yaml:"enableLog"`
	// 是否启用缓存
	// EnableCache bool `json:"enableCache" yaml:"enableCache"`
	// 是否开启全局禁用复数，true表示开启
	Singular bool `json:"singular" yaml:"singular"`
}
