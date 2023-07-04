package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func RequestLogging() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		startTime := time.Now().UnixNano()
		ctx.Next()
		latency := time.Now().UnixNano() - startTime
		log.Info().
			Str("method", ctx.Request.Method).
			Str("endpoint", ctx.Request.RequestURI).
			Int("code", ctx.Writer.Status()).
			Int64("latency", latency).Send()
	}
}
