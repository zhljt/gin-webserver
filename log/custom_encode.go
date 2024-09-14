package log

import (
	"math"
	"time"

	"github.com/zhljt/webserver-go/log/bufferpool"
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
)

type Pcore struct {
	zapcore.Core
}

func encodeTimeLayout(t time.Time, layout string, enc zapcore.PrimitiveArrayEncoder) {
	type appendTimeEncoder interface {
		AppendTimeLayout(time.Time, string)
	}

	if enc, ok := enc.(appendTimeEncoder); ok {
		enc.AppendTimeLayout(t, layout)
		return
	}

	enc.AppendString(t.Format(layout))
}

func custNameEncoder(loggerName string, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + loggerName + "] ")
}

func NameAndTimeOfLayout(loggerName, layout string) zapcore.TimeEncoder {
	return func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		custNameEncoder(loggerName, enc)
		encodeTimeLayout(t, layout, enc)
	}
}

type TextEncode struct {
	*zapcore.EncoderConfig
	buf            *buffer.Buffer
	spaced         bool // include spaces after colons and commas
	openNamespaces int

	// for encoding generic values by reflection
	// 	reflectBuf *buffer.Buffer
	// 	reflectEnc ReflectedEncoder
}

func NewTextEncode(cfg zapcore.EncoderConfig) *TextEncode {
	if cfg.SkipLineEnding {
		cfg.LineEnding = ""
	} else if cfg.LineEnding == "" {
		cfg.LineEnding = "\n"
	}

	return &TextEncode{
		EncoderConfig: &cfg,
		buf:           bufferpool.Get(),
		spaced:        true,
	}
}

func (enc *TextEncode) AppendComplex64(v complex64)   { enc.appendComplex(complex128(v), 32) }
func (enc *TextEncode) AppendComplex128(v complex128) { enc.appendComplex(complex128(v), 64) }
func (enc *TextEncode) AppendFloat64(v float64)       { enc.appendFloat(v, 64) }
func (enc *TextEncode) AppendFloat32(v float32)       { enc.appendFloat(float64(v), 32) }
func (enc *TextEncode) AppendInt(v int)               { enc.AppendInt64(int64(v)) }
func (enc *TextEncode) AppendInt32(v int32)           { enc.AppendInt64(int64(v)) }
func (enc *TextEncode) AppendInt16(v int16)           { enc.AppendInt64(int64(v)) }
func (enc *TextEncode) AppendInt8(v int8)             { enc.AppendInt64(int64(v)) }
func (enc *TextEncode) AppendUint(v uint)             { enc.AppendUint64(uint64(v)) }
func (enc *TextEncode) AppendUint32(v uint32)         { enc.AppendUint64(uint64(v)) }
func (enc *TextEncode) AppendUint16(v uint16)         { enc.AppendUint64(uint64(v)) }
func (enc *TextEncode) AppendUint8(v uint8)           { enc.AppendUint64(uint64(v)) }
func (enc *TextEncode) AppendUintptr(v uintptr)       { enc.AppendUint64(uint64(v)) }

func (enc *TextEncode) AppendInt64(val int64) {
	enc.addElementSeparator()
	enc.buf.AppendInt(val)
}

func (enc *TextEncode) AppendUint64(val uint64) {
	enc.addElementSeparator()
	enc.buf.AppendUint(val)
}

func (enc *TextEncode) appendFloat(val float64, bitSize int) {
	enc.addElementSeparator()
	switch {
	case math.IsNaN(val):
		enc.buf.AppendString(`"NaN"`)
	case math.IsInf(val, 1):
		enc.buf.AppendString(`"+Inf"`)
	case math.IsInf(val, -1):
		enc.buf.AppendString(`"-Inf"`)
	default:
		enc.buf.AppendFloat(val, bitSize)
	}
}
func (enc *TextEncode) appendComplex(val complex128, precision int) {
	enc.addElementSeparator()
	// Cast to a platform-independent, fixed-size type.
	r, i := float64(real(val)), float64(imag(val))
	enc.buf.AppendByte('"')
	// Because we're always in a quoted string, we can use strconv without
	// special-casing NaN and +/-Inf.
	enc.buf.AppendFloat(r, precision)
	// If imaginary part is less than 0, minus (-) sign is added by default
	// by AppendFloat.
	if i >= 0 {
		enc.buf.AppendByte('+')
	}
	enc.buf.AppendFloat(i, precision)
	enc.buf.AppendByte('i')
	enc.buf.AppendByte('"')
}

func (enc *TextEncode) closeOpenNamespaces() {
	for i := 0; i < enc.openNamespaces; i++ {
		enc.buf.AppendByte('}')
	}
	enc.openNamespaces = 0
}

func (enc *TextEncode) addKey(key string) {
	enc.addElementSeparator()
	enc.buf.AppendByte('"')
	// enc.safeAddString(key)
	enc.buf.AppendByte('"')
	enc.buf.AppendByte(':')
	if enc.spaced {
		enc.buf.AppendByte(' ')
	}
}

func (enc *TextEncode) addElementSeparator() {
	last := enc.buf.Len() - 1
	if last < 0 {
		return
	}
	switch enc.buf.Bytes()[last] {
	case '{', '[', ':', ',', ' ':
		return
	default:
		enc.buf.AppendByte(',')
		if enc.spaced {
			enc.buf.AppendByte(' ')
		}
	}
}
