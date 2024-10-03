package initialize

import (
	"fmt"
	"os"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"github.com/zhljt/gin-webserver/global"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

const (
	ConfigEnv         = "CONFIG"
	confgiPrefix      = "config/"
	ConfigDefaultFile = confgiPrefix + "config.yaml"
	ConfigTestFile    = confgiPrefix + "config.test.yaml"
	ConfigDebugFile   = confgiPrefix + "config.debug.yaml"
	ConfigReleaseFile = confgiPrefix + "config.release.yaml"
)

func InitViper() *viper.Viper {
	// TODO: implement viper initialization
	var config string
	switch gin.Mode() {
	case gin.DebugMode:
		config = ConfigDefaultFile
	case gin.TestMode:
		config = ConfigTestFile
	default:
		config = ConfigDefaultFile
	}
	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	// viper.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file : %s", config)
		panic(err)
	}

	if err := v.Unmarshal(&global.G_SystemConfig, viper.DecodeHook(ecodeHookFuncType())); err != nil {
		panic(err)
	}
	return v
}

// 自定义解码钩子函数
func ecodeHookFuncType() mapstructure.DecodeHookFunc {
	return func(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
		switch t {
		case reflect.TypeOf(zap.AtomicLevel{}):
			if f.Kind() == reflect.String {
				if str, ok := data.(string); ok {
					var zapLevel zap.AtomicLevel
					if err := zapLevel.UnmarshalText([]byte(str)); err != nil {
						return nil, err
					}
					return zapLevel, nil
				}
			}
			// 源类型 不为 strint
			return nil, fmt.Errorf("expected string for zap.AtomicLevel, got %T", data)
		default:
			// 对于其他类型，使用默认行为
			return data, nil
		}
	}
}

func InitYaml() {
	// TODO: implement yaml initialization
	var path string
	switch gin.Mode() {
	case gin.DebugMode:
		path = ConfigDefaultFile
	case gin.TestMode:
		path = ConfigTestFile
	default:
		path = ConfigDefaultFile
	}

	// 读取配置文件
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("Error reading config file : %s", path)
		panic(err)
	}

	// 解析配置文件
	config := global.G_SystemConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		panic(err)
	}
}
