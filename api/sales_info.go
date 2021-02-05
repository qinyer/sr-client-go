package api

import (
	"encoding/json"
	"github.com/qinyer/sr-client-go/context"
	"github.com/qinyer/sr-client-go/models"
	"github.com/qinyer/sr-client-go/utils"
)

type SalesInfo struct {
	*context.Context
}

// NewSalesInfo 获取销售信息上报实例
func NewSalesInfo(ctx *context.Context) *SalesInfo {
	return &SalesInfo{ctx}
}

// Add 添加/更新销售信息
func (ss *SalesInfo) Add(req *models.SalesInfoReq) (res utils.ReportCommonRes, err error) {
	response, err := utils.NewHttpClient(ss.Context).HttpPostJson("/salesinfo/add", req)
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(response), &res)
	if err != nil {
		return
	}
	return
}
