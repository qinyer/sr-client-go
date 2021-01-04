package test

import (
	"fmt"
	sr "github.com/qinyer/sr-client-go"
	"github.com/qinyer/sr-client-go/models"
	"testing"
)

func TestSpu(t *testing.T)  {
	req := &models.SpuReq{
		DataSourceId: "11262",
		Spus:         &[]models.Spu{
			models.Spu{
				ExternalSpuId: "SPU01",
				DescProps:     &models.DescProps{ProductNameChinese: "牛仔裤"},
			},
		},
	}
	res, err := sr.NewYouShuClient(cfg).Spu().Add(req)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(res)
}