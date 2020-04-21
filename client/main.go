package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
	"log"
	"micro_demo/client/router"
)

func main() {
	consulReg := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{"localhost:8500"}
	})

	engine := gin.Default()

	server := web.NewService(
		web.Name("prod_client"),
		web.Address(":8080"),
		web.Handler(engine),
		web.Registry(consulReg),
	)


	router.Register(engine)

	if err := server.Init(); err != nil {
		log.Fatal(err)
	}
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
