package config

import "time"

// 有数基础配置
type YouShuConfig struct {
	AppId      string        // appId 数据中心分配的有数应用ID
	AppSecret  string        // appSecret 数据中心分配的有数签名密钥，需要绝对保密
	Prod       bool          // prod 是否为生产环境
	Timeout    time.Duration // 请求超时设置
	PrefixPath string        // 请求前缀
}
