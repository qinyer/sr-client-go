package models

// SalesInfoReq 添加/更新销售信息请求
type SalesInfoReq struct {
	DataSourceId string       `json:"dataSourceId"` // 数据源id（create方法返回的data.dataSource.id）
	SalesInfo    *[]SalesInfo `json:"salesinfo"`    // 商品库存价格列表 数组最大长度 50
}

// SalesInfo
type SalesInfo struct {
	ExternalSkuId   string          `json:"external_sku_id"`   // 您为商品 sku 分配的唯一 id。 色码款商品必填。 字段长度最小 1 字节，长度最大 128 字节
	ExternalStoreId string          `json:"external_store_id"` // 客户侧店铺/仓库 id 字段长度最小 1 字节，长度最大 32 字节
	Price           *Price          `json:"price,omitempty"`             // 商品价格信息
	Stock           *Stock          `json:"stock,omitempty"`             // 库存信息
	TargetUrlProps  *TargetUrlProps `json:"target_url_props"`  // 商品目标页属性
}

// SkuPromotion
type SkuPromotion struct {
	ExternalPromotionId string  `json:"external_promotion_id"` // 活动 id
	PromotionPrice      float64 `json:"promotion_price"`       // 活动价格
}

// Price
type Price struct {
	CurrentPrice  float64         `json:"current_price"`  // 商品目前售价，单位：元，保留2位小数；大于等于0，最小值0；current_price需≤sku_price
	DailyPrice    float64         `json:"daily_price,omitempty"`    // 商品日常售价，单位：元，保留两位小数
	SkuPrice      float64         `json:"sku_price"`      // 商品原价，单位：元，保留2位小数；大于等于0，最小值0；current_price需≤sku_price
	SkuPromotions *[]SkuPromotion `json:"sku_promotions,omitempty"` // 商品活动及促销价列表 数组最大长度 50
}

// Stock
type Stock struct {
	SkuStockStatus int `json:"sku_stock_status"` // 同一个商品在不同店仓中的销售/上下架状态 1：销售中；2：停售中
}

// TargetUrlProps
type TargetUrlProps struct {
	MiniprogramAppid string `json:"miniprogram_appid"` // 微信小程序appid，字段长度最小 1 字节，长度最大 1024 字节
	UrlMiniprogram   string `json:"url_miniprogram"`   // 小程序落地页 url，长度不超过2048 字段长度最小 1 字节，长度最大 2048 字节
}
