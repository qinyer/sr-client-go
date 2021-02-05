package models

// OrderReq 添加/更新订单请求
type OrderReq struct {
	DataSourceId string   `json:"dataSourceId"` // 数据源id（create方法返回的data.dataSource.id）
	Orders       *[]Order `json:"orders"`       // 订单列表 数组最大长度 50
}

// Order
type Order struct {
	ExternalOrderId  string           `json:"external_order_id"`        // 商家订单号
	CreateTime       string           `json:"create_time"`              // 订单创建时间，unix时间戳 字段长度为 13 字节
	OrderSource      string           `json:"order_source"`             // 订单来源,枚举值:商家小程序：wxapp；商家app：app；商家H5：mobileweb；商家pcweb：pcweb；线下人工pos：offstore_pos_manual；线下自助收银：offstore_pos_self_help；其他：other
	OrderType        int              `json:"order_type"`               // 订单类型；1：普通订单；2：充值订单；3：消费订单；普通购买商品订单传 1 即可
	BrandId          string           `json:"brand_id,omitempty"`       // 订单品牌 id
	BrandName        string           `json:"brand_name,omitempty"`     // 订单品牌名称
	GoodsNumTotal    int              `json:"goods_num_total"`          // 订单商品总数量
	GoodsWeight      float64          `json:"goods_weight,omitempty"`   // 订单商品总重量，默认单位为克
	GoodsAmountTotal float64          `json:"goods_amount_total"`       // 商品总金额，单位默认为元 注：已含单品级别优惠的商品金额，如单品直降
	FreightAmount    float64          `json:"freight_amount"`           // 订单运费，单位默认为元 注：运费为0时，传0.00
	OrderAmount      float64          `json:"order_amount"`             // 订单金额，单位默认为元 注：商品总金额+运费金额=订单金额
	PayableAmount    float64          `json:"payable_amount"`           // 订单应付金额，单位默认为元 注：订单金额-订单级别的优惠金额（如：订单满减）=订单应付金额
	PaymentAmount    float64          `json:"payment_amount"`           // 实付金额，单位默认为元 注：订单应付金额-支付优惠金额（如：微信支付优惠、招商银行优惠等）=订单实付金额
	OrderStatus      string           `json:"order_status"`             // 主订单状态，1110待支付，1150已支付待发货，1160已发货，1180销售完成/已收货，1280退款中，1290退货完成
	StatusChangeTime string           `json:"status_change_time"`       // 状态变更时间，unix毫秒级时间，如 order_status状态为 1150 ，则传 1150（已支付待发货）状态变更的时间
	UserInfo         *UserInfo        `json:"user_info"`                // 用户信息，json格式
	GoodsInfo        *[]GoodsInfo     `json:"goods_info"`               // 主订单商品信息，数组类型，每个sku存一个数组单位
	PromotionInfo    *[]PromotionInfo `json:"promotion_info,omitempty"` // 主订单商品参与的活动信息，数组类型
	CouponInfo       *[]CouponInfo    `json:"coupon_info,omitempty"`    // 主订单用到的券信息，数组类型
	PaymentInfo      *[]PaymentInfo   `json:"payment_info,omitempty"`   // 主订单每种支付方式的支付信息,order_status = 1110时 payment_info非必填，其他状态码必填
	ExpressInfo      *ExpressInfo     `json:"express_info,omitempty"`   // 快递信息
	InvoiceInfo      *[]InvoiceInfo   `json:"invoice_info,omitempty"`   // 发票信息，类型为数组
	PointsTotal      float64          `json:"points_total,omitempty"`   // 订单赠送总积分
	IsDeleted        int              `json:"is_deleted"`               // 商家标记订单已删除，0-未删除，1-已删除
}

