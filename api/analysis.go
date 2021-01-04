package api

import (
	"encoding/json"
	"sr-client-go/context"
	"sr-client-go/models"
	"sr-client-go/utils"
)

type Analysis struct {
	*context.Context
}

// NewAnalysis 获取分析上报实例
func NewAnalysis(ctx *context.Context) *Analysis {
	analysis := new(Analysis)
	analysis.Context = ctx
	return analysis
}

// AddVisitPage 上报页面访问
func (ay *Analysis) AddVisitPage(req *models.AnalysisVisitPageReq) (res utils.ReportCommonRes, err error) {
	response, err := utils.NewHttpClient(ay.AppId, ay.AppSecret, ay.Prod).HttpPostJson("/analysis/add_wxapp_visit_page", req)
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(response), &res)
	if err != nil {
		return
	}
	return
}

// AddVisitTrend 上报日访问趋势
func (ay *Analysis) AddVisitTrend(req *models.AnalysisVisitTrendReq) (res utils.ReportCommonRes, err error) {
	response, err := utils.NewHttpClient(ay.AppId, ay.AppSecret, ay.Prod).HttpPostJson("/analysis/add_wxapp_daily_visit_trend", req)
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(response), &res)
	if err != nil {
		return
	}
	return
}

// AddVisitDistribution 上报访问分布
func (ay *Analysis) AddVisitDistribution(req *models.AnalysisVisitDistributionReq) (res utils.ReportCommonRes, err error) {
	response, err := utils.NewHttpClient(ay.AppId, ay.AppSecret, ay.Prod).HttpPostJson("/analysis/add_wxapp_visit_distribution", req)
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(response), &res)
	if err != nil {
		return
	}
	return
}
