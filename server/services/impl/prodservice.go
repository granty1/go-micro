package impl

import (
	context "context"
	"errors"
	"micro_demo/server/services"
)

type ProdService struct{}

func (p *ProdService) GetProdsList(ctx context.Context, request *protos.ProdsRequest, response *protos.ProdsResponse) error {
	if request.Size == 0 {
		return errors.New("size cannot be zero")
	}
	var data []*protos.ProdModel
	for i := 0; i < int(request.Size); i++ {
		data = append(data, &protos.ProdModel{
			ProdId:   int32(i + 1),
			ProdName: "name",
		})
	}
	response.Data = data
	return nil
}
