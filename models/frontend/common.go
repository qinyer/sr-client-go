package frontend

type PropsCommon struct {
	Page         string     `json:"page"`           //行为发生的小程序页面路径，开头不要加 /，路径后可带查询参数 例pages/product?sku_id=AOdjf7u
	PageTitle    string     `json:"page_title"`     //行为发生的小程序页面标题 例商品详情、商城首页
	SrSdkVersion string     `json:"sr_sdk_version"` //sdk版本号 例1.1.6
	Time         string     `json:"time"`           //行为发生时间 13位时间戳 例1560409473714
	WxUser       *WxUser    `json:"wx_user"`        //用户相关属性
	Chan         *Chan      `json:"chan"`           //渠道相关属性
	Component    *Component `json:"component,omitempty"`      //组件是视图上的一个区块元素
}

type WxUser struct {
	AppId   string       `json:"app_id"`   //公众号或小程序唯一标识 例wx9d4f5f22pa099f82
	OpenId  string       `json:"open_id"`  //微信用户在小程序下的唯一标识符 例ogN6X0T-ilsH-XmIdzXtuR1f1r3Q
	UserId  string       `json:"user_id,omitempty"`  //用户在商户的唯一标识符 例548019854034
	UnionId string       `json:"union_id,omitempty"` //微信用户在开放平台账号下的唯一标识符 例o6_bmlsdaXds8d6_sgVt7hM3OPfL
	LocalId string       `json:"local_id,omitempty"` //用户在有数的唯一标识符 例360b8853-64bf-3fba-e9a0-5abb1e4d7721
	Tag     *[]WxUserTag `json:"tag,omitempty"`      //用户的个性化标签
	Extra   struct{}     `json:"extra,omitempty"`    //
}

type WxUserTag struct {
	TagId   string `json:"tag_id,omitempty"`   //标签ID，唯一标识符 例游客，若填写了`tag_name`则必填
	TagName string `json:"tag_name,omitempty"` //标签名称 例游客，若填写了`tag_id`则必填
}

type Chan struct {
	ChanWxappScene int         `json:"chan_wxapp_scene"`  //小程序场景值，必填 例1037
	ChanId         string      `json:"chan_id,omitempty"`           //引流渠道的标识符
	ChanReferAppId string      `json:"chan_refer_app_id,omitempty"` //来源小程序或公众号appid 例wx9d6f5f6gea059654
	ChanShopId     string      `json:"chan_shop_id,omitempty"`      //微信用门店ID，若需要计算门店业务则必填
	ChanShopName   string      `json:"chan_shop_name,omitempty"`    //用户在有数的门店名称，展示字段，若chan_shop_id存在则必填
	ChanCustom     *ChanCustom `json:"chan_custom,omitempty"`       //自定义渠道
	Extra          struct{}    `json:"extra,omitempty"`             //
}

type ChanCustom struct {
	ChanCustomId       string `json:"chan_custom_id,omitempty"`         // 自定义渠道的标识符，是自定义渠道的最小粒度
	ChanCustomIdDesc   string `json:"chan_custom_id_desc,omitempty"`    // 自定义渠道的描述
	ChanCustomCat3     string `json:"chan_custom_cat_3,omitempty"`      // 3级自定义渠道的标识符，3级是针对4级的分类，要求4级数据必须存在
	ChanCustomCat3Desc string `json:"chan_custom_cat_3_desc,omitempty"` // 3级自定义渠道的描述，若chan_custom_cat_3存在则必须存在
	ChanCustomCat2     string `json:"chan_custom_cat_2,omitempty"`      // 2级自定义渠道的标识符，2级是针对3级的分类，要求3、4级数据必须存在
	ChanCustomCat2Desc string `json:"chan_custom_cat_2_desc,omitempty"` // 2级自定义渠道的描述，若chan_custom_cat_2存在则必须存在
	ChanCustomCat1     string `json:"chan_custom_cat_1,omitempty"`      // 1级自定义渠道的标识符，1级是针对2级的分类，要求2、3、4级数据必须存在
	ChanCustomCatDesc  string `json:"chan_custom_cat_desc,omitempty"`   // 1级自定义渠道的描述，若chan_custom_cat_1存在则必须存在
}

type Component struct {
	ComponentId   string `json:"component_id,omitempty"`   //组件ID
	ComponentName string `json:"component_name,omitempty"` //组件名称
}
