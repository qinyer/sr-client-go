package models

// AnalysisVisitDistributionReq 上报访问分布请求
type AnalysisVisitDistributionReq struct {
	DataSourceId string                     `json:"dataSourceId"`
	RawMsg       *[]VisitDistributionRawMsg `json:"rawMsg"`
}

// VisitDistributionRawMsg
type VisitDistributionRawMsg struct {
	RefDate string                         `json:"ref_date"`
	List    *[]VisitDistributionRawMsgList `json:"list"`
}

// VisitDistributionRawMsgList
type VisitDistributionRawMsgList struct {
	Index    string                         `json:"index"`
	ItemList *[]VisitDistributionRawMsgItem `json:"item_list"`
}

// VisitDistributionRawMsgItem
type VisitDistributionRawMsgItem struct {
	Key   int `json:"key"`
	Value int `json:"value"`
}
