package trace

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	TraceIDHeaderKey = "trace-id"
)

//TraceLogger 日志打印抽象接口
type TraceLogger interface {
	InfoW(msg string, m map[string]interface{})
}

func Trace() gin.HandlerFunc {
	return func(context *gin.Context) {
		traceId := context.Request.Header.Get(TraceIDHeaderKey)
		context.Set(TraceIDHeaderKey, traceId)

		if traceId == "" {
			newTraceId, _ := uuid.NewRandom()
			context.Request.Header.Set(TraceIDHeaderKey, newTraceId.String())
			context.Set(TraceIDHeaderKey, newTraceId.String())
		}
	}
}

