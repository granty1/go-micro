package protos

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"micro_demo/client/wrappers"
)

const (
	Prod_Service = "prod_service"
)

var c client.Client

func init() {
	selfClient := micro.NewService(
		micro.Name("api_client"),
		micro.WrapClient(wrappers.NewLogWrapper))
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
