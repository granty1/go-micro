package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
	"net/http"
)

func main() {

	consulReg := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{"localhost:8500"}
	})

	engine := gin.Default()

	v1 := engine.Group("/v1")
	{
		v1.POST("/list", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"data": []string{
					"Grant",
					"Kral",
					"Yuhn",
				},
			})
		})
	}

	s := web.NewService(
		web.Address(":8002"),
		web.Name("user_service"),
		web.Version("0.1.1"),
		web.Handler(engine),
		web.Registry(consulReg),
	)
	//s.Init()
	s.Run()
}
