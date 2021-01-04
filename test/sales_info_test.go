package test

import (
	"fmt"
	sr "sr-client-go"
	"sr-client-go/models"
	"testing"
)

func TestSaleInfo(t *testing.T)  {
	req := &models.SalesInfoReq{
		DataSourceId: "11264",
		SalesInfo:    &[]models.SalesInfo{
			models.SalesInfo{
				ExternalSkuId:   "SKU01",
				ExternalStoreId: "9L11",
				Price:           &models.Price{
					CurrentPrice:  99.00,
					DailyPrice:    99.00,
					SkuPrice:      199.00,
					SkuPromotions: &[]models.SkuPromotion{
						models.SkuPromotion{
							ExternalPromotionId: "P01",
							PromotionPrice:      99.00,
						},
					},
				},
				Stock:           &models.Stock{SkuStockStatus: 2},
				TargetUrlProps:  &models.TargetUrlProps{
					MiniprogramAppid: "wx63cb835d679b9188",
					UrlMiniprogram:   "pages/magazine/index?sku_id=SKU01&from=wechat&to=productDetail",
				},
			},
		},
	}
	res, err := sr.NewYouShuClient(cfg).SalesInfo().Add(req)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(res)
}