package models

// UserReq 添加会员信息
type UserReq struct {
	DataSourceId string  `json:"dataSourceId"` // 数据源id（create方法返回的data.dataSource.id）
	Users        *[]User `json:"users"`        // 会员列表 数组最大长度 50
}

// User
type User struct {
	UserId      string         `json:"user_id"`      // 会员 id
	PhoneNumber string         `json:"phone_number"` // 手机号码
	UserSpec    *[]UserSpec    `json:"user_spec"`    // 会员应用信息
	CardSpec    *[]CardSpec    `json:"card_spec"`    // 会员卡信息
	DeviceSpec  *[]DeviceSpec  `json:"device_spec"`  // 会员设备信息
	BasicSpec   *BasicSpec     `json:"basic_spec"`   // 会员基础信息
	InvoiceSpec *[]InvoiceSpec `json:"invoice_spec"` // 会员开票信息
}

// UserSpec
type UserSpec struct {
	AppType         int    `json:"app_type"`          // 应用类型，1：微信小程序；2：qq小程序；3：支付宝小程序；4：百度小程序
	Appid           string `json:"appid"`             // 应用的appid
	Openid          string `json:"openid"`            // 对应应用的用户 id
	UserCreatedTime string `json:"user_created_time"` // 会员注册的时间，时间戳，13 位，毫秒
	UserTouchTime   string `json:"user_touch_time"`   // 会员首次触达应用时间，时间戳，13 位，毫秒
}

// CardSpec
type CardSpec struct {
	CardId             string  `json:"card_id"`               // 会员卡号
	CardStatus         int     `json:"card_status"`           // 卡状态，1：启用；2：停用
	CardType           int     `json:"card_type"`             // 主副卡标志，1：主卡；2：副卡
	CardMasterId       string  `json:"card_master_id"`        // 主卡卡号
	CardCreatedTime    string  `json:"card_created_time"`     // 卡创建时间，时间戳，13位，毫秒
	CardLevel          int     `json:"card_level"`            // 卡等级，分为1-9个等级，9为最高级别卡，1为最低级别卡，以此类推
	CardLevelStartTime string  `json:"card_level_start_time"` // 等级开始时间，时间戳，13位，毫秒
	CardLevelEndTime   string  `json:"card_level_end_time"`   // 等级结束时间，时间戳，13位，毫秒
	StoreId            string  `json:"store_id"`              // 所属店仓 id
	Point              float64 `json:"point"`                 // 卡积分，2 位小数
}

// DeviceSpec
type DeviceSpec struct {
	DeviceType int    `json:"device_type"` // 设备类型 1：手机；2：台式机；3：笔记本
	Idfa       string `json:"idfa"`        // idfa
	Uuid       string `json:"uuid"`        // uuid
	Imei       string `json:"imei"`        // imei
	Mac        string `json:"mac"`         // mac地址
}

// BasicSpec
type BasicSpec struct {
	Name         string        `json:"name"`          // 姓名
	Nickname     string        `json:"nickname"`      // 昵称
	HeaderUrl    string        `json:"header_url"`    // 头像 url
	Gender       int           `json:"gender"`        //	性别，1：男：2：女；3：未知
	Age          int           `json:"age"`           // 年龄
	Qq           int           `json:"qq"`            // qq 号码
	IdCardSpec   *[]IdCardSpec `json:"id_card_spec"`  // 身份证件信息
	Birthday     string        `json:"birthday"`      // 生日，时间戳，13 位，毫秒
	Province     string        `json:"province"`      // 居住地省
	ProvinceCode int           `json:"province_code"` // 省份编码，使用《统计用区划代码和城乡分代码编制规则》
	City         string        `json:"city"`          // 居住地市
	CityCode     int           `json:"city_code"`     // 城市编码，使用《统计用区划代码和城乡分代码编制规则》
	District     string        `json:"district"`      // 居住行政区
	DistrictCode int           `json:"district_code"` // 行政区编码，使用《统计用区划代码和城乡分代码编制规则》
	Address      string        `json:"address"`       // 居住地址
	Email        []string      `json:"email"`         // 电子邮件，格式为 xx@xx
}

// InvoiceSpec
type InvoiceSpec struct {
	InvoiceType     int    `json:"invoice_type"`     // 发票类型，1：普通发票；2：增值税专用发票
	InvoiceTitle    string `json:"invoice_title"`    // 发票抬头
	InvoiceCompany  string `json:"invoice_company"`  // 公司名称
	InvoiceTaxpayer string `json:"invoice_taxpayer"` // 纳税人识别号
}

// IdCardSpec
type IdCardSpec struct {
	IdCardType int    `json:"id_card_type"` // 1：身份证；2：护照；3：港澳通行证；4：台湾通行证；5：台胞证；6：回乡证：7：军官证；8：警官证；9：士兵证；
	IdCardNo   string `json:"id_card_no"`   // 证件号码
}
