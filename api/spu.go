package api

import (
	"encoding/json"
	"github.com/qinyer/sr-client-go/context"
	"github.com/qinyer/sr-client-go/models"
	"github.com/qinyer/sr-client-go/utils"
)

type Spu struct {
	*context.Context
}

// NewSpu 获取商品spu上报实例
func NewSpu(ctx *context.Context) *Spu {
	return &Spu{ctx}
}

// Add 添加/更新商品spu
func (spu *Spu) Add(req *models.SpuReq) (res utils.ReportCommonRes, err error) {
	response, err := utils.NewHttpClient(spu.Context).HttpPostJson("/spus/add", req)
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(response), &res)
	if err != nil {
		return
	}
	return
}