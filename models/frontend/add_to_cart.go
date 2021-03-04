package frontend

type AddToCartReq struct {
	Type       string          `json:"type"`
	From       string          `json:"from"`
	TrackingId string          `json:"tracking_id"`
	LogId      int             `json:"log_id"`
	Props      *AddToCartProps `json:"props"`
}

type AddToCartProps struct {
	Page           string     `json:"page"`                     //行为发生的小程序页面路径，开头不要加 /，路径后可带查询参数 例pages/product?sku_id=AOdjf7u
	PageTitle      string     `json:"page_title,omitempty"`     //行为发生的小程序页面标题 例商品详情、商城首页
	SrSdkVersion   string     `json:"sr_sdk_version,omitempty"` //sdk版本号 例1.1.6
	IsSdkAutoTrack bool       `json:"is_sdk_auto_track"`
	Time           string     `json:"time"`      //行为发生时间 13位时间戳 例1560409473714
	WxUser         *WxUser    `json:"wx_user"`   //用户相关属性
	Chan           *Chan      `json:"chan"`      //渠道相关属性
	Component      *Component `json:"component"` //组件是视图上的一个区块元素

	ActionType   string         `json:"action_type"`             //动作类型 append_to_cart：商品详情页中点击加入购物车按钮；append_to_cart_in_cart：点击购物车加号；remove_from_cart：点击购物车减号或从购物车彻底删除
	Sku          *Sku           `json:"sku"`                     //商品sku
	Spu          *Spu           `json:"spu"`                     //商品spu
	SkuCategory  *[]SkuCategory `json:"sku_category,omitempty"`  //商品类目路径
	ShippingShop *ShippingShop  `json:"shipping_shop,omitempty"` //商品发货门店
	Sale         *Sale          `json:"sale,omitempty"`          //销售信息
	SkuNum       int            `json:"sku_num"`                 //本次操作的加购件数，如购物车原有 1 件商品，此次加购了 2 件，sku_num 传 2，或购物车原有 3 件商品，此次减去了 1 件，sku_num 传 1
}

type Sku struct {
	SkuId   string `json:"sku_id"`   //商品skuid,唯一标识符 例AKDFie8-1
	SkuName string `json:"sku_name"` //商品名称 例妃子笑荔枝礼盒装（2KG）
}

type Spu struct {
	SpuId   string `json:"spu_id"`   //商品spuid 例AKDFie8
	SpuName string `json:"spu_name"` //商品名称 例妃子笑荔枝
}

type SkuCategory struct {
	SkuCateId      string `json:"sku_cate_id,omitempty"`       //商品类目id，唯一标识符，若填写了`sku_category`下任一字段则必填
	SkuCatName     string `json:"sku_cat_name,omitempty"`      //商品类目名称，若填写了`sku_category`下任一字段则必填
	SkuParentCatId string `json:"sku_parent_cat_id,omitempty"` //商品父级类目id，若填写了`sku_category`下任一字段则必填，若已是顶级类目，填null
}

type ShippingShop struct {
	ShippingShopId   string `json:"shipping_shop_id,omitempty"`   //商品发货门店id；当商品所属门店同于行为发生门店时，shop_id 与 chan_shop_id 值一致，若填写了`shipping_shop_name`则必填
	ShippingShopName string `json:"shipping_shop_name,omitempty"` //商品发货门店名称，若填写了`shipping_shop_id`则必填
}

type Sale struct {
	OriginalPrice float64 `json:"original_price,omitempty"` //商品当前原价单价（元），一般是划线价 例49.9，若填写了`current_price`则必填
	CurrentPrice  float64 `json:"current_price,omitempty"`  //商品当前现价（元） 例30.0，若填写了`current_price`则必填
}
