package api

import (
	"encoding/json"
	"github.com/qinyer/sr-client-go/context"
	"github.com/qinyer/sr-client-go/models"
	"github.com/qinyer/sr-client-go/utils"
)

type Order struct {
	*context.Context
}

// NewOrder 获取订单相关上报实例
func NewOrder(ctx *context.Context) *Order {
	return &Order{ctx}
}

// AddOrder 添加/更新订单
func (order *Order) AddOrder(req *models.OrderReq) (res utils.ReportCommonRes, err error) {
	response, err := utils.NewHttpClient(order.Context).HttpPostJson("/order/add_order", req)
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(response), &res)
	if err != nil {
		return
	}
	return
}

// AddOrderSum 汇总订单
func (order *Order) AddOrderSum(req *models.OrderSumReq) (res utils.ReportCommonRes, err error) {
	response, err := utils.NewHttpClient(order.Context).HttpPostJson("/order/add_order_sum", req)
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(response), &res)
	if err != nil {
		return
	}
	return
}

// AddPromotion 添加/更新活动信息
func (order *Order) AddPromotion(req *models.OrderPromotionReq) (res utils.ReportCommonRes, err error) {
	response, err := utils.NewHttpClient(order.Context).HttpPostJson("/order/add_promotion", req)
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(response), &res)
	if err != nil {
		return
	}
	return
}

// AddCoupon 添加/更新卡券信息
func (order *Order) AddCoupon(req *models.OrderCouponReq) (res utils.ReportCommonRes, err error) {
	response, err := utils.NewHttpClient(order.Context).HttpPostJson("/order/add_coupon", req)
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(response), &res)
	if err != nil {
		return
	}
	return
}
