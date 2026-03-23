package logger

import (
	"fmt"
	"one-dock/app/config"
	"one-dock/pkgs/console"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	*zap.Logger
}

type LogFactor struct {
	Info, Warn, Error, Debug func(data map[string]interface{}, msg ...interface{})
}

var (
	levels = []zapcore.Level{zap.InfoLevel, zap.WarnLevel, zap.ErrorLevel, zap.DebugLevel}
)

func Init(cfg config.LogConfig) *LogFactor {
	logFactor := &LogFactor{}
	loggers := make(map[zapcore.Level]*Logger)

	on := cfg.On
	age := cfg.Age
	size := cfg.Size
	backups := cfg.Backups

	for _, level := range levels {
		logger := newLogger(level, age, size, backups)
		loggers[level] = logger
	}

	logFactor.Info = makeLogFunc(on, loggers[zap.InfoLevel])
	logFactor.Warn = makeLogFunc(on, loggers[zap.WarnLevel])
	logFactor.Error = makeLogFunc(on, loggers[zap.ErrorLevel])
	logFactor.Debug = makeLogFunc(on, loggers[zap.DebugLevel])

	console.Info("日志服务初始化成功！")
	return logFactor
}

func newLogger(level zapcore.Level, age, size, backups int) *Logger {
	path := fmt.Sprintf("./runtime/logs/%s/%s.log",
		time.Now().Format("2006-01-02"),
		level.String())

	writer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   path,
		MaxAge:     age,
		MaxSize:    size,
		MaxBackups: backups,
	})

	encoder := zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		TimeKey:      "time",
		CallerKey:    "caller", // 指定 caller 字段名
		EncodeTime:   zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000"),
		EncodeCaller: zapcore.ShortCallerEncoder, // 短路径格式（如：logger/logger.go:100）
		// 如果需要完整路径，可改用：zapcore.FullCallerEncoder
		LineEnding:  zapcore.DefaultLineEnding,
		LevelKey:    "level", // 补充 level 字段（可选，增强日志可读性）
		EncodeLevel: zapcore.LowercaseLevelEncoder,
		// MessageKey:  "msg", // 补充 msg 字段（可选）
	})

	// 关键修改2：创建 zap.Logger 时添加 AddCaller() 选项，启用调用者信息采集
	core := zapcore.NewCore(encoder, writer, level)
	logger := zap.New(
		core,
		// zap.AddCaller(),      // 启用 caller 信息采集（核心配置）
		// zap.AddCallerSkip(2), // 跳过当前封装层，获取真实的调用位置（关键）
	)

	return &Logger{logger}
}

func makeLogFunc(on bool, logger *Logger) func(data map[string]interface{}, msg ...interface{}) {
	return func(data map[string]interface{}, msg ...interface{}) {
		if !on {
			return
		}
		if len(msg) == 0 {
			msg = []interface{}{logger.Level().String()}
		}

		fields := make([]zap.Field, 0, len(data))
		for k, v := range data {
			fields = append(fields, zap.Any(k, v))
		}

		switch logger.Level() {
		case zap.InfoLevel:
			logger.Info(fmt.Sprint(msg[0]), fields...)
		case zap.WarnLevel:
			logger.Warn(fmt.Sprint(msg[0]), fields...)
		case zap.ErrorLevel:
			logger.Error(fmt.Sprint(msg[0]), fields...)
		case zap.DebugLevel:
			logger.Debug(fmt.Sprint(msg[0]), fields...)
		}
	}
}
