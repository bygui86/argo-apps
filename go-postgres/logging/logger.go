package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	defaultEncoding = "console"
	defaultLevel    = zapcore.InfoLevel
)

var Config *zap.Config
var Log *zap.Logger
var SugaredLog *zap.SugaredLogger

func init() {

	Config = &zap.Config{
		Encoding:         defaultEncoding,
		Level:            zap.NewAtomicLevelAt(defaultLevel),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:  "message",
			TimeKey:     "time",
			EncodeTime:  zapcore.ISO8601TimeEncoder,
			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,
		},
	}
	Log, _ = Config.Build()
	SugaredLog = Log.Sugar()
}
