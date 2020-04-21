package api

import (
	"github.com/gin-gonic/gin"
	protos "micro_demo/client/services"
	"net/http"
)

func GetProdsList(ctx *gin.Context) {
	prodService, ok := ctx.Keys[protos.Prod_Service].(protos.ProdService)
	if !ok {
		ctx.JSON(http.StatusBadGateway, "cannot get prod service")
	} 

	resp, err := prodService.GetProdsList(ctx, &protos.ProdsRequest{
		Size: 3,
	})
	if err != nil {
		ctx.JSON(http.StatusOK, err)
		return
	}
	ctx.JSON(http.StatusOK, resp.Data)
}
