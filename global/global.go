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
