package utils

import (
	"fmt"
	"github.com/qinyer/sr-client-go/context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// 有数前端上报地址
const YsFrontendUrl = "https://zhls.qq.com/api/report"

// 有数前端上报
type YsFrontendHttp struct {
	ctx *context.Context
}

// NewFrontendHttpClient 获取有数前端上报实例
func NewFrontendHttpClient(ctx *context.Context) *YsFrontendHttp {
	return &YsFrontendHttp{ctx}
}

// RequestFrontend 请求
func (ys *YsFrontendHttp) RequestFrontend(method string, reqUrl string, reqBody string, headers ...map[string]string) (body string, err error) {
	var req *http.Request
	fullUrl := ys.buildFrontendRequestUrl(reqUrl)
	// 解析HOST
	u, err := url.Parse(fullUrl)
	if err != nil {
		ys.ctx.Logger.Error(err)
		return
	}
	defaultDomain := u.Host
	// 默认超时10S
	if ys.ctx.Timeout == 0 {
		ys.ctx.Timeout = time.Duration(10) * time.Second
	}
	client := &http.Client{
		Timeout: ys.ctx.Timeout,
	}
	req, err = http.NewRequest(method, fullUrl, strings.NewReader(reqBody))
	if err != nil {
		ys.ctx.Logger.Error(err)
		return
	}

	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Charset", "utf-8;")
	req.Header.Set("Content-Type", "application/json;")
	req.Header.Set("Host", defaultDomain)
	req.Header.Set("Timeout", fmt.Sprintf("%v", ys.ctx.Timeout))
	if len(headers) > 0 {
		for k, v := range headers[0] {
			req.Header.Set(k, v)
		}
	}
	resp, err := client.Do(req)
	if err != nil {
		ys.ctx.Logger.Error(err)
		return
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ys.ctx.Logger.Error(err)
		return
	}

	body = string(content)
	ys.ctx.Logger.Printf("\n request: \n method: %s, reqUrl: %s, reqBody: %s\n response: \n %s", method, reqUrl, reqBody, body)
	if resp.StatusCode != 200 || body == "" {
		err = fmt.Errorf(fmt.Sprintf("url：%s，body：%s", fullUrl, body))
		return
	}
	// 打印日志
	return
}

// HttpGetJson GET
func (ys *YsFrontendHttp) GetJson(path string, params url.Values, headers ...map[string]string) (body string, err error) {
	reqUrl := fmt.Sprintf("%s?%s", path, params.Encode())
	body, err = ys.RequestFrontend("GET", reqUrl, "", headers...)

	return
}

// HttpPostJson POST
func (ys *YsFrontendHttp) PostJson(path string, params interface{}, headers ...map[string]string) (body string, err error) {
	body, err = ys.RequestFrontend("POST", path, ToBufferString(params), headers...)

	return
}

// getFullUrl 获取请求url
func (ys *YsFrontendHttp) buildFrontendRequestUrl(reqUrl string) string {
	// 请求时间戳（秒级）
	timeLocation, _ := time.LoadLocation("Etc/GMT")
	timestamp := time.Now().In(timeLocation).Unix()
	// 随机字符串
	nonce := RandomStr(32)
	return fmt.Sprintf("%s%s?app_id=%s&nonce=%s&timestamp=%d", YsFrontendUrl, reqUrl, ys.ctx.AppId, nonce, timestamp)
}
