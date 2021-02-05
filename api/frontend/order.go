package frontend

import (
	"encoding/json"
	"github.com/qinyer/sr-client-go/context"
	"github.com/qinyer/sr-client-go/models/frontend"
	"github.com/qinyer/sr-client-go/utils"
)

type Order struct {
	*context.Context
}

// NewOrder 获取订单行为上报实例
func NewOrder(ctx *context.Context) *Order {
	return &Order{ctx}
}

// Report 订单行为上报
func (c *Order) Report(req *frontend.CustomOrderReq) (res utils.ReportFrontendRes, err error) {
	response, err := utils.NewFrontendHttpClient(c.Context).PostJson("", req)
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(response), &res)
	if err != nil {
		return
	}
	return
}
