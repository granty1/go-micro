package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
	"micro_demo/models"
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
			var req models.ProdsRequest
			if err := context.ShouldBindJSON(&req); err != nil {
				fmt.Println(err)
			}

			var data []*models.ProdModel
			for i := 0 ; i < int(req.Size) ; i ++ {
				data = append(data, &models.ProdModel{
					ProdId:               int32(i) + 1,
					ProdName:             "Product",
				})
			}

			context.JSON(http.StatusOK, models.ProdsResponse{
				Data:                 data,
			})
		})
	}

	s := web.NewService(
		web.Address(":8002"),
		web.Name("product_service"),
		web.Version("0.1.1"),
		web.Handler(engine),
		web.Registry(consulReg),
	)
	//s.Init()
	s.Run()
}
