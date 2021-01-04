package test

import (
	"fmt"
	sr "github.com/qinyer/sr-client-go"
	"github.com/qinyer/sr-client-go/models"
	"testing"
)

func TestOrder(t *testing.T) {
	req := &models.OrderReq{
		DataSourceId: "11258",
		Orders: &[]models.Order{
			models.Order{
				ExternalOrderId:  "20201228112150001",
				CreateTime:       "1609125710601",
				OrderSource:      "wxapp",
				OrderType:        1,
				BrandId:          "B01",
				BrandName:        "打卡呀服饰小商品",
				GoodsNumTotal:    1,
				GoodsWeight:      300.00,
				GoodsAmountTotal: 99.00,
				FreightAmount:    6.00,
				OrderAmount:      105.00,
				PayableAmount:    95.10,
				PaymentAmount:    95.10,
				OrderStatus:      "1150",
				UserInfo: &models.UserInfo{
					OpenId:             "ok4Qb0U0dFqJ2KiSoGBXKCIqXnFY",
					AppId:              "wx63cb835d679b9188",
					UnionId:            "UNION01",
					UserPhone:          "13647623697",
					UserId:             "U001",
					MemberId:           "M0001",
					UserFirstOrderTime: "1609125710601",
				},
				GoodsInfo: &[]models.GoodsInfo{
					models.GoodsInfo{
						ExternalSkuId:  "SKU01",
						SkuNameChinese: "纯棉中腰直筒牛仔长裤",
						GoodsAmount:    99.00,
						PaymentAmount:  89.10,
						IsGift:         0,
						ExternalSpuId:  "SPU01",
						SpuNameChinese: "牛仔裤",
						SaleUnit:       "条",
						CategoryId:     "C02",
						CategoryName:   "服饰",
						GoodsNum:       1,
						GoodsWeight:    300.00,
						StoreInfo: &models.StoreInfo{
							ExternalStoreId: "9L11",
							StoreName:       "品牌名称（海岸城店）",
							StoreCity:       "深圳市",
						},
						ChanInfo: &[]models.ChanInfo{
							//models.ChanInfo{
							//	ChanWeapp: &models.ChanWeapp{ChanScene: "1140"},
							//	ChanCustom: &models.ChanCustom{
							//		ChanCustomId:       "666666",
							//		ChanCustomIdDesc:   "666666",
							//		ChanCustomCat3:     "666666",
							//		ChanCustomCat3Desc: "666666",
							//		ChanCustomCat2:     "666666",
							//		ChanCustomCat2Desc: "666666",
							//		ChanCustomCat1:     "666666",
							//		ChanCustomCatDesc:  "666666",
							//	},
							//	ChanId:         "8_5a23p4ded6ff6384",
							//	ChanReferAppId: "wx63cb835d679b9188",
							//},
						},
						CommissionInfo: &[]models.CommissionInfo{
							models.CommissionInfo{
								CommissionType: 1,
								CommissionFee:  10.00,
							},
						},
					},
				},
				PromotionInfo: &[]models.PromotionInfo{
					models.PromotionInfo{
						PromotionType:       1,
						ExternalPromotionId: "P01",
						PromotionName:       "双旦百亿补贴",
						PromotionAmount:     9.90,
					},
				},
				CouponInfo: &[]models.CouponInfo{
					models.CouponInfo{
						CouponType:        1,
						CouponAmountTotal: 9.90,
						CouponDetail: &[]models.CouponDetail{
							models.CouponDetail{
								ExternalCouponId: "T01",
								CouponBatchId:    "B01",
								CouponName:       "9折优惠",
								CouponAmount:     9.90,
							},
						},
					},
				},
				PaymentInfo: &[]models.PaymentInfo{
					models.PaymentInfo{
						PaymentType: "00009",
						TransId:     "TRANS001",
						TransAmount: 95.10,
					},
				},
				ExpressInfo: &models.ExpressInfo{
					LogisticsStatus:      "3010",
					GoodsTotalWeight:     300.00,
					ReceiverName:         "打卡呀助手",
					ReceiverPhone:        "13772727277",
					ReceiverAddress:      "重庆市腾讯大厦",
					ReceiverCountryCode:  "CN",
					ReceiverProvinceCode: "440000",
					ReceiverCityCode:     "440001",
					ReceiverDistrictCode: "440001",
					ExpectedDeliveryTime: "10:00-22:00",
					ExpectedDeliveryDate: "20201228",
					ExpressPackageInfo: &[]models.ExpressPackageInfo{
						models.ExpressPackageInfo{
							ExpressCompanyCode: "10003",
							ExpressCompanyName: "express-speed",
							ExpressCode:        "288383838",
							ShipTime:           "1572395131732",
							ExpressPage: &models.ExpressPage{
								MiniprogramPath:  "page/index/index",
								MiniprogramAppid: "wxc8f3e27b33ae9e8e",
								MiniprogramH5:    "https://haodongxi.com",
							},
							ExpressPackageInfo: &[]models.ExpressPackageInfoItem{
								models.ExpressPackageInfoItem{
									ExternalSkuId: "SKU01",
									Number:        1,
								},
							},
						},
					},
				},
				InvoiceInfo: &[]models.InvoiceInfo{
					models.InvoiceInfo{
						IfNeedInvoice:          false,
						InvoiceType:            "1002",
						InvoiceTitle:           "UUnw829auH",
						InvoiceContent:         "F371BYSWfe",
						InvoiceAdditionInfo:    "无",
						InvoiceCompany:         "重庆分公司",
						InvoiceTaxpayer:        "510107723420661",
						RegistryAddress:        "重庆市",
						RegistryPhone:          "13882882822",
						RegistryBankName:       "工商银行",
						RegistryBankAccount:    "62188282828282",
						InvoiceDeliveryAddress: "重庆市南岸区",
						InvoiceDeliveryName:    "QinY",
						InvoiceDeliveryPhone:   "13844884848",
						InvoiceNum:             "882828288",
					},
				},
				PointsTotal: 95.10,
				IsDeleted:   0,
			},
		},
	}
	res, err := sr.NewYouShuClient(cfg).Order().AddOrder(req)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(res)
}

