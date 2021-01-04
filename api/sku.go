package api

import (
	"encoding/json"
	"github.com/qinyer/sr-client-go/context"
	"github.com/qinyer/sr-client-go/models"
	"github.com/qinyer/sr-client-go/utils"
)

type Sku struct {
	*context.Context
}

// NewSku 获取商品sku上报实例
func NewSku(ctx *context.Context) *Sku {
	sku := new(Sku)
	sku.Context = ctx
	return sku
}

// Add 添加/更新商品sku
func (sku *Sku) Add(req *models.SkuReq) (res utils.ReportCommonRes, err error) {
	response, err := utils.NewHttpClient(sku.AppId, sku.AppSecret, sku.Prod).HttpPostJson("/sku/add", req)
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(response), &res)
	if err != nil {
		return
	}
	return
}