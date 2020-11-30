package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		return fmt.Sprintf("%s %s %s - [%s] %d \n",
			params.ClientIP,
			params.Latency,
			params.Method,
			params.TimeStamp.Format(time.RFC822),
			params.StatusCode)
	})
}
