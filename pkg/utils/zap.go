package utils

import (
	"io"
	"os"
	"time"

	"github.com/lestrrat/go-file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Option = zap.Option

type RotateOptions struct {
	MaxAge       time.Duration
	RotationTime time.Duration
}

func DefaultRotate() RotateOptions {
	return RotateOptions{
		MaxAge:       time.Hour * 24 * 30,
		RotationTime: time.Hour * 24,
	}
}

type TeeOption struct {
	Filename string
	Ropt     RotateOptions
	Lef      zap.LevelEnablerFunc
	Level    zapcore.Level
}

func NewTeeWithRotate(tops []TeeOption, opts ...Option) *zap.Logger {
	var cores []zapcore.Core
	cfg := zap.NewProductionEncoderConfig()
	cfg.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
	}

	for _, top := range tops {

		r, _ := rotatelogs.New(
			top.Filename+".%Y%m%d",
			rotatelogs.WithLinkName(top.Filename),
			rotatelogs.WithMaxAge(time.Hour*24*30),
			rotatelogs.WithRotationTime(time.Hour*24),
		)

		core := zapcore.NewCore(
			zapcore.NewJSONEncoder(cfg),
			zapcore.AddSync(zapcore.AddSync(r)),
			top.Lef,
		)
		cores = append(cores, core)
	}
	cores = append(cores, zapcore.NewCore(
		zapcore.NewJSONEncoder(cfg),
		zapcore.AddSync(os.Stdin),
		zap.InfoLevel,
	))
	return zap.New(zapcore.NewTee(cores...), opts...)
}

// New create a new logger (not support log rotating).
func New(writer io.Writer, level zapcore.Level, opts ...Option) *zap.Logger {
	if writer == nil {
		panic("the writer is nil")
	}
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(cfg.EncoderConfig),
		zapcore.AddSync(writer),
		level,
	)

	return zap.New(core, opts...)
}
