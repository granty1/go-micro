package protos

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/client/selector"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/registry"
)

const (
	Prod_Service = "prod_service"
)

type logWrapper struct {
	client.Client
}

var c client.Client

func (l *logWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	opts = append(opts, client.WithSelectOption(func(options *selector.SelectOptions) {
		// Add load balance strategy
		options.Strategy = selector.Random
	}), client.WithCallWrapper(func(callFunc client.CallFunc) client.CallFunc {
		return func(ctx context.Context, node *registry.Node, req client.Request, rsp interface{}, opts client.CallOptions) error {
			md, _ := metadata.FromContext(ctx)
			fmt.Printf("[Req]->  ctx: %v, address: %s, service: %s, method: %s\n", md, node.Address, req.Service(), req.Endpoint())
			return callFunc(ctx, node, req, rsp, opts)
		}
	}))
	return l.Client.Call(ctx, req, rsp, opts...)
}

func NewLogWrapper(cli client.Client) client.Client {
	return &logWrapper{
		cli,
	}
}

func init() {
	selfClient := micro.NewService(
		micro.Name("api_client"),
		micro.WrapClient(NewLogWrapper))
	c = selfClient.Client()
}

func RegisterRPCService(ctx *gin.Context, services ...string) {
	for _, v := range services {
		switch v {
		case Prod_Service:
			ctx.Keys[Prod_Service] = NewProdService(Prod_Service, c)
		}
	}

}
