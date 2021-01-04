package models

// AnalysisVisitPageReq 上报页面访问请求
type AnalysisVisitPageReq struct {
	DataSourceId string             `json:"dataSourceId"`
	RawMsg       *[]VisitPageRawMsg `json:"rawMsg"`
}

// VisitPageRawMsg
type VisitPageRawMsg struct {
	RefDate string                 `json:"ref_date"`
	List    *[]VisitPageRawMsgList `json:"list"`
}

// VisitPageRawMsgList
type VisitPageRawMsgList struct {
	PagePath       string  `json:"page_path"`
	PageVisitPv    int     `json:"page_visit_pv"`
	PageVisitUv    int     `json:"page_visit_uv"`
	PageStaytimePv float64 `json:"page_staytime_pv"`
	EntrypagePv    int     `json:"entrypage_pv"`
	ExitpagePv     int     `json:"exitpage_pv"`
	PageSharePv    int     `json:"page_share_pv"`
	PageShareUv    int     `json:"page_share_uv"`
}
