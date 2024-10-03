/*
 * @Author: Lin Jin Ting
 * @LastEditors: Lin Jin Ting
 * @Email: ljt930@gmail.com
 * @Description:
 * @Date: 2024-09-18 22:20:34
 * @LastEditTime: 2024-10-01 21:09:47
 */
package config

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// LogConfigImpl 日志配置实现
type ZapLogConfig struct {
	// 是否开发者模式
	Development bool `json:"development" yaml:"development"`
	// 是否启用调用者记录
	DisableCaller bool `json:"disableCaller" yaml:"disableCaller" `
	// 是否启用堆栈跟踪记录
	DisableStacktrace bool `json:"disableStacktrace" yaml:"disableStacktrace" mapstructure:"disableStacktrace"`
	// Sampling *zap.SamplingConfig `json:"sampling" yaml:"sampling"`
	ZapCores []ZapCore `json:"zapCores" yaml:"zapCores" mapstructure:"zapCores"`
}

type ZapCore struct {
	// 定义日志级别
	Level zap.AtomicLevel `json:"level" yaml:"level" `
	// 定义 encoder 类型
	Encoding string `json:"encoding" yaml:"encoding"`
	// 定义 名称
	Name string `json:"name" yaml:"name" `
	// 时间日期格式
	Layout       string       `json:"layout" yaml:"layout"`
	CustomConfig CustomConfig `json:"customConfig" yaml:"customConfig"`
	// 文件输出路径，支持os.Stdout
	OutputPath string `json:"outputPath" yaml:"outputPath"`
}

type CustomConfig struct {
	TimeKey       string `json:"timeKey" yaml:"timeKey"`
	LevelKey      string `json:"levelKey" yaml:"levelKey"`
	NameKey       string `json:"nameKey" yaml:"nameKey"`
	CallerKey     string `json:"callerKey" yaml:"callerKey"`
	MessageKey    string `json:"messageKey" yaml:"messageKey"`
	FunctionKey   string `json:"functionKey" yaml:"functionKey"`
	StacktraceKey string `json:"stacktraceKey" yaml:"stacktraceKey"`
}

// SetLevel 设置日志级别
func (lc *ZapLogConfig) SetLevel(name string, l zapcore.Level) {
	for _, out := range lc.ZapCores {
		if out.Name == name {
			out.Level.SetLevel(l)
		}
	}
}

// GetLevel 获取日志级别
func (lc *ZapLogConfig) GetLevel(name string) zapcore.Level {
	for _, out := range lc.ZapCores {
		if out.Name == name {
			return out.Level.Level()
		}
	}
	return zapcore.InfoLevel
}

// func init() {
// 	path := "config/log.yaml"

// 	// 读取 YAML 配置文件
// 	yamlBytes, err := os.ReadFile(path)
// 	if err != nil {
// 		panic(err)
// 	}

// 	err = yaml.Unmarshal(yamlBytes, LogConfigImpl)

// 	if err != nil {
// 		panic(err)
// 	}
// }
