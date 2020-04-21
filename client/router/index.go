package router

import (
	"github.com/gin-gonic/gin"
	"micro_demo/client/api"
	"micro_demo/client/middleware"
	protos "micro_demo/client/services"
)

func Register(engine *gin.Engine) {
	v1Group := engine.Group("v1")
	{
		prod := v1Group.Group("/prod").Use(middleware.RPCService(protos.Prod_Service))
		{
			prod.GET("/list",api.GetProdsList)
		}
	}
}