// UserInfo
type UserInfo struct {
	OpenId             string `json:"open_id"`                         // 下单人 open_id，order_source = wxapp时，必填
	AppId              string `json:"app_id,omitempty"`                // 小程序或公众号的appid
	UnionId            string `json:"union_id,omitempty"`              // 下单人 union_id
	UserPhone          string `json:"user_phone,omitempty"`            // 下单人手机号
	UserId             string `json:"user_id,omitempty"`               // 下单人用户 id
	MemberId           string `json:"member_id,omitempty"`             //	下单人会员号
	UserFirstOrderTime string `json:"user_first_order_time,omitempty"` // 下单人在KA注册后首次下单时间，格式为UNIX时间戳 字段长度为 13 字节
}

//GoodsInfo
type GoodsInfo struct {
	ExternalSkuId   string            `json:"external_sku_id"`             // sku 编号
	SkuNameChinese  string            `json:"sku_name_chinese"`            // sku 名称
	PrimaryImageUrl string            `json:"primary_image_url,omitempty"` // 商品主图
	GoodsAmount     float64           `json:"goods_amount"`                // 单件商品原价，单位默认为元
	PaymentAmount   float64           `json:"payment_amount"`              // 多件商品实付金额（分摊了优惠的金额）,单位默认为元
	IsGift          int               `json:"is_gift,omitempty"`           // 是否赠品，0代表非赠品，1代表赠品
	ExternalSpuId   string            `json:"external_spu_id"`             // sku 所属 spu 编号，若无 spu，传输内容请与 external_sku_id 保持一致
	SpuNameChinese  string            `json:"spu_name_chinese"`            // spu 名称，若无 spu，传输内容请与 sku_name_chinese 保持一致
	SaleUnit        string            `json:"sale_unit,omitempty"`         // 商品售卖单位
	CategoryId      string            `json:"category_id,omitempty"`       // 末级类目 id
	CategoryName    string            `json:"category_name,omitempty"`     // 末级类目名称
	GoodsNum        int               `json:"goods_num"`                   // 商品数量
	GoodsWeight     float64           `json:"goods_weight,omitempty"`      // 商品重量，单位默认为克
	StoreInfo       *StoreInfo        `json:"store_info,omitempty"`        // 主订单销售门店信息
	ChanInfo        *[]ChanInfo       `json:"chan_info,omitempty"`         // 主订单来源渠道，数组类型
	CommissionInfo  *[]CommissionInfo `json:"commission_info,omitempty"`   // 佣金，json字符串
}

type StoreInfo struct {
	ExternalStoreId string `json:"external_store_id,omitempty"` // 主订单销售门店id
	StoreName       string `json:"store_name,omitempty"`        // 主订单销售门店名称
	StoreCity       string `json:"store_city,omitempty"`        // 主订单销售门店所属城市
}

// ExpressInfo
type ExpressInfo struct {
	LogisticsStatus      string                `json:"logistics_status,omitempty"`       // 订单物流状态
	GoodsTotalWeight     float64               `json:"goods_total_weight,omitempty"`     // 商品总重量，单位默认为克
	ReceiverName         string                `json:"receiver_name,omitempty"`          // 收件人姓名
	ReceiverPhone        string                `json:"receiver_phone,omitempty"`         // 收件人联系电话
	ReceiverAddress      string                `json:"receiver_address,omitempty"`       // 收件人地址
	ReceiverCountryCode  string                `json:"receiver_country_code,omitempty"`  // 国家编码
	ReceiverProvinceCode string                `json:"receiver_province_code,omitempty"` // 省份编码
	ReceiverCityCode     string                `json:"receiver_city_code,omitempty"`     // 城市编码
	ReceiverDistrictCode string                `json:"receiver_district_code,omitempty"` // 区编码
	ExpectedDeliveryTime string                `json:"expected_delivery_time,omitempty"` // 期望送货时间段，格式为“起始时间-结束时间”，如"9:00-12:00"
	ExpectedDeliveryDate string                `json:"expected_delivery_date,omitempty"` // 期望送货日期，格式“YYYYMMDD”
	ExpressPackageInfo   *[]ExpressPackageInfo `json:"express_package_info,omitempty"`   // 包裹信息，struct类型
}

