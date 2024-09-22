package global

import (
	"github.com/spf13/viper"
	"github.com/zhljt/gin-webserver/config"
	"gorm.io/gorm"
)

var (
	// Global variable to store the configuration
	GormDB       *gorm.DB
	Viper        *viper.Viper
	SystemConfig *config.SystemConfig
)

const (
	// 系错误码
	ARG_ERROR      = 1001000 // 参数错误
	ARG_REQ_ERROR  = 1001001 // 参数缺失
	ARG_BIND_ERROR = 1001002 // 参数绑定错误
	ARG_TYPE_ERROR = 1001003 // 参数类型错误

	DB_ERROR         = 1002000 // 数据库错误
	DB_CONNECT_ERROR = 1002001 // 数据库连接错误
	DB_TIMEOUT_ERROR = 1002002 // 数据库操作超时

	CACHE_ERROR         = 1003000 // 缓存错误
	CACHE_CONNECT_ERROR = 1003001 // 缓存连接错误

	AUTH_ERROR                = 1004000 // 认证错误
	AUTH_TOKEN_ERROR          = 1004100 // 认证token错误
	LOGIN_AUTH_ERROR          = 1004101 // 登录认证错误
	REGISTER_USER_EXIST_ERROR = 1004102 // 用户注册失败
	CHANGE_PWD_ERROR          = 1004103 // 修改密码失败
	RESET_PWD_ERROR           = 1004104 //重置密码失败

	SYS_ERROR = 1000001 // 系统错误
)
