package logoptionb

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	Log *zap.Logger
)

const (
	DebugLevel = zap.InfoLevel
	
)

func init() {
	// var logConfig zap.Config
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "/tmp/logs",
		MaxSize:    1, // megabytes
		MaxBackups: 0,
		MaxAge:     0, // days
	})
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:   "msg",
		LevelKey:     "level",
		TimeKey:      "time",
		EncodeTime:   zapcore.ISO8601TimeEncoder,
		EncodeLevel:  zapcore.LowercaseLevelEncoder,
		EncodeCaller: zapcore.ShortCallerEncoder,
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		w,
		zap.InfoLevel,
	)

	Log = zap.New(core)

	// logConfig = zap.Config{
	// 	OutputPaths: []string{"stdout", "/tmp/logs"},
	// 	Encoding:    "json",
	// 	Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
	// 	w,

	// 	EncoderConfig: zapcore.EncoderConfig{
	// 		MessageKey:   "msg",
	// 		LevelKey:     "level",
	// 		TimeKey:      "time",
	// 		EncodeTime:   zapcore.ISO8601TimeEncoder,
	// 		EncodeLevel:  zapcore.LowercaseLevelEncoder,
	// 		EncodeCaller: zapcore.ShortCallerEncoder,
	// 	},
	// }
	// var err error
	// Log, err = logConfig.Build()
	// if err != nil {
	// 	panic(err)
	// }
}

func Field(key string, value interface{}) zap.Field {
	return zap.Any(key, value)
}

func Info(msg string, tags ...zap.Field) {
	Log.Info(msg, tags...)
	Log.Sync()
}

func Error(msg string, tags ...zap.Field) {
	Log.Error(msg, tags...)
	Log.Sync()
}

func Debug(msg string, tags ...zap.Field) {
	Log.Debug(msg, tags...)
	Log.Sync()
}
