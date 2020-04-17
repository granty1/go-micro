package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/client/selector"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-plugins/client/http"
	"github.com/micro/go-plugins/registry/consul"
)

func main() {

	consulReg := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{"127.0.0.1:8500"}
	})
	s := selector.NewSelector(
			selector.Registry(consulReg),
			selector.SetStrategy(selector.RoundRobin),
	)

	callAPI(s)
}

func callAPI(selector selector.Selector)  {
	c := http.NewClient(
		client.Selector(selector),
		client.ContentType("application/json"),
	)
	req := c.NewRequest("product_service", "/v1/list", nil, )

	var resp map[string]interface{}
	if err := c.Call(context.Background(), req, &resp); err != nil {
		log.Error(err)
	}

	fmt.Println(resp)
}
