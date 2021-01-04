package test

import (
	"fmt"
	sr "sr-client-go"
	"sr-client-go/models"
	"testing"
)

func TestSku(t *testing.T) {
	req := &models.SkuReq{
		DataSourceId: "11261",
		Skus: &[]models.Sku{
			models.Sku{
				ExternalSkuId: "SKU01",
				ExternalSpuId: "SPU01",
				SkuBarcode:    "233233233233",
				ImgUrls: &[]models.ImgUrl{
					models.ImgUrl{
						PrimaryImgs: &[]models.Img{
							models.Img{ImgUrl: "https://bkimg.cdn.bcebos.com/pic/b21c8701a18b87d6bad5f50d0a0828381f30fd09?x-bce-process=image/resize,m_lfit,w_268,limit_1/format,f_jpg"},
						},
						Imgs: &[]models.Img{
							models.Img{ImgUrl: "https://bkimg.cdn.bcebos.com/pic/b21c8701a18b87d6bad5f50d0a0828381f30fd09?x-bce-process=image/resize,m_lfit,w_268,limit_1/format,f_jpg"},
						},
						DetailImgs: &[]models.Img{
							models.Img{ImgUrl: "https://bkimg.cdn.bcebos.com/pic/b21c8701a18b87d6bad5f50d0a0828381f30fd09?x-bce-process=image/resize,m_lfit,w_268,limit_1/format,f_jpg"},
						},
					},
				},
				ProductProps: &models.ProductProps{
					Color: &models.Color{
						ColorRgb:  "#5E96D7",
						ColorName: "天蓝",
					},
					IsO2o: false,
					Size:  "M",
				},
				CategoryProps: &[]models.CategoryProp{
					models.CategoryProp{
						ExternalCategoryIdLeaf: "C02",
						CategoryType:           2,
					},
				},
				SalesProps:          &models.SalesProps{IsAvailable: false},
				DescProps:           &models.DescProps{ProductNameChinese: "纯棉中腰直筒牛仔长裤"},
				ExternalCreatedTime: "1609124845481",
			},
		},
	}
	res, err := sr.NewYouShuClient(cfg).Sku().Add(req)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(res)
}
