package log

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/yaml.v2"
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

var LogConfigImpl = new(LogConfig)

func InitLogger() error {
	err := loadConfig("config/log.yaml")
	if err != nil {
		return err
	}

	var cores []zapcore.Core
	var lg *zap.Logger
	for _, out := range LogConfigImpl.Outputs {
		core, err := genCore(out)

		if err != nil {
			return err
		}
		cores = append(cores, core)
	}
	core := zapcore.NewTee(cores...)
	var opts []zap.Option
	// opts = append(opts, zap.WrapCore(warecore))
	if LogConfigImpl.DisableCaller {
		opts = append(opts, zap.Development())
	}
	if LogConfigImpl.DisableCaller {
		opts = append(opts, zap.AddCaller(), zap.AddCallerSkip(1))
	}
	if LogConfigImpl.DisableStacktrace {
		opts = append(opts, zap.AddStacktrace(zap.ErrorLevel))

	}
	lg = zap.New(core, opts...)

	// 替换zap包中全局的logger实例，后续在其他包中只需使用zap.L()调用即可
	zap.ReplaceGlobals(lg)
	return nil
}

func loadConfig(path string) error {
	// 读取 YAML 配置文件
	yamlBytes, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(yamlBytes, LogConfigImpl)

	if err != nil {
		return err
	}
	return nil
}

func genEncoder(output LogOutput) zapcore.Encoder {
	// encoderConfig1 := zap.NewProductionEncoderConfig()
	encoderConfig := output.EncoderConfig

	encoderConfig.EncodeTime = customTimeEncoder(output.Layout)
	encoderConfig.EncodeLevel = customLevelEncoder()
	encoderConfig.EncodeName = customNameEncoder()
	encoderConfig.EncodeCaller = customCallerEncoder()
	if output.Encoding == "json" {
		return zapcore.NewJSONEncoder(encoderConfig)
	}
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func genCore(output LogOutput) (zapcore.Core, error) {
	sink, _, err := zap.Open(output.OutputPath)
	if err != nil {
		return nil, err
	}
	encoder := genEncoder(output)

	return zapcore.NewCore(encoder, sink, output.Level), nil

}

func customTimeEncoder(layout string) zapcore.TimeEncoder {
	return func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString("[" + time.Format(layout) + "]")
	}
}

func customLevelEncoder() zapcore.LevelEncoder {
	return func(l zapcore.Level, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString("[" + l.CapitalString() + "]")
	}
}

func customNameEncoder() zapcore.NameEncoder {
	return func(name string, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString("[" + name + "]")
	}
}

func customCallerEncoder() zapcore.CallerEncoder {
	return func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
		// enc.AppendString("[" + l.traceId + "]")
		enc.AppendString("[" + caller.TrimmedPath() + "]")
	}
}

func (c *LogConfig) SetLevel(name string, l zapcore.Level) {
	for _, out := range c.Outputs {
		if out.Name == name {
			out.Level.SetLevel(l)
		}
	}
}
