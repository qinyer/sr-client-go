package test

import (
	"fmt"
	sr "github.com/qinyer/sr-client-go"
	"github.com/qinyer/sr-client-go/models"
	"testing"
)

func TestUser(t *testing.T) {
	req := &models.UserReq{
		DataSourceId: "11260",
		Users: &[]models.User{
			models.User{
				UserId:      "U001",
				PhoneNumber: "a486828e0524c7f10f4ad56785b64b93",
				UserSpec: &[]models.UserSpec{
					models.UserSpec{
						AppType:         1,
						Appid:           "wx63cb835d679b9188",
						Openid:          "ok4Qb0U0dFqJ2KiSoGBXKCIqXnFY",
						UserCreatedTime: "1609125710601",
						UserTouchTime:   "1609125710601",
					},
				},
				CardSpec: &[]models.CardSpec{
					models.CardSpec{
						CardId:             "CARD001",
						CardStatus:         1,
						CardType:           1,
						CardMasterId:       "MASTER001",
						CardCreatedTime:    "1609125710601",
						CardLevel:          1,
						CardLevelStartTime: "1609125710601",
						CardLevelEndTime:   "1609125710601",
						StoreId:            "9L11",
						Point:              99.00,
					},
				},
				DeviceSpec: &[]models.DeviceSpec{
					models.DeviceSpec{
						DeviceType: 1,
						Idfa:       "IDFA001",
						Uuid:       "UUID001",
						Imei:       "imei",
						Mac:        "40:b0:76:81:70:86",
					},
				},
				BasicSpec: &models.BasicSpec{
					Name:      "小秦",
					Nickname:  "tabqin",
					HeaderUrl: "www.avatar.com/xiaoq",
					Gender:    1,
					Age:       26,
					Qq:        13773737,
					IdCardSpec: &[]models.IdCardSpec{
						models.IdCardSpec{
							IdCardType: 1,
							IdCardNo:   "320106199002228751",
						},
					},
					Birthday:     "1993-12-20",
					Province:     "广东省",
					ProvinceCode: 440000,
					City:         "深圳市",
					CityCode:     440300,
					District:     "南山区",
					DistrictCode: 440305,
					Address:      "文心五路保利文化广场B区二楼",
					Email:        []string{"12828828@qq.com"},
				},
				InvoiceSpec: &[]models.InvoiceSpec{
					models.InvoiceSpec{
						InvoiceType:     1,
						InvoiceTitle:    "发票抬头",
						InvoiceCompany:  "公司名称",
						InvoiceTaxpayer: "纳税人识别号",
					},
				},
			},
		},
	}
	res, err := sr.NewYouShuClient(cfg).User().Add(req)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(res)
}
