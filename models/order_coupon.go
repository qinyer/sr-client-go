package models

// OrderCouponReq 添加/更新卡券信息
type OrderCouponReq struct {
	DataSourceId string         `json:"dataSourceId"` // 数据源id（create方法返回的data.dataSource.id）
	Coupons      *[]OrderCoupon `json:"coupons"`      // 卡券列表，数组最大长度 50
}

// OrderCoupon
type OrderCoupon struct {
	ExternalCouponId       string   `json:"external_coupon_id"`         // 卡券批次唯一 id。字段长度最小 1 字节，长度最大 100 字节
	CouponName             string   `json:"coupon_name"`                // 卡券名称，字段长度最小 1 字节，长度最大 100 字节
	PromotionIds           []string `json:"promotion_ids"`              // 活动 id 列表，字段长度最小 1 字节，长度最大 100 字节
	CouponCodes            []string `json:"coupon_codes"`               // 卡券编码
	CouponType             int      `json:"coupon_type"`                // 卡券类型；1：商家券；2：微信券
	CouponSubType          int      `json:"coupon_sub_type"`            // 卡券子类型；1：代金券；2：打折券
	PlanCount              int      `json:"plan_count"`                 // 券发行量
	StoreCount             int      `json:"store_count"`                // 券库存量
	StartTime              string   `json:"start_time"`                 // 券有效的开始时间，时间戳，毫秒
	EndTime                string   `json:"end_time"`                   // 券有效的结束时间，时间戳，毫秒
	RuleDescription        string   `json:"rule_description"`           // 券规则描述
	AmountCoupon           float64  `json:"amount_coupon"`              // 券金额，卡券子类型=1时必填
	DiscountCoupon         float64  `json:"discount_coupon"`            // 券折扣，卡券子类型=2时必填，大于等于 0，小于等于 1，至多 4 位小数
	AmountMinimum          float64  `json:"amount_minimum"`             // 使用卡券的最低消费金额
	ReleaseStatus          int      `json:"release_status"`             // 发放状态；1：有效；2；过期；3：冻结
	MaxCouponNumberPerUser int      `json:"max_coupon_number_per_user"` // 每位用户最大领取数量
	MaxAmountByDay         float64  `json:"max_amount_by_day"`          // 单天发放上限金额
	MaxCouponsByDay        int      `json:"max_coupons_by_day"`         // 单天发放上线个数
	NaturalPersonLimit     bool     `json:"natural_person_limit"`       // 是否开启自然人限制，true:开启；false:关闭
	Product                *Product `json:"product"`                    // 可使用卡券的商品列表
	UrlList                string   `json:"url_list"`                   // 领券链接
	ShowStartTime          string   `json:"show_start_time"`            // 券展示的开始时间，时间戳，毫秒
	ShowEndTime            string   `json:"show_end_time"`              // 券展示的结束时间，时间戳，毫秒
}

type Product struct {
	IsAll       bool     `json:"is_all"`       // 卡券使用范围是否为该小程序的全部商品，true:是；false：否
	CategoryIds []string `json:"category_ids"` // 可使用卡券的类目id列表
	SkuIds      []string `json:"sku_ids"`      // 可使用卡券的skuid列表
	SpuIds      []string `json:"spu_ids"`      // 可使用卡券的spuid列表
}
