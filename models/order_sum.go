package models

// OrderSumReq 汇总订单请求
type OrderSumReq struct {
	DataSourceId string      `json:"dataSourceId"` // 数据源id（create方法返回的data.dataSource.id）
	Orders       *[]OrderSum `json:"orders"`       // 订单列表 数组最大长度 50
}

// OrderSum
type OrderSum struct {
	RefDate            string  `json:"ref_date"`              // 日期，unix时间戳，字段长度为13字节
	GiveOrderAmountSum float64 `json:"give_order_amount_sum"` // 该日期的下单金额之和
	GiveOrderNumSum    int     `json:"give_order_num_sum"`    // 该日期的下单数量之和
	PaymentAmountSum   float64 `json:"payment_amount_sum"`    // 该日期的支付金额之和
	PayedNumSum        int     `json:"payed_num_sum"`         // 该日期的支付数量之和
}
