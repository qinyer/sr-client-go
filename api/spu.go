package api

import (
	"encoding/json"
	"sr-client-go/context"
	"sr-client-go/models"
	"sr-client-go/utils"
)

type Spu struct {
	*context.Context
}

// NewSpu 获取商品spu上报实例
func NewSpu(ctx *context.Context) *Spu {
	spu := new(Spu)
	spu.Context = ctx
	return spu
}

// Add 添加/更新商品spu
func (spu *Spu) Add(req *models.SpuReq) (res utils.ReportCommonRes, err error) {
	response, err := utils.NewHttpClient(spu.AppId, spu.AppSecret, spu.Prod).HttpPostJson("/spus/add", req)
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(response), &res)
	if err != nil {
		return
	}
	return
}