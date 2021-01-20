package clog

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type FileHandler struct {
	logger        *zap.SugaredLogger
	level         Level
	appId         string //服务名
	filePathDepth int
}

// FileConfig 日志配置，如果配置了RotationTime，则优先按时间做日志切割
type FileConfig struct {
	LogFilePath   string `yaml:"logFilePath"`   // 日志路径
	MaxSize       int    `yaml:"maxSize"`       // 单个日志最大的文件大小. 单位: MB
	MaxBackups    int    `yaml:"maxBackups"`    // 日志文件最多保存多少个备份
	MaxAge        int    `yaml:"maxAge"`        // 文件最多保存多少天
	Console       bool   `yaml:"console"`       // 是否命令行输出，开发环境可以使用
	LevelString   string `yaml:"levelString"`   // 输出的日志级别, 值：debug,info,warn,error,panic,fatal
	RotationTime  int    `yaml:"rotationTime"`  // 按时间间隔分割日志，单位为小时
	FilePathDepth int    `yaml:"filePathDepth"` // 日志中path字段记录的文件目录的层级深度
}

func NewFileHandler(c *FileConfig, serviceName string) *FileHandler {
	xLogTimeEncoder := func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format(time.RFC3339))
	}

	// Optimize the xLog output for machine consumption and the console output
	// for human operators.
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.MessageKey = ""
	encoderConfig.TimeKey = "time"
	//encoderConfig.CallerKey = "path" // 原定的path字段含义太多，建议还是分开，然后log调用的地方就叫caller
	encoderConfig.EncodeTime = xLogTimeEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	encoderConfig.EncodeName = zapcore.FullNameEncoder

	xLogEncoder := zapcore.NewJSONEncoder(encoderConfig)
	// Join the outputs, encoders, and level-handling functions into
	// zapcore.Cores, then tee the cores together.
	var allCore = getSplitZapCores(c, xLogEncoder)

	if c.Console {
		consoleDebugging := zapcore.Lock(os.Stdout)
		allCore = append(allCore, zapcore.NewCore(xLogEncoder, consoleDebugging, infoPriority))
	}

	core := zapcore.NewTee(allCore...)

	var opts []zap.Option

	logger := zap.New(core).WithOptions(opts...).Sugar()
	defer logger.Sync()

	appID := ""
	if serviceName != "" {
		appID = serviceName
	}

	return &FileHandler{logger: logger, level: LevelStringToCode(c.LevelString), appId: appID, filePathDepth: c.FilePathDepth}
}

func log(logger *zap.SugaredLogger, l Level, keysAndValues []interface{}) {
	switch l {
	case DebugLevel:
		logger.Debugw("", keysAndValues...)
	case InfoLevel:
		logger.Infow("", keysAndValues...)
	case WarnLevel:
		logger.Warnw("", keysAndValues...)
	case ErrorLevel:
		logger.Errorw("", keysAndValues...)
	case PanicLevel:
		logger.Panicw("", keysAndValues...)
	case FatalLevel:
		logger.Fatalw("", keysAndValues...)
	}
}

func (fh *FileHandler) Log(ctx context.Context, l Level, format string, args ...interface{}) {
	if l < fh.level {
		return
	}

	logger := fh.getLogger()

	traceID := ctx.Value(TraceIDHeaderKey)
	logger = logger.With(AppIdKey, fh.appId, TraceIDKey, traceID)

	msg := format
	if msg == "" && len(args) > 0 {
		msg = fmt.Sprint(args...)
	} else if msg != "" && len(args) > 0 {
		msg = fmt.Sprintf(format, args...)
	}

	var keysAndValues []interface{}

	keysAndValues = append(keysAndValues, "msg")
	keysAndValues = append(keysAndValues, msg)

	log(logger, l, keysAndValues)
}

func (fh *FileHandler) LogWith(ctx context.Context, l Level, msg string, m map[string]interface{}) {
	if l < fh.level {
		return
	}

	logger := fh.getLogger()

	if msg != "" {
		m["msg"] = msg
	}
	if fh.appId != "" {
		m[AppIdKey] = fh.appId
	}

	var keysAndValues []interface{}

	for k, v := range m {
		keysAndValues = append(keysAndValues, k)
		keysAndValues = append(keysAndValues, v)
	}

	log(logger, l, keysAndValues)
}

func (fh *FileHandler) Close() (err error) {
	return
}

func getCaller(skip int, filePathDepth int) string {
	fileName, line, funcName := "???", 0, "???"
	pc, fileName, line, ok := runtime.Caller(skip)
	if ok {
		funcName = runtime.FuncForPC(pc).Name()
		funcName = filepath.Base(funcName)

		fileName = getFileName(fileName, filePathDepth)
	}

	ca := fileName + ":" + strconv.Itoa(line) + "(" + funcName + ")"
	return ca
}

func getFileName(fileName string, filePathDepth int) string {
	if filePathDepth <= 1 {
		return filepath.Base(fileName)
	}

	newFileName := strings.Trim(fileName, string(os.PathSeparator))
	paths := strings.Split(newFileName, string(os.PathSeparator))

	if len(paths) > filePathDepth {
		newFileName = filepath.Join(paths[len(paths)-filePathDepth:]...)
	}

	return newFileName
}

func (fh *FileHandler) getLogger() (logger *zap.SugaredLogger) {
	logger = fh.logger.With("path", getCaller(5, fh.filePathDepth))
	return
}
