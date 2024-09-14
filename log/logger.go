/*
 * @Author: Lin Jin Ting
 * @Date: 2024-08-15 14:31:28
 * @LastEditors: Lin Jin Ting
 * @LastEditTime: 2024-08-22 09:34:47
 * @FilePath: \gin-web\log\logger.go
 * @Description:
 *
 * Copyright (c) 2024 by ljt930@gmail.com, All Rights Reserved.
 */
package log

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func InitLogger() {

	encoder := getEncoder()

	level := zap.DebugLevel

	core := zapcore.NewCore(encoder, os.Stdout, level)
	Logger = zap.New(core, zap.AddCaller()).Named("TEST").With(
		zap.String("key1", "FunTester"),
		zap.String("key2", "FunTester--"),
	)

	zap.ReplaceGlobals(Logger) // 替换zap包中全局的logger实例，后续在其他包中只需使用zap.L()调用即可
	return
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()

	encoderConfig.TimeKey = "time"

	encoderConfig.EncodeTime = func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString("[" + time.Format("2006-01-02 15:04:05.000") + "]")
	}

	encoderConfig.EncodeLevel = func(l zapcore.Level, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString("[" + l.CapitalString() + "]")
	}
	encoderConfig.EncodeName = func(name string, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString("[" + name + "]")
	}
	encoderConfig.EncodeCaller = func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
		// enc.AppendString("[" + l.traceId + "]")
		enc.AppendString("[" + caller.TrimmedPath() + "]")
	}

	return zapcore.NewConsoleEncoder(encoderConfig)
}
