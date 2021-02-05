package models

// OrderPromotionReq 添加/更新活动信息
type OrderPromotionReq struct {
	DataSourceId string            `json:"dataSourceId"` // 数据源id（create方法返回的data.dataSource.id）
	Promotions   *[]OrderPromotion `json:"promotions"`   // 订单列表 数组最大长度 50
}

//OrderPromotion
type OrderPromotion struct {
	ExternalPromotionId  string                `json:"external_promotion_id"`           // 您为促销活动分配的唯一 id。一般而言，是您在商品库为该促销活动分配的 id。字段长度最小 1 字节，长度最大 100 字节
	PromotionName        string                `json:"promotion_name"`                  // 促销名称字段长度最小 1 字节，长度最大 100 字节
	PromotionDescription string                `json:"promotion_description,omitempty"` //	促销描述字段长度最大 100 字节
	PromotionType        int                   `json:"promotion_type"`                  // 促销类型描述字段长度最小 1 字节，长度最大 100 字节，10001：秒杀
	PromotionStatus      int                   `json:"promotion_status"`                // 活动状态1：启用；2：停用
	Channels             int                   `json:"channels"`                        // 活动渠道列表字段长度最小 1 字节，1：商家小程序；2：商家app；3：商家H5；4：商家pcweb；5：线下人工pos；6：线下自助收银；99：其他
	PromotionUrl         string                `json:"promotion_url,omitempty"`         // 促销活动 url字段长度最大 100 字节
	StartTime            string                `json:"start_time"`                      // 活动开始时间，时间戳，毫秒
	EndTime              string                `json:"end_time"`                        // 活动结束时间，时间戳，毫秒
	PromotionSalesInfo   *[]PromotionSalesInfo `json:"promotion_salesinfo"`             // 参与店仓的列表
}

// PromotionSalesInfo
type PromotionSalesInfo struct {
	ExternalStoreId string  `json:"external_store_id"` // 店仓分配的唯一ID
	ExternalSkuId   string  `json:"external_sku_id"`   // 商品SKU分配的唯一ID
	PromotionPrice  float64 `json:"promotion_price"`   // 活动的促销价格
}
