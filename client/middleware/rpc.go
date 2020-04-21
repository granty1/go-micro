package middleware

import (
	"github.com/gin-gonic/gin"
	protos "micro_demo/client/services"
)

func RPCService(services ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Keys = make(map[string]interface{})
		protos.RegisterRPCService(ctx,services...)
		ctx.Next()
	}
}
