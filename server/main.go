package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	"log"
	protos "micro_demo/server/services"
	"micro_demo/server/services/impl"
)

func main() {

	consulReg := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{"localhost:8500"}
	})

	service := micro.NewService(
		micro.Name("prod_service"),
		//micro.Address(":8001"),
		micro.Registry(consulReg),
	)

	if err := protos.RegisterProdServiceHandler(service.Server(), &impl.ProdService{}); err != nil {
		log.Fatal(err)
	}
	service.Init()
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
