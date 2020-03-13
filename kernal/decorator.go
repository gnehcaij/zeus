package kernal

import (
	"github.com/gin-gonic/gin"
	"github.com/gnehcaij/zeus/constant"
	"net/http"
)

func CommonAPI(logicHandler func(ctx *RequestContext) ([]byte, *constant.CommonStatus)) func(ctx *gin.Context) {

	handler := func(c *gin.Context) {
		ctx := NewRequestContext(c)
		resp, _ := logicHandler(ctx)

		c.Data(http.StatusOK, "application/json; charset=utf-8", resp)
	}
	return handler
}
