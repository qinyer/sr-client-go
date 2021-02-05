package models

// StoreReq 添加/更新门店仓库请求
type StoreReq struct {
	DataSourceId string   `json:"dataSourceId"` // 数据源id（create方法返回的data.dataSource.id）
	Stores       *[]Store `json:"stores"`       // 仓库列表 数组最大长度 50
}

// Store
type Store struct {
	ExternalStoreId       string        `json:"external_store_id"`                 // 您为仓库分配的唯一ID。 一般而言，是您在商品库为该仓库分配的id。 字段长度最小 1 字节，长度最大 32 字节
	Type                  int           `json:"type"`                              // 门店类型；1：门店仓（可进行线下购物的门店）；2：前置仓（支持同城配/落地配的仓库；3：电商仓（支持快递配送的仓库）
	BusinessType          int           `json:"business_type"`                     // 门店经营类型 可固定填写 99：其他
	OperationStatus       int           `json:"operation_status"`                  // 运营状态 1：开业；2：关店；3：暂停营业 与daojia_operation_status两者必填其一
	DaojiaOperationStatus int           `json:"daojia_operation_status,omitempty"` // 暂缺
	PhoneNumbers          []string      `json:"phone_numbers,omitempty"`           // 店仓电话列表，格式如：075588888888
	LocationInfo          *LocationInfo `json:"location_info"`                     // 地址信息
	BasicProps            *BasicProps   `json:"basic_props"`                       // 基础信息
	DeliveryInfo          *DeliveryInfo `json:"delivery_info,omitempty"`           // 配送信息
}

// LocationInfo
type LocationInfo struct {
	CountryCode  string     `json:"country_code,omitempty"`  // 国家编码，CN 字段长度最小 1 字节，长度最大 64 字节
	CountryName  string     `json:"country_name"`            // 国家名称，中国 字段长度最小 1 字节，长度最大 64 字节
	ProvinceCode int        `json:"province_code,omitempty"` // 省份编码，使用《统计用区划代码和城乡分代码编制规则》，440000
	ProvinceName string     `json:"province_name"`           // 省份名称，XX省 字段长度最小 1 字节，长度最大 64 字节
	CityCode     int        `json:"city_code,omitempty"`     // 城市编码，使用《统计用区划代码和城乡分代码编制规则》
	CityName     string     `json:"city_name"`               // 城市名称，XX市 字段长度最小 1 字节，长度最大 64 字节
	DistrictCode int        `json:"district_code,omitempty"` // 行政区编码，使用《统计用区划代码和城乡分代码编制规则》
	DistrictName string     `json:"district_name"`           // 行政区名称XX区 字段长度最小 1 字节，长度最大 64 字节
	Address      string     `json:"address"`                 // 店仓详细地址
	GeoInfo      *[]GeoInfo `json:"geo_info,omitempty"`      // 坐标信息
}

// BasicProps
type BasicProps struct {
	Name          string           `json:"name"`                     // 店仓名称建议格式：品牌名称（XX店） 字段长度最小 1 字节，长度最大 32 字节
	OpeningTime   string           `json:"opening_time,omitempty"`   // 店仓开业时间 时间戳 毫秒 字段长度为长度为 13 字节
	OperatingTime *[]OperatingTime `json:"operating_time,omitempty"` // 为空默认周一至周日，10:00-22:00
}

// DeliveryInfo
type DeliveryInfo struct {
	RangeType int         `json:"range_type,omitempty"` // 配送范围类型 1：圆形；3：多边形；store支持到家配送时必填
	Radius    string      `json:"radius,omitempty"`     // 配送半径 range_type=1时，必填 字段长度最小 1 字节，长度最大 128 字节
	GeoGroup  *[]GeoGroup `json:"geo_group,omitempty"`  // 坐标信息组合，坐标点的二维数组，表示多个形状 range_type=3时，必填
}

// GeoInfo
type GeoInfo struct {
	Latitude  float64 `json:"latitude,omitempty"`  // 店仓纬度，如：22.5200
	Longitude float64 `json:"longitude,omitempty"` // 店仓经度，如：113.9400
	Type      int     `json:"type,omitempty"`      // 填写了longitude和latitude则必填 1：腾讯/高德地图；2：百度地图：3：国际标准
}

// OperatingTime
type OperatingTime struct {
	DateZone string `json:"date_zone,omitempty"` // 日期区间，例：周一至周日；字段长度最小 1 字节，长度最大 64 字节
	TimeZone string `json:"time_zone,omitempty"` // 时间区间，例：10:00-22:00；字段长度最小 1 字节，长度最大 64 字节
}

// GeoGroup
type GeoGroup struct {
	Geos *[]GeoInfo `json:"geos,omitempty"` // 坐标点信息
}
