package api

import (
	"encoding/json"
	"sr-client-go/context"
	"sr-client-go/models"
	"sr-client-go/utils"
)

type SalesInfo struct {
	*context.Context
}

// NewSalesInfo 获取销售信息上报实例
func NewSalesInfo(ctx *context.Context) *SalesInfo {
	salesInfo := new(SalesInfo)
	salesInfo.Context = ctx
	return salesInfo
}

// Add 添加/更新销售信息
func (ss *SalesInfo) Add(req *models.SalesInfoReq) (res utils.ReportCommonRes, err error) {
	response, err := utils.NewHttpClient(ss.AppId, ss.AppSecret, ss.Prod).HttpPostJson("/salesinfo/add", req)
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(response), &res)
	if err != nil {
		return
	}
	return
}
