package easy_serve

import (
	"os"

	"github.com/Sunqi43797189/easy_serve/config"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var loggerMap = make(map[string]*zap.Logger)

func initLogger() {
	for _, logConf := range config.C.Log {
		// 日志分割
		hook := lumberjack.Logger{
			Filename:   logConf.LogPath,  // 日志文件路径，默认 os.TempDir()
			MaxSize:    logConf.MaxSize,  // 每个日志文件保存10M，默认 100M
			MaxBackups: 30,               // 保留30个备份，默认不限
			MaxAge:     7,                // 保留7天，默认不限
			Compress:   logConf.Compress, // 是否压缩，默认不压缩
		}
		write := zapcore.AddSync(&hook)
		var level zapcore.Level
		switch logConf.Level {
		case "debug":
			level = zap.DebugLevel
		case "info":
			level = zap.InfoLevel
		case "error":
			level = zap.ErrorLevel
		default:
			level = zap.InfoLevel
		}
		encoderConfig := zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "linenum",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
			EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
			EncodeDuration: zapcore.SecondsDurationEncoder, //
			EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
			EncodeName:     zapcore.FullNameEncoder,
		}
		atomicLevel := zap.NewAtomicLevel()
		atomicLevel.SetLevel(level)
		core := zapcore.NewCore(
			// zapcore.NewConsoleEncoder(encoderConfig),
			zapcore.NewJSONEncoder(encoderConfig),
			zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), write), // 打印到控制台和文件
			level,
		)
		// 开启开发模式，堆栈跟踪
		options := []zap.Option{}
		options = append(options, zap.AddCaller(), zap.Development())

		// 构造日志
		logger := zap.New(core, options...)
		loggerMap[logConf.Name] = logger
	}
}
