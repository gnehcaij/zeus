package middleware

import (
	"github.com/gin-gonic/gin"
)

func SessionMiddleware() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		//c.Set(constant.LOGIDKEY, genLogId())
		//time.Sleep(3*time.Second)
		c.Next()
	}
}
