package models

import "sr-client-go/utils"

// 数据源类型对照表
// 每个接口上报数据前都需都添加一个数据仓库，每种数据源类型的数据仓库只可添加一次，该数据仓库在传输数据时可重复使用。
const (
	DsOrderMain                 = 0  // 主订单 订单数据仓库, 对应上报接口 /order/add_order
	DsOrderSub                  = 1  // 子订单 子订单数据仓库，对应上报接口 /order/add_sub_order
	DsSaleInfo                  = 2  // 销售信息 销售信息数据仓库，对应上报接口 /salesinfo/add
	DsSku                       = 3  // 商品sku 商品sku数据仓库，对应上报接口 /sku/add
	DsStore                     = 4  // 门店信息 门店信息数据仓库，对应上报接口 /store/add
	DsSpu                       = 5  // 商品spu 商品spu数据仓库，对应上报接口 /spus/add
	DsProductCategories         = 6  // 商品类目 商品类目数据仓库，对应上报接口 /product_categories/add
	DsAnalysisVisitPage         = 7  // 微信小程序页面访问 微信小程序页面访问信息数据仓库，对应微信接口 analysis.getVisitPage，对应上报接口 /analysis/add_wxapp_visit_page
	DsAnalysisVisitDistribution = 8  // 微信小程序访问分布数据 微信小程序页面访问信息数据仓库，对应微信接口 analysis.getVisitDistribution，对应上报接口 /analysis/add_wxapp_visit_distribution
	DsUser                      = 11 // 用户信息 用户信息数据仓库，对应上报接口 /user/add_user
	DsOrderPromotion            = 12 // 活动信息 活动信息数据仓库，对应上报接口 /order/add_promotion
	DsOrderCoupon               = 13 // 卡券信息数据仓库，对于上报接口 /order/add_coupon
	DsAnalysisVisitTrend        = 16 // 微信小程序日访问趋势数据仓库，对应微信接口 analysis.getDailyVisitTrend，对应上报接口 /analysis/add_wxapp_daily_visit_trend
	DsOrderSum                  = 51 // 汇总订单数据仓库, 对应上报接口 /order/add_order_sum
)

// DataSourceIdAndFunc 数据源类型和接口方法对照表，假定一对多关系，便于后期新增接口方法归集到指定数据源类型
var DataSourceIdAndFunc = map[int][]string{
	DsOrderMain:                 {"/order/add_order"},
	DsOrderSub:                  {"/order/add_sub_order"},
	DsSaleInfo:                  {"/salesinfo/add"},
	DsSku:                       {"/sku/add"},
	DsStore:                     {"/store/add"},
	DsSpu:                       {"/spus/add"},
	DsProductCategories:         {"/product_categories/add"},
	DsAnalysisVisitPage:         {"/analysis/add_wxapp_visit_page"},
	DsAnalysisVisitDistribution: {"/analysis/add_wxapp_visit_distribution"},
	DsUser:                      {"/user/add_user"},
	DsOrderPromotion:            {"/order/add_promotion"},
	DsOrderCoupon:               {"/order/add_coupon"},
	DsAnalysisVisitTrend:        {"/analysis/add_wxapp_daily_visit_trend"},
	DsOrderSum:                  {"/order/add_order_sum"},
}

// DataSourceRes 添加数据仓库结果
type DataSourceRes struct {
	utils.CommonError

	Data struct {
		DataSource DataSource `json:"dataSource"`
	} `json:"data"`
}

// DataSourceGetRes 获取数据仓库结果
type DataSourceGetRes struct {
	utils.CommonError

	Data struct {
		DataSources []DataSource `json:"dataSources"`
	} `json:"data"`
}

// 数据仓库对象
type DataSource struct {
	Id         string `json:"id"`
	Type       int    `json:"type"`
	MerchantId string `json:"merchantId"`
}
