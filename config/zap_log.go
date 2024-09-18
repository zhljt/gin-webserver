package config

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LogConfig struct {
	// 是否开发者模式
	Development bool `json:"development" yaml:"development"`
	// 是否启用调用者记录
	DisableCaller bool `json:"disableCaller" yaml:"disableCaller"`
	// 是否启用堆栈跟踪记录
	DisableStacktrace bool `json:"disableStacktrace" yaml:"disableStacktrace"`
	// Sampling *zap.SamplingConfig `json:"sampling" yaml:"sampling"`
	Outputs []LogOutput `json:"outputs" yaml:"outputs"`
}

type LogOutput struct {
	// 定义日志级别
	Level zap.AtomicLevel `json:"level" yaml:"level"`
	// 定义 encoder 类型
	Encoding string `json:"encoding" yaml:"encoding"`
	// 定义 名称
	Name string `json:"name" yaml:"name"`
	// 时间日期格式
	Layout        string                `json:"layout" yaml:"layout"`
	EncoderConfig zapcore.EncoderConfig `json:"encoderConfig" yaml:"encoderConfig"`
	// 文件输出路径，支持os.Stdout
	OutputPath string `json:"outputPath" yaml:"outputPath"`
}

// SetLevel 设置日志级别
func (lc *LogConfig) SetLevel(name string, l zapcore.Level) {
	for _, out := range lc.Outputs {
		if out.Name == name {
			out.Level.SetLevel(l)
		}
	}
}

// GetLevel 获取日志级别
func (lc *LogConfig) GetLevel(name string) zapcore.Level {
	for _, out := range lc.Outputs {
		if out.Name == name {
			return out.Level.Level()
		}
	}
	return zapcore.InfoLevel
}

func (lc *LogConfig) Init() {
	for _, out := range lc.Outputs {
		out.Level.SetLevel(out.Level.Level())
		if out.Encoding == "json" {
			out.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		}
		if out.OutputPath == "os.Stdout" {
			out.OutputPath = ""
		}
	}

	for _, out := range lc.Outputs {
		sink, _, err := zap.Open(out.OutputPath)
		if err != nil {
			panic(err)
		}
		output := zap.New(sink)
		encoder := genEncoder(output)

		return zapcore.NewCore(encoder, sink, out.Level), nil

		if err != nil {
			return err
		}
		cores = append(cores, core)
	}

}
