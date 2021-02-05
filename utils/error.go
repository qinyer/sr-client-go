package utils

// CommonRes 公共返回
type CommonError struct {
	RetCode int    `json:"retcode"`
	ErrMsg  string `json:"errmsg"`
}

// ReportCommonRes 上报公共返回
type ReportCommonRes struct {
	CommonError
	Data interface{} `json:"data"`
}

// ReportFrontendRes 前端上报公共返回
type ReportFrontendRes struct {
	Code int    `json:"code"` // 0-成功 其他-失败
	Msg  string `json:"msg"`
}