func TestOrderSum(t *testing.T) {
	req := &models.OrderSumReq{
		DataSourceId: "11267",
		Orders: &[]models.OrderSum{
			models.OrderSum{
				RefDate:            "1609125710601",
				GiveOrderAmountSum: 95.10,
				GiveOrderNumSum:    1,
				PaymentAmountSum:   95.10,
				PayedNumSum:        1,
			},
		},
	}
	res, err := sr.NewYouShuClient(cfg).Order().AddOrderSum(req)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(res)
}

func TestOrderPromotion(t *testing.T) {
	req := &models.OrderPromotionReq{
		DataSourceId: "11265",
		Promotions: &[]models.OrderPromotion{
			models.OrderPromotion{
				ExternalPromotionId:  "P01",
				PromotionName:        "秒杀",
				PromotionDescription: "双旦百亿补贴",
				PromotionType:        10001,
				PromotionStatus:      1,
				Channels:             1,
				PromotionUrl:         "pages/promotion/promotion_id=895420584093285",
				StartTime:            "1609124845000",
				EndTime:              "1610680045000",
				PromotionSalesInfo: &[]models.PromotionSalesInfo{
					models.PromotionSalesInfo{
						ExternalStoreId: "9L11",
						ExternalSkuId:   "SKU01",
						PromotionPrice:  88.00,
					},
				},
			},
		},
	}
	res, err := sr.NewYouShuClient(cfg).Order().AddPromotion(req)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(res)
}

func TestOrderCoupon(t *testing.T) {
	req := &models.OrderCouponReq{
		DataSourceId: "11266",
		Coupons: &[]models.OrderCoupon{
			models.OrderCoupon{
				ExternalCouponId:       "T01",
				CouponName:             "9折优惠",
				PromotionIds:           []string{"P01"},
				CouponCodes:            []string{"CC01"},
				CouponType:             1,
				CouponSubType:          2,
				PlanCount:              100,
				StoreCount:             500,
				StartTime:              "1609125710601",
				EndTime:                "1610680045000",
				RuleDescription:        "打卡呀购买服饰类商品，享9折优惠",
				AmountCoupon:           0.9,
				DiscountCoupon:         0,
				AmountMinimum:          1,
				ReleaseStatus:          1,
				MaxCouponNumberPerUser: 1,
				MaxAmountByDay:         10,
				MaxCouponsByDay:        10,
				NaturalPersonLimit:     false,
				Product: &models.Product{
					IsAll:       false,
					CategoryIds: []string{"C02"},
					SkuIds:      []string{"SKU01"},
					SpuIds:      []string{"SPU01"},
				},
				UrlList:       "https://dakaya.qqmylife.cn",
				ShowStartTime: "1609125710601",
				ShowEndTime:   "1610680045000",
			},
		},
	}
	res, err := sr.NewYouShuClient(cfg).Order().AddCoupon(req)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(res)
}
