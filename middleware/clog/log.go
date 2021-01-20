package clog

import (
	"context"
)

type CLogger struct {
	h *Handlers
}

type Config struct {
	FileConfig *FileConfig `yaml:"fileConfig"`
}

var (
	cLogger *CLogger
)

func Init(c *Config, serviceName string) {
	if c == nil {
		fileConfig := &FileConfig{
			Console: true,
		}
		c = &Config{FileConfig: fileConfig}
	}
	fileHandler := NewFileHandler(c.FileConfig, serviceName)
	cLogger = &CLogger{}
	cLogger.h = NewHandlers(fileHandler)
}

func NewLogger(c *Config, serviceName string) *CLogger {
	var cLogger = &CLogger{}
	if c == nil {
		fileConfig := &FileConfig{
			Console: true,
		}
		c = &Config{FileConfig: fileConfig}
	}
	fileHandler := NewFileHandler(c.FileConfig, serviceName)
	cLogger.SetHandlers(NewHandlers(fileHandler))
	return cLogger
}

func (cLogger *CLogger) SetHandlers(hs *Handlers) {
	cLogger.h = hs
}

func GetDefaultLogger() *CLogger {
	return cLogger
}

func DebugC(ctx context.Context, format string, args ...interface{}) {
	cLogger.h.Log(ctx, DebugLevel, format, args...)
}

func InfoC(ctx context.Context, format string, args ...interface{}) {
	cLogger.h.Log(ctx, InfoLevel, format, args...)
}

func WarnC(ctx context.Context, format string, args ...interface{}) {
	cLogger.h.Log(ctx, WarnLevel, format, args...)
}

func ErrorC(ctx context.Context, format string, args ...interface{}) {
	cLogger.h.Log(ctx, ErrorLevel, format, args...)
}

func PanicC(ctx context.Context, format string, args ...interface{}) {
	cLogger.h.Log(ctx, PanicLevel, format, args...)
}

func FatalC(ctx context.Context, format string, args ...interface{}) {
	cLogger.h.Log(ctx, FatalLevel, format, args...)
}

func DebugW(msg string, m map[string]interface{}) {
	cLogger.h.LogWith(context.Background(), DebugLevel, msg, m)
}

func InfoW(msg string, m map[string]interface{}) {
	cLogger.h.LogWith(context.Background(), InfoLevel, msg, m)
}

func WarnW(msg string, m map[string]interface{}) {
	cLogger.h.LogWith(context.Background(), WarnLevel, msg, m)
}

func ErrorW(msg string, m map[string]interface{}) {
	cLogger.h.LogWith(context.Background(), ErrorLevel, msg, m)
}

func FatalW(msg string, m map[string]interface{}) {
	cLogger.h.LogWith(context.Background(), FatalLevel, msg, m)
}

func PanicW(msg string, m map[string]interface{}) {
	cLogger.h.LogWith(context.Background(), PanicLevel, msg, m)
}

func Debug(format string, args ...interface{}) {
	cLogger.h.Log(context.Background(), DebugLevel, format, args...)
}

func Info(format string, args ...interface{}) {
	cLogger.h.Log(context.Background(), InfoLevel, format, args...)
}

func Warn(format string, args ...interface{}) {
	cLogger.h.Log(context.Background(), WarnLevel, format, args...)
}

func Error(format string, args ...interface{}) {
	cLogger.h.Log(context.Background(), ErrorLevel, format, args...)
}

func Panic(format string, args ...interface{}) {
	cLogger.h.Log(context.Background(), PanicLevel, format, args...)
}

func Fatal(format string, args ...interface{}) {
	cLogger.h.Log(context.Background(), FatalLevel, format, args...)
}

func (cLogger *CLogger) Debug(format string, args ...interface{}) {
	cLogger.h.Log(context.Background(), DebugLevel, format, args...)
}

func (cLogger *CLogger) Info(format string, args ...interface{}) {
	cLogger.h.Log(context.Background(), InfoLevel, format, args...)
}

func (cLogger *CLogger) Warn(format string, args ...interface{}) {
	cLogger.h.Log(context.Background(), WarnLevel, format, args...)
}

func (cLogger *CLogger) Error(format string, args ...interface{}) {
	cLogger.h.Log(context.Background(), ErrorLevel, format, args...)
}

func (cLogger *CLogger) Panic(format string, args ...interface{}) {
	cLogger.h.Log(context.Background(), PanicLevel, format, args...)
}

func (cLogger *CLogger) Fatal(format string, args ...interface{}) {
	cLogger.h.Log(context.Background(), FatalLevel, format, args...)
}

func (cLogger *CLogger) DebugW(msg string, m map[string]interface{}) {
	cLogger.h.LogWith(context.Background(), DebugLevel, msg, m)
}

func (cLogger *CLogger) InfoW(msg string, m map[string]interface{}) {
	cLogger.h.LogWith(context.Background(), InfoLevel, msg, m)
}

func (cLogger *CLogger) WarnW(msg string, m map[string]interface{}) {
	cLogger.h.LogWith(context.Background(), WarnLevel, msg, m)
}

func (cLogger *CLogger) ErrorW(msg string, m map[string]interface{}) {
	cLogger.h.LogWith(context.Background(), ErrorLevel, msg, m)
}

func (cLogger *CLogger) FatalW(msg string, m map[string]interface{}) {
	cLogger.h.LogWith(context.Background(), FatalLevel, msg, m)
}

func (cLogger *CLogger) PanicW(msg string, m map[string]interface{}) {
	cLogger.h.LogWith(context.Background(), PanicLevel, msg, m)
}

// DebugC with context logs a message.
func (cLogger *CLogger) DebugC(ctx context.Context, format string, args ...interface{}) {
	cLogger.h.Log(ctx, DebugLevel, format, args...)
}

// InfoC with context logs a message.
func (cLogger *CLogger) InfoC(ctx context.Context, format string, args ...interface{}) {
	cLogger.h.Log(ctx, InfoLevel, format, args...)
}

// WarnC with context logs a message.
func (cLogger *CLogger) WarnC(ctx context.Context, format string, args ...interface{}) {
	cLogger.h.Log(ctx, WarnLevel, format, args...)
}

// ErrorC with context logs a message.
func (cLogger *CLogger) ErrorC(ctx context.Context, format string, args ...interface{}) {
	cLogger.h.Log(ctx, ErrorLevel, format, args...)
}

// PanicC with context logs a message.
func (cLogger *CLogger) PanicC(ctx context.Context, format string, args ...interface{}) {
	cLogger.h.Log(ctx, PanicLevel, format, args...)
}

// FatalC with context logs a message.
func (cLogger *CLogger) FatalC(ctx context.Context, format string, args ...interface{}) {
	cLogger.h.Log(ctx, FatalLevel, format, args...)
}