//ChanInfo
type ChanInfo struct {
	ChanWeapp      *ChanWeapp  `json:"chan_weapp,omitempty"`        // 小程序渠道
	ChanCustom     *ChanCustom `json:"chan_custom,omitempty"`       // 自定义渠道
	ChanId         string      `json:"chan_id,omitempty"`           // 智慧零售入口小程序必传，引流渠道编码
	ChanReferAppId string      `json:"chan_refer_app_id,omitempty"` // 智慧零售入口小程序必传，来源小程序或公众号appid
}

// ChanWeapp
type ChanWeapp struct {
	ChanScene string `json:"chan_scene,omitempty"` // 小程序场景值
}

// ChanCustom
type ChanCustom struct {
	ChanCustomId       string `json:"chan_custom_id,omitempty"`         // 自定义渠道的标识符，是自定义渠道的最小粒度
	ChanCustomIdDesc   string `json:"chan_custom_id_desc,omitempty"`    // 自定义渠道的描述
	ChanCustomCat3     string `json:"chan_custom_cat_3,omitempty"`      // 3级自定义渠道的标识符，3级是针对4级的分类，要求4级数据必须存在
	ChanCustomCat3Desc string `json:"chan_custom_cat_3_desc,omitempty"` // 3级自定义渠道的描述，若chan_custom_cat_3存在则必须存在
	ChanCustomCat2     string `json:"chan_custom_cat_2,omitempty"`      // 2级自定义渠道的标识符，2级是针对3级的分类，要求3、4级数据必须存在
	ChanCustomCat2Desc string `json:"chan_custom_cat_2_desc,omitempty"` // 2级自定义渠道的描述，若chan_custom_cat_2存在则必须存在
	ChanCustomCat1     string `json:"chan_custom_cat_1,omitempty"`      // 1级自定义渠道的标识符，1级是针对2级的分类，要求2、3、4级数据必须存在
	ChanCustomCatDesc  string `json:"chan_custom_cat_desc,omitempty"`   // 1级自定义渠道的描述，若chan_custom_cat_1存在则必须存在
}

//PromotionInfo
type PromotionInfo struct {
	PromotionType       int     `json:"promotion_type,omitempty"`        // 活动类型，枚举值如下：1：订单满减；2：订单满折；3：订单赠品；4：订单积分；5：单品降价；6：单品打折；7：单品赠品；8：单品积分；9：X件N折；10：X元N件；11：拼团
	ExternalPromotionId string  `json:"external_promotion_id,omitempty"` // 活动 id
	PromotionName       string  `json:"promotion_name,omitempty"`        // 活动名称
	PromotionAmount     float64 `json:"promotion_amount,omitempty"`      // 订单维度的活动优惠金额，单位默认为元
}

//CouponInfo
type CouponInfo struct {
	CouponType        int             `json:"coupon_type,omitempty"`         // 卡券类型；1：商家券；2：微信支付券
	CouponAmountTotal float64         `json:"coupon_amount_total,omitempty"` // 该类券优惠金额总额，单位默认为元
	CouponDetail      *[]CouponDetail `json:"coupon_detail,omitempty"`       // 该类券的细节券信息
}

//PaymentInfo
type PaymentInfo struct {
	PaymentType string  `json:"payment_type,omitempty"` // 支付方式，见<枚举列表>页
	TransId     string  `json:"trans_id,omitempty"`     // 微信支付订单ID/流水号
	TransAmount float64 `json:"trans_amount,omitempty"` // 金额，单位默认为元
}

