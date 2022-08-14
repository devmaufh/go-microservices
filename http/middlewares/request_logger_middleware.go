package middlewares

import (
	"github.com/devmaufh/go_books/domain/request_logs"
	"github.com/devmaufh/go_books/services"
	"github.com/devmaufh/go_books/utils/date_utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

func RequestLoggerStart() gin.HandlerFunc {
	requestLog := request_logs.RequestLog{}
	requestLog.Timestamp = date_utils.GetNowString()

	return func(context *gin.Context) {
		requestLog.ClientIp = context.ClientIP()
		context.Writer.Header().Set("X-Request-Time", date_utils.GetNowString())

		requestLog, err := services.CreateRequestLog(requestLog)
		if err != nil {
			context.AbortWithStatusJSON(500, err)
		}
		context.Writer.Header().Set("X-Request-ID", strconv.FormatInt(requestLog.Id, 10))
		context.Next()
	}
}
