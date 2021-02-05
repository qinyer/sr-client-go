package frontend

import (
	"encoding/json"
	"github.com/qinyer/sr-client-go/context"
	"github.com/qinyer/sr-client-go/models/frontend"
	"github.com/qinyer/sr-client-go/utils"
)

type Cart struct {
	*context.Context
}

// NewCart 获取商品加购实例
func NewCart(ctx *context.Context) *Cart {
	return &Cart{ctx}
}

// Report 商品加购上报
func (c *Cart) Report(req *frontend.AddToCartReq) (res utils.ReportFrontendRes, err error) {
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
