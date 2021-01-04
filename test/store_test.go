package test

import (
	"fmt"
	sr "github.com/qinyer/sr-client-go"
	"github.com/qinyer/sr-client-go/models"
	"testing"
)

func TestStore(t *testing.T)  {
	req := &models.StoreReq{
		DataSourceId: "11259",
		Stores:       &[]models.Store{
			models.Store{
				ExternalStoreId:       "9L11",
				Type:                  1,
				BusinessType:          99,
				OperationStatus:       1,
				DaojiaOperationStatus: 1,
				PhoneNumbers:          []string{"02362781081"},
				LocationInfo:          &models.LocationInfo{
					CountryCode:  "CN",
					CountryName:  "中国",
					ProvinceCode: 440000,
					ProvinceName: "广东省",
					CityCode:     440300,
					CityName:     "深圳市",
					DistrictCode: 440305,
					DistrictName: "南山区",
					Address:      "文心五路保利文化广场B区二楼",
					GeoInfo:      &[]models.GeoInfo{
						models.GeoInfo{
							Latitude:  45.73622,
							Longitude: 126.54209,
							Type:      3,
						},
					},
				},
				BasicProps:            &models.BasicProps{
					Name:          "品牌名称（海岸城店）",
					OpeningTime:   "1515118117111",
					OperatingTime: &[]models.OperatingTime{
						models.OperatingTime{
							DateZone: "周一至周日",
							TimeZone: "10:00-22:00",
						},
					},
				},
				DeliveryInfo:          &models.DeliveryInfo{
					RangeType: 3,
					Radius:    "300km",
					GeoGroup:  &[]models.GeoGroup{
						models.GeoGroup{Geos: &[]models.GeoInfo{
							models.GeoInfo{
								Latitude:  26.075368,
								Longitude: 119.329487,
								Type:      1,
							},
						}},
					},
				},
			},
		},
	}
	res, err := sr.NewYouShuClient(cfg).Store().Add(req)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(res)
}
