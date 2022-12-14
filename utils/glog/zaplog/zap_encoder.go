package zaplog

import (
	"go.uber.org/zap/zapcore"
	"time"
)

// 负责设置 encoding 的日志格式
type FormatConfig struct {
	Prefix        string //前缀
	Format        string
	StacktraceKey string
	LogInConsole  bool
	ShowLine      bool
	EncodeLevel   zapcore.LevelEncoder
}

// 通过config获取
func GetConfigEncoder(cfg FormatConfig) zapcore.Encoder {
	// 获取一个指定的的EncoderConfig，进行自定义
	//encodeConfig := zap.NewProductionEncoderConfig()
	encodeConfig := zapcore.EncoderConfig{
		MessageKey: "msg",
		LevelKey:   "level",
		NameKey:    "log",
		TimeKey:    "time",
		CallerKey:  "caller",
		//FunctionKey:    "func",
		StacktraceKey:  cfg.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    cfg.EncodeLevel,
		EncodeTime:     cfg.PrefixTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}

	if cfg.Format == "json" {
		//json格式
		return zapcore.NewJSONEncoder(encodeConfig)
	}
	return zapcore.NewConsoleEncoder(encodeConfig)
}

// 控制台日志格式，大写带颜色
func GetConsoleEncoder(cfg FormatConfig) zapcore.Encoder {
	// 获取一个指定的的EncoderConfig，进行自定义
	//encodeConfig := zap.NewProductionEncoderConfig()
	encodeConfig := zapcore.EncoderConfig{
		MessageKey: "msg",
		LevelKey:   "level",
		NameKey:    "log",
		TimeKey:    "time",
		CallerKey:  "caller",
		//FunctionKey:    "func",
		StacktraceKey:  cfg.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     cfg.PrefixTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
	//控制台格式
	return zapcore.NewConsoleEncoder(encodeConfig)
}

// PrefixTimeEncoder 自定义带前缀日志输出时间格式
func (m FormatConfig) PrefixTimeEncoder(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
	encoder.AppendString(m.Prefix + t.Format("2006/01/02-15:04:05.000"))
}

func (m FormatConfig) timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

func (m FormatConfig) timeUnixNano(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendInt64(t.UnixNano() / 1e6)
}
