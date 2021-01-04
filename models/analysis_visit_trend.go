package models

// AnalysisVisitTrendReq 上报日访问趋势请求
type AnalysisVisitTrendReq struct {
	DataSourceId string              `json:"dataSourceId"`
	RawMsg       *[]VisitTrendRawMsg `json:"rawMsg"`
}

// AnalysisVisitTrendRawMsg
type VisitTrendRawMsg struct {
	List *[]VisitTrendRawMsgList `json:"list"`
}

// VisitTrendRawMsgList
type VisitTrendRawMsgList struct {
	RefDate         string  `json:"ref_date"`
	SessionCnt      int     `json:"session_cnt"`
	VisitPv         int     `json:"visit_pv"`
	VisitUv         int     `json:"visit_uv"`
	VisitUvNew      int     `json:"visit_uv_new"`
	StayTimeSession int     `json:"stay_time_session"`
	VisitDepth      float64 `json:"visit_depth"`
}
