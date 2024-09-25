package request

import (
	"fmt"
	"os"
)

type InitDBRequest struct {
	DBType     string `json:"dbType"`     // 数据库类型
	DBName     string `json:"dbName"`     // 数据库名称
	DBHost     string `json:"dbHost"`     // 数据库连接地址
	DBPort     int    `json:"dbPort"`     // 数据库端口
	DBUser     string `json:"dbUser"`     // 数据库用户名
	DBPassword string `json:"dbPassword"` // 数据库密码
	DBPath     string `json:"dbPath"`     // SQLITE数据库文件路径
}

func (r *InitDBRequest) MysqlEmptyDsn() string {
	if r.DBHost == "" {
		r.DBHost = "localhost"
	}
	if r.DBPort == 0 {
		r.DBPort = 3306
	}

	return fmt.Sprintf("%s:%s@tcp(%s:%d)/?charset=utf8mb4&parseTime=True&loc=Local", r.DBUser, r.DBPassword, r.DBHost, r.DBPort)
}

func (r *InitDBRequest) MysqlDsn() string {
	if r.DBHost == "" {
		r.DBHost = "localhost"
	}
	if r.DBPort == 0 {
		r.DBPort = 3306
	}

	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", r.DBUser, r.DBPassword, r.DBHost, r.DBPort, r.DBName)
}

func (r *InitDBRequest) SqliteDsn() string {
	sep := string(os.PathSeparator)
	return r.DBPath + sep + r.DBName + ".db"
}
