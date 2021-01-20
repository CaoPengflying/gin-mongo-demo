package clog

import (
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"path/filepath"
	"time"
)

var infoPriority = zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
	return true
})

var errPriority = zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
	return lvl >= zapcore.ErrorLevel
})

type splitSize struct {
	fileConfig *FileConfig
	encoder    zapcore.Encoder
}

func (splitSize *splitSize) getWriter(logLevel string) io.Writer {
	c := splitSize.fileConfig
	return &lumberjack.Logger{
		Filename:   c.LogFilePath + logLevel,
		MaxSize:    c.MaxSize,
		MaxBackups: c.MaxBackups,
		MaxAge:     c.MaxAge,
	}
}

func (splitSize *splitSize) getZapCores() []zapcore.Core {
	var zapCoreArray []zapcore.Core
	infoWriter := splitSize.getWriter("info.log")
	errorWriter := splitSize.getWriter("error.log")

	synInfo := zapcore.AddSync(infoWriter)
	synError := zapcore.AddSync(errorWriter)

	infoCore := zapcore.NewCore(splitSize.encoder, synInfo, infoPriority)
	errorCore := zapcore.NewCore(splitSize.encoder, synError, errPriority)

	zapCoreArray = append(zapCoreArray, infoCore)
	zapCoreArray = append(zapCoreArray, errorCore)

	return zapCoreArray
}

type splitTime struct {
	fileConfig *FileConfig
	encoder    zapcore.Encoder
}

func (splitTime *splitTime) getWriter(logLevel string) io.Writer {
	c := splitTime.fileConfig
	// 解决相对路径软连接错误问题
	fileName, err := filepath.Abs(c.LogFilePath)
	if err != nil {
		panic(err)
	}

	fileName = fileName + logLevel + ".log"
	hook, err := rotatelogs.New(
		fileName+"_%Y%m%d%H.log",
		rotatelogs.WithLinkName(fileName),
		rotatelogs.WithMaxAge(time.Hour*24*time.Duration(c.MaxAge)),
		rotatelogs.WithRotationTime(time.Hour*time.Duration(c.RotationTime)),
	)

	if err != nil {
		panic(err)
	}
	return hook
}

func (splitTime *splitTime) getZapCores() []zapcore.Core {
	var zapCoreArray []zapcore.Core
	infoWriter := splitTime.getWriter("info")
	errorWriter := splitTime.getWriter("error")

	synInfo := zapcore.AddSync(infoWriter)
	synError := zapcore.AddSync(errorWriter)

	infoCore := zapcore.NewCore(splitTime.encoder, synInfo, infoPriority)
	errorCore := zapcore.NewCore(splitTime.encoder, synError, errPriority)

	zapCoreArray = append(zapCoreArray, infoCore)
	zapCoreArray = append(zapCoreArray, errorCore)

	return zapCoreArray
}

func getSplitZapCores(fileConfig *FileConfig, encoder zapcore.Encoder) []zapcore.Core {
	if fileConfig.LogFilePath == "" {
		return nil
	}
	var retZapCores []zapcore.Core

	if fileConfig.RotationTime > 0 {
		splitTime := splitTime{fileConfig: fileConfig, encoder: encoder}
		retZapCores = splitTime.getZapCores()
	} else {
		splitSize := splitSize{fileConfig: fileConfig, encoder: encoder}
		retZapCores = splitSize.getZapCores()
	}

	return retZapCores
}
