/*
 * @Author: Lin Jin Ting
 * @Email: ljt930@gmail.com
 * @Date: 2024-09-19 08:54:10
 * @Last Modified by: linjinting@gs
 * @Last Modified time: 2024-09-19 16:48:05
 * @FilePath: \gin-web\initialize\logger.go
 * @Description:
 * Copyright (c) 2024 by ljt930@gmail.com, All Rights Reserved.
 */
package initialize

import (
	"time"

	"github.com/zhljt/gin-webserver/config"
	"github.com/zhljt/gin-webserver/global"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type _zapLogger struct {
	*config.ZapCore
}

func InitLogger() error {
	opt := initOptions()
	lg := zap.New(getAllCores(), opt...)

	// 替换zap包中全局的logger实例，后续在其他包中只需使用zap.L()调用即可
	zap.ReplaceGlobals(lg)
	return nil
}

func getAllCores() zapcore.Core {
	cores := make([]zapcore.Core, 0)
	for _, out := range global.SystemConfig.ZapLog.ZapCores {
		zl := &_zapLogger{
			ZapCore: out,
		}
		core, err := zl.GetCore()
		if err != nil {
			continue
		}
		cores = append(cores, core)
	}

	return zapcore.NewTee(cores...)
}

func initOptions() []zap.Option {
	var opts []zap.Option
	// opts = append(opts, zap.WrapCore(core))
	if global.SystemConfig.ZapLog.DisableCaller {
		opts = append(opts, zap.Development())
	}
	if global.SystemConfig.ZapLog.DisableCaller {
		opts = append(opts, zap.AddCaller(), zap.AddCallerSkip(1))
	}
	if global.SystemConfig.ZapLog.DisableStacktrace {
		opts = append(opts, zap.AddStacktrace(zap.ErrorLevel))

	}
	return opts
}

func (z *_zapLogger) GetCore() (zapcore.Core, error) {
	writer, _, err := z.GetWriteSyncer()
	if err != nil {
		return nil, err
	}
	encoder := z.GetEncoder()

	return zapcore.NewCore(encoder, writer, z.Level), nil

}

func (z *_zapLogger) GetEncoder() zapcore.Encoder {
	if z.Encoding == "json" {
		return zapcore.NewJSONEncoder(z.GetEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(z.GetEncoderConfig())
}

func (z *_zapLogger) GetEncoderConfig() zapcore.EncoderConfig {
	encoderConfig := z.EncoderConfig
	encoderConfig.EncodeTime = z.customTimeEncoder(z.Layout)
	encoderConfig.EncodeLevel = z.customLevelEncoder()
	encoderConfig.EncodeName = z.customNameEncoder()
	encoderConfig.EncodeCaller = z.customCallerEncoder()
	return encoderConfig
}

func (z *_zapLogger) customTimeEncoder(layout string) zapcore.TimeEncoder {
	return func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString("[" + time.Format(layout) + "]")
	}
}

func (z *_zapLogger) customLevelEncoder() zapcore.LevelEncoder {
	return func(l zapcore.Level, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString("[" + l.CapitalString() + "]")
	}
}

func (z *_zapLogger) customNameEncoder() zapcore.NameEncoder {
	return func(name string, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString("[" + name + "]")
	}
}

func (z *_zapLogger) customCallerEncoder() zapcore.CallerEncoder {
	return func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
		// enc.AppendString("[" + l.traceId + "]")
		enc.AppendString("[" + caller.TrimmedPath() + "]")
	}
}

func (z *_zapLogger) GetWriteSyncer() (zapcore.WriteSyncer, func(), error) {
	return zap.Open(z.OutputPath)

}
