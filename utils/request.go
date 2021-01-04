package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const TestBaseUrl = "https://test.zhls.qq.com" // 测试环境URL
const ProdBaseUrl = "https://zhls.qq.com"      // 生产环境URL
const PrefixPath = "/data-api/v1"               // 路径前缀

// YsHttp 详见 -> https://mp.zhls.qq.com/youshu-docs/develop/dev_account/dev_account_access.html
type YsHttp struct {
	appId      string        // appId 数据中心分配的有数应用ID
	appSecret  string        // appSecret 数据中心分配的有数签名密钥，需要绝对保密
	prod       bool          // 是否为生产环境
	nonce      string        // 随机字符串标识，不超过32个字符
	signType   string        // 签名算法，目前只支持 sha256
	timestamp  int64         // 当前时间戳，单位为秒
	signature  string        // 最终的签名参数
	prefixPath string        // 请求路径前缀
	timeout    time.Duration // 请求超时设置
}

// 获取实例
func NewHttpClient(appId, appSecret string, prod bool) (ysHttp *YsHttp) {
	return &YsHttp{
		appId:     appId,
		appSecret: appSecret,
		prod:      prod,
	}
}

// SetPathPrefix 设置请求前缀
func (ys *YsHttp) SetPathPrefix(pathPrefix string) *YsHttp {
	ys.prefixPath = pathPrefix
	return ys
}

// SetTimeout 设置超时时间
func (ys *YsHttp) SetTimeout(timeout int64) *YsHttp {
	ys.timeout = time.Duration(timeout) * time.Second
	return ys
}

// Request 请求
func (ys *YsHttp) Request(method string, reqUrl string, reqBody string, headers ...map[string]string) (body string, err error) {
	var reqResMsg string
	var req *http.Request
	for n := 0; n < 1; n++ {
		reqResMsg += fmt.Sprintf("\n%s\n", "<<<<<<<<<<<<<<<<<<<<<<<<request>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
		reqResMsg += fmt.Sprintf("Method:%s\n", method)
		reqResMsg += fmt.Sprintf("reqUrl:%s\n", reqUrl)
		// 签名
		err := ys.generateSignature()
		if err != nil {
			logrus.Error(err)
			reqResMsg += fmt.Sprintf("\nerr:%s\n", err.Error())
			logrus.Debug(reqResMsg)
			break
		}
		// 完整请求地址
		fullUrl := ys.getFullUrl(reqUrl)
		reqResMsg += fmt.Sprintf("fullUrl:%s\n", fullUrl)
		// 解析HOST
		u, err := url.Parse(fullUrl)
		if err != nil {
			logrus.Error(err)
			logrus.Debug(reqResMsg)
			break
		}
		defaultDomain := u.Host
		// 默认超时10S
		if ys.timeout == 0 {
			ys.timeout = time.Duration(10) * time.Second
		}
		client := &http.Client{
			Timeout: ys.timeout,
		}
		req, err = http.NewRequest(method, fullUrl, strings.NewReader(reqBody))
		if err != nil {
			logrus.Error(err)
			reqResMsg += fmt.Sprintf("\nerr:%s\n", err.Error())
			logrus.Debug(reqResMsg)
			break
		}

		req.Header.Set("Accept", "*/*")
		req.Header.Set("Accept-Charset", "utf-8;")
		req.Header.Set("Content-Type", "application/json;")
		req.Header.Set("Host", defaultDomain)
		req.Header.Set("Timeout", fmt.Sprintf("%v", ys.timeout))
		for k, v := range req.Header {
			reqResMsg += fmt.Sprintf("%s:%s\n", k, v[0])
		}
		if len(headers) > 0 {
			for k, v := range headers[0] {
				req.Header.Set(k, v)
				reqResMsg += fmt.Sprintf("%s:%v\n", k, v)
			}
		}

		reqResMsg += fmt.Sprintf("\n%s\n", reqBody)

		var timeDebug = time.Now().UnixNano()
		resp, err := client.Do(req)
		if err != nil {
			logrus.Error(err)
			reqResMsg += fmt.Sprintf("\nerr:%s\n", err.Error())
			logrus.Debug(reqResMsg)
			break
		}
		defer resp.Body.Close()

		reqResMsg += fmt.Sprintf("\n%s\n", "<<<<<<<<<<<<<<<<<<<<<<<<response>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
		for key, _ := range resp.Header {
			reqResMsg += fmt.Sprintf("%s:%s\n", key, resp.Header.Get(key))
		}

		reqResMsg += fmt.Sprintf("status code:%d\n", resp.StatusCode)
		reqResMsg += fmt.Sprintf("execute time:%f\n", float64(time.Now().UnixNano()-timeDebug)/float64(1e9))

		content, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logrus.Error(err)
			reqResMsg += fmt.Sprintf("\nerr:%s\n", err.Error())
			logrus.Debug(reqResMsg)
			break
		}

		body = string(content)
		reqResMsg += fmt.Sprintf("\n%s\n", body)
		if resp.StatusCode != 200 || body == "" {
			logrus.Debug(reqResMsg)
			err = fmt.Errorf(fmt.Sprintf("url：%s，body：%s", fullUrl, body))
			break
		}
	}
	// 打印日志
	logrus.Debug(reqResMsg)
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

// generateSignature 签名
func (ys *YsHttp) generateSignature() (err error) {
	// 签名算法
	ys.signType = "sha256"
	// 请求时间戳（秒级）
	timeLocation, _ := time.LoadLocation("Etc/GMT")
	ys.timestamp = time.Now().In(timeLocation).Unix()
	// 随机字符串
	ys.nonce = RandomStr(32)
	// 待签名字符串
	conStr := fmt.Sprintf("app_id=%s&nonce=%s&sign=%s&timestamp=%d", ys.appId, ys.nonce, ys.signType, ys.timestamp)
	// sha256
	h := hmac.New(sha256.New, []byte(ys.appSecret))
	io.WriteString(h, conStr)
	ys.signature = hex.EncodeToString(h.Sum(nil))
	return
}

// getFullUrl 获取请求url
func (ys *YsHttp) getFullUrl(reqUrl string) (url string) {
	host := TestBaseUrl // 默认走测试环境
	// 开启生成环境
	if ys.prod {
		host = ProdBaseUrl
	}
	prefix := PrefixPath
	if ys.prefixPath != "" {
		prefix = ys.prefixPath
	}
	signatureStr := fmt.Sprintf("app_id=%s&nonce=%s&sign=%s&timestamp=%d&signature=%s", ys.appId, ys.nonce, ys.signType, ys.timestamp, ys.signature)
	// 查询参数拼接签名公共参数
	if strings.Contains(reqUrl, "?") {
		tmp := strings.Split(reqUrl, "?")
		reqUrl = fmt.Sprintf("%s?%s&%s", tmp[0], tmp[1], signatureStr)
	} else {
		reqUrl = fmt.Sprintf("%s?%s", reqUrl, signatureStr)
	}

	return fmt.Sprintf("%s%s%s", host, prefix, reqUrl)
}
