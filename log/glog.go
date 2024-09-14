package log

// import (
// 	"fmt"
// 	"os"

// 	"go.uber.org/zap"
// 	"go.uber.org/zap/zapcore"
// 	"gopkg.in/yaml.v2"
// )

// // 全局日志对象
// var GLoger *GlobalLogger

// type GlobalLogger struct {
// 	*zap.Logger // 服务内部普通日志记录
// 	// Request *FileLogger // 服务内部入库日志记录
// }

// // func (l *GlobalLogger) DEBUG(msg string, fields ...zap.Field) {
// // 	l.Debug(msg, fields...)
// // }

// // func (l *GlobalLogger) INFO(msg string, fields ...zap.Field) {
// // 	l.Info(msg, fields...)
// // }

// // func (l *GlobalLogger) WARNING(msg string, fields ...zap.Field) {
// // 	l.Warn(msg, fields...)
// // }

// // func (l *GlobalLogger) FATAL(msg string, fields ...zap.Field) {
// // 	l.Fatal(msg, fields...)
// // }

// // func (l *GlobalLogger) SYNC() {
// // 	l.Sync()
// // }

// type Config struct {
// 	// 是否启用调用者记录
// 	DisableCaller bool `json:"disableCaller" yaml:"disableCaller"`
// 	// 是否启用堆栈跟踪记录
// 	DisableStacktrace bool `json:"disableStacktrace" yaml:"disableStacktrace"`
// 	// Sampling *zap.SamplingConfig `json:"sampling" yaml:"sampling"`
// 	Outputs []ConfigOutpust `json:"outputs" yaml:"outputs"`
// }

// type ConfigOutpust struct {
// 	// 定义日志级别
// 	Level         zap.AtomicLevel       `json:"level" yaml:"level"`
// 	Encoding      string                `json:"encoding" yaml:"encoding"`
// 	EncoderConfig zapcore.EncoderConfig `json:"encoderConfig" yaml:"encoderConfig"`
// 	OutputPath    string                `json:"outputPath" yaml:"outputPath"`
// }

// var _encoderNameToConstructor = map[string]func(zapcore.EncoderConfig) zapcore.Encoder{
// 	"console": func(encoderConfig zapcore.EncoderConfig) zapcore.Encoder {
// 		return zapcore.NewConsoleEncoder(encoderConfig)
// 	},
// 	"json": func(encoderConfig zapcore.EncoderConfig) zapcore.Encoder {
// 		return zapcore.NewJSONEncoder(encoderConfig)
// 	},
// }

// func warecore(core zapcore.Core) zapcore.Core {

// 	return core
// }

// func ConfigLogger() error {

// 	// 读取 YAML 配置文件
// 	yamlBytes, err := os.ReadFile("config/log.yaml")
// 	if err != nil {
// 		return err
// 	}

// 	var config Config
// 	var cores []zapcore.Core
// 	err = yaml.Unmarshal(yamlBytes, &config)

// 	if err != nil {
// 		return err
// 	}
// 	for _, cfg := range config.Outputs {
// 		cfg.EncoderConfig.EncodeName = custNameEncoder
// 		core, err := cfg.GeneralCore()

// 		if err != nil {
// 			return err
// 		}
// 		cores = append(cores, core)
// 	}
// 	core := zapcore.NewTee(cores...)
// 	var opts []zap.Option
// 	opts = append(opts, zap.WrapCore(warecore))

// 	if config.DisableCaller {
// 		opts = append(opts, zap.AddCaller(), zap.AddCallerSkip(1))
// 	}
// 	if config.DisableStacktrace {
// 		opts = append(opts, zap.AddStacktrace(zap.ErrorLevel))
// 	}
// 	// 5 : 创建logger对象
// 	logger := zap.New(core, opts...)
// 	// 6 : 初始化全局日志对象:
// 	GLoger = &GlobalLogger{logger}

// 	return nil

// }

// func (cfg *ConfigOutpust) GeneralCore() (zapcore.Core, error) {
// 	sink, _, err := zap.Open(cfg.OutputPath)
// 	if err != nil {
// 		return nil, err
// 	}
// 	constructor, ok := _encoderNameToConstructor[cfg.Encoding]
// 	if !ok {
// 		return nil, fmt.Errorf("no encoder registered for name %q", cfg.Encoding)
// 	}
// 	encoder := constructor(cfg.EncoderConfig)
// 	return zapcore.NewCore(encoder, sink, cfg.Level), nil

// }

// func DefaultEncoder() zapcore.EncoderConfig {
// 	return zapcore.EncoderConfig{
// 		// Keys can be anything except the empty string.
// 		MessageKey:     "msg",
// 		TimeKey:        "time",
// 		LevelKey:       "level",
// 		NameKey:        "name",
// 		CallerKey:      "caller",
// 		StacktraceKey:  "stacktrace",
// 		FunctionKey:    "func",
// 		LineEnding:     zapcore.DefaultLineEnding,
// 		EncodeLevel:    zapcore.CapitalLevelEncoder,
// 		EncodeTime:     NameAndTimeOfLayout("test-log", "2006-01-02 15:04:05.000"),
// 		EncodeDuration: zapcore.StringDurationEncoder,
// 		EncodeCaller:   zapcore.ShortCallerEncoder,
// 	}

// }
