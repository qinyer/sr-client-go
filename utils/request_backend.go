package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/qinyer/sr-client-go/context"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const TestBaseUrl = "https://test.zhls.qq.com" // 测试环境URL
const ProdBaseUrl = "https://zhls.qq.com"      // 生产环境URL
const PrefixPath = "/data-api/v1"              // 路径前缀

// YsHttp 详见 -> https://mp.zhls.qq.com/youshu-docs/develop/dev_account/dev_account_access.html
type YsHttp struct {
	ctx *context.Context
}

// 获取实例
func NewHttpClient(ctx *context.Context) (ysHttp *YsHttp) {
	return &YsHttp{ctx}
}

// SetPrefix 设置请求前缀
func (ys *YsHttp) SetPrefix(prefix string) *YsHttp {
	ys.ctx.PrefixPath = prefix
	return ys
}

// SetTimeout 设置超时时间
func (ys *YsHttp) SetTimeout(timeout int64) *YsHttp {
	ys.ctx.Timeout = time.Duration(timeout) * time.Second
	return ys
}

// Request 请求
func (ys *YsHttp) Request(method string, reqUrl string, reqBody string, headers ...map[string]string) (body string, err error) {
	var req *http.Request
	fullUrl, err := ys.signatureAndBuildRequestUrl(reqUrl)
	if err != nil {
		ys.ctx.Logger.Error(err)
		return
	}
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
func (ys *YsHttp) HttpGetJson(path string, params url.Values, headers ...map[string]string) (body string, err error) {
	reqUrl := fmt.Sprintf("%s?%s", path, params.Encode())
	body, err = ys.Request("GET", reqUrl, "", headers...)

	return
}

// HttpPostJson POST
func (ys *YsHttp) HttpPostJson(path string, params interface{}, headers ...map[string]string) (body string, err error) {
	body, err = ys.Request("POST", path, ToBufferString(params), headers...)

	return
}

// getFullUrl 获取请求url
func (ys *YsHttp) signatureAndBuildRequestUrl(reqUrl string) (url string, err error) {
	// 签名算法
	signType := "sha256"
	// 请求时间戳（秒级）
	timeLocation, _ := time.LoadLocation("Etc/GMT")
	timestamp := time.Now().In(timeLocation).Unix()
	// 随机字符串
	nonce := RandomStr(32)
	// 待签名字符串
	conStr := fmt.Sprintf("app_id=%s&nonce=%s&sign=%s&timestamp=%d", ys.ctx.AppId, nonce, signType, timestamp)
	// sha256
	h := hmac.New(sha256.New, []byte(ys.ctx.AppSecret))
	_, err = io.WriteString(h, conStr)
	if err != nil {
		return
	}
	signature := hex.EncodeToString(h.Sum(nil))
	// 默认走测试环境
	host := TestBaseUrl
	// 开启生成环境
	if ys.ctx.Prod {
		host = ProdBaseUrl
	}
	prefix := PrefixPath
	if ys.ctx.PrefixPath != "" {
		prefix = ys.ctx.PrefixPath
	}
	signatureStr := fmt.Sprintf("app_id=%s&nonce=%s&sign=%s&timestamp=%d&signature=%s", ys.ctx.AppId, nonce, signType, timestamp, signature)
	// 查询参数拼接签名公共参数
	if strings.Contains(reqUrl, "?") {
		tmp := strings.Split(reqUrl, "?")
		reqUrl = fmt.Sprintf("%s?%s&%s", tmp[0], tmp[1], signatureStr)
	} else {
		reqUrl = fmt.Sprintf("%s?%s", reqUrl, signatureStr)
	}

	url = fmt.Sprintf("%s%s%s", host, prefix, reqUrl)
	return
}
