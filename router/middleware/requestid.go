package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestId := c.Request.Header.Get("X-Request-Id")
		if requestId == "" {
			uuid := uuid.New()
			requestId = uuid.String()
		}

		c.Set("X-Request-Id", requestId)
		c.Writer.Header().Set("X-RequestId-Id", requestId)
		c.Next()
	}
}
