package kernal

import (
	"github.com/gin-gonic/gin"
	"github.com/gnehcaij/zeus/util"
)

type RequestContext struct {
	GinContext *gin.Context
}

func NewRequestContext(ctx *gin.Context) *RequestContext {
	return &RequestContext{
		GinContext: ctx,
	}
}

func (r *RequestContext) DefaultRequestParam(key string, defaultValue interface{}, valueType string) interface{} {
	v := r.GinContext.DefaultQuery(key, "")
	return util.StringConv(v, defaultValue, valueType)
}
