package frontend

type CustomOrderReq struct {
	Type  string           `json:"type"`
	Props *CustomOrderProp `json:"props"`
}

type CustomOrderProp struct {
	Page         string     `json:"page"`           //行为发生的小程序页面路径，开头不要加 /，路径后可带查询参数 例pages/product?sku_id=AOdjf7u
	PageTitle    string     `json:"page_title"`     //行为发生的小程序页面标题 例商品详情、商城首页
	SrSdkVersion string     `json:"sr_sdk_version"` //sdk版本号 例1.1.6
	Time         string     `json:"time"`           //行为发生时间 13位时间戳 例1560409473714
	WxUser       *WxUser    `json:"wx_user"`        //用户相关属性
	Chan         *Chan      `json:"chan"`           //渠道相关属性
	Component    *Component `json:"component"`      //组件是视图上的一个区块元素

	Order     *Order      `json:"order"`      //用户订单信息
	SubOrders *[]SubOrder `json:"sub_orders"` //订单的金额信息，注意为[]结构，
}

type Order struct {
	OrderId     string `json:"order_id"`     //商户侧订单号，在商户系统内订单的唯一标识符
	OrderTime   string `json:"order_time"`   //订单下单时间，13位，毫秒，任何 order_status都必填
	OrderStatus string `json:"order_status"` //订单状态，订单状态与后端状态的对应关系见本页底部，give_order：用户提交订单；cancel_pay：用户关闭支付密码浮层；cancel_give_order：用户取消订单；pay：用户发起支付；payed：用户完成支付（影响有数实时订单统计）refund：用户发起退货退款
}

type SubOrder struct {
	SubOrderId string   `json:"sub_order_id"`    //同order_id
	OrderAmt   float64  `json:"order_amt"`       //填写订单金额，单位默认为元
	PayAmt     float64  `json:"pay_amt"`         //订单应付金额，单位默认为元
	Extra      struct{} `json:"extra,omitempty"` //订单额外属性
}