//InvoiceInfo
type InvoiceInfo struct {
	IfNeedInvoice          bool   `json:"if_need_invoice"`                    // 是否需要发票，true代表需要，false代表不需要
	InvoiceType            string `json:"invoice_type,omitempty"`             // 发票类型，枚举值，取值如下： 1000(增值税专用发票) 1001(普通发票) 1002(机动车专用发票) 1003(机打发票) 1004(定额发票 ) 1005(剪开式发票) 1006（其他）
	InvoiceTitle           string `json:"invoice_title,omitempty"`            // 发票抬头
	InvoiceContent         string `json:"invoice_content,omitempty"`          // 发票内容
	InvoiceAdditionInfo    string `json:"invoice_addition_info,omitempty"`    // 发票附加信息
	InvoiceCompany         string `json:"invoice_company,omitempty"`          // 公司名称
	InvoiceTaxpayer        string `json:"invoice_taxpayer,omitempty"`         // 纳税人识别号
	RegistryAddress        string `json:"registry_address,omitempty"`         // 注册地址
	RegistryPhone          string `json:"registry_phone,omitempty"`           // 注册电话
	RegistryBankName       string `json:"registry_bank_name,omitempty"`       // 开户银行
	RegistryBankAccount    string `json:"registry_bank_account,omitempty"`    // 开户账号
	InvoiceDeliveryAddress string `json:"invoice_delivery_address,omitempty"` // 发票收件地址
	InvoiceDeliveryName    string `json:"invoice_delivery_name,omitempty"`    // 发票收件人姓名
	InvoiceDeliveryPhone   string `json:"invoice_delivery_phone,omitempty"`   // 发票收件人电话
	InvoiceNum             string `json:"invoice_num,omitempty"`              // 发票号码
}

//CommissionInfo
type CommissionInfo struct {
	CommissionType int     `json:"commission_type,omitempty"` // 佣金类型，枚举值如下：1：按比例提成；2：按金额提成
	CommissionFee  float64 `json:"commission_fee,omitempty"`  // 佣金金额，单位元
}

//CouponDetail
type CouponDetail struct {
	ExternalCouponId string  `json:"external_coupon_id,omitempty"` // 券 id
	CouponBatchId    string  `json:"coupon_batch_id,omitempty"`    // 券批次 id（该字段需要在优惠券接口中添加卡券批次）
	CouponName       string  `json:"coupon_name,omitempty"`        // 券名称
	CouponAmount     float64 `json:"coupon_amount,omitempty"`      // 该张券优惠金额，单位默认为元
}

//ExpressPackageInfo
type ExpressPackageInfo struct {
	ExpressCompanyCode string                    `json:"express_company_code,omitempty"` // 物流公司编码，枚举类型，枚举值请参见文章后面的“物流商 code”
	ExpressCompanyName string                    `json:"express_company_name,omitempty"` // 物流公司名称
	ExpressCode        string                    `json:"express_code,omitempty"`         // 运单号
	ShipTime           string                    `json:"ship_time,omitempty"`            // 发货时间，格式为时间戳 字段长度为 13 字节
	ExpressPage        *ExpressPage              `json:"express_page,omitempty"`         // 运费跳转页面，json字符串
	ExpressPackageInfo *[]ExpressPackageInfoItem `json:"express_package_info,omitempty"` // 物流包裹信息
}

// ExpressPage
type ExpressPage struct {
	MiniprogramPath  string `json:"miniprogram_path,omitempty"`  // 快递详情页跳转链接（小程序页面，小程序填此字段）
	MiniprogramAppid string `json:"miniprogram_appid,omitempty"` // 小程序APPID，填写了miniprogram_path需填此字段
	MiniprogramH5    string `json:"miniprogram_h5,omitempty"`    // 快递详情页跳转链接（h5页面，公众号填此字段）
}

// ExpressPackageInfoItem
type ExpressPackageInfoItem struct {
	ExternalSkuId string `json:"external_sku_id,omitempty"` // 商品sku id
	Number        int    `json:"number,omitempty"`          // 商品数量
}
