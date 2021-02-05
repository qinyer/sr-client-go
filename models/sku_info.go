package models

// SkuInfoReq 商详新增接口
type SkuInfoReq struct {
	DataSourceId string     `json:"dataSourceId"` // 数据源id（create方法返回的data.dataSource.id）
	Skus         *[]SkuInfo `json:"skus"`         // 商品SKU列表 数组最大长度 50
}

type SkuInfo struct {
	ExternalSkuId      string              `json:"external_sku_id"`           // 您为商品SKU分配的唯一ID。 一般而言，是您在商品库为该商品SKU分配的id。 字段长度最小 1 字节，长度最大 128 字节
	ExternalSpuId      string              `json:"external_spu_id,omitempty"` // 您为商品SPU分配的唯一ID。 色码款商品必填。 字段长度最小 1 字节，长度最大 128 字节
	SkuBarcode         string              `json:"sku_barcode,omitempty"`     // 商品条形码，超市商品必填。字段长度最小 1 字节，长度最大 64 字节
	BrandName          string              `json:"brand_name"`                // 品牌名称。字段长度最小 1 字节，长度最大 64 字节
	MediaInfo          *MediaInfo          `json:"media_info"`                // 商品图片及视频信息 数组最大长度 10
	CategoryInfo       *[]CategoryInfo     `json:"category_info"`             // 类目信息，数据结构
	SalesInfo          *SkuInfoSalesInfo   `json:"sales_info"`                // 商品销售信息
	TagInfo            *TagInfo            `json:"tag_info"`                  // 商品的标签信息 数组最大长度 10
	DescInfo           *DescInfo           `json:"desc_info"`                 // 商品描述信息
	Price              *SkuInfoPrice       `json:"price"`                     // 商品销售信息
	TargetUrl          *TargetUrl          `json:"target_url"`                // 商品目标页信息
	StoreInfo          *SkuInfoStoreInfo   `json:"store_info"`                // 店铺信息
	CouponInfo         *SkuInfoCouponInfo  `json:"coupon_info"`               // 优惠券信息
	CommentInfo        *CommentInfo        `json:"comment_info"`              // 评论信息
	ThirdPromotionInfo *ThirdPromotionInfo `json:"third_promotion_info"`      // 第三方推广信息
	IsDelete           int                 `json:"is_delete"`                 // 商家标记商品已删除；0：未删除，1：已删除，为空默认为0
}

type MediaInfo struct {
	PrimaryImgs *MediaInfoPrimaryImgs  `json:"primary_imgs"` // 主图图片信息列表 数组最大长度 10
	Imgs        *[]MediaInfoImgs       `json:"imgs"`         // 图片信息列表 数组最大长度 10
	DetailImgs  *[]MediaInfoDetailImgs `json:"detail_imgs"`  // 详情页图片信息列表 数组最大长度 50
	Videos      *[]MediaInfoVideos     `json:"videos"`       // 商品视频信息，数组最大长度 10
}

type CategoryInfo struct {
	CategoryType       int    `json:"category_type"`         // 类目类型，1：前台类目；2：后台类目
	CategoryLevel1Id   string `json:"category_level_1_id"`   // 一级类目id
	CategoryLevel1Name string `json:"category_level_1_name"` // 一级类目名称
	CategoryLevel2Id   string `json:"category_level_2_id"`   // 二级类目id
	CategoryLevel2Name string `json:"category_level_2_name"` // 二级类目名称
	CategoryLevel3Id   string `json:"category_level_3_id"`   // 三级类目id
	CategoryLevel3Name string `json:"category_level_3_name"` // 三级类目名称
}

type SkuInfoSalesInfo struct {
	IsAvailable bool `json:"is_available"` // 上架状态；true：上架；false：下架
	SalesNum    int  `json:"sales_num"`    // 展示销量
}

type TagInfo struct {
	CustomTag *[]CustomTag `json:"custom_tag"` // 商品的标签信息 数组最大长度 50
}

type DescInfo struct {
	ProductNameChinese string `json:"product_name_chinese"` // 商品中文名称 字段长度最小 1 字节，长度最大 100 字节
}

type SkuInfoPrice struct {
	CurrentPrice float64 `json:"current_price"` // 商品现价，单位元
	SkuPrice     float64 `json:"sku_price"`     // 商品原价，单位元
}

type TargetUrl struct {
	UrlMiniprogram      string `json:"url_miniprogram"`      // 微信小程序落地页 url，当落地页为微信小程序时必填 字段长度最小 1 字节，长度最大 100 字节
	MiniprogramAppid    string `json:"miniprogram_appid"`    // 微信小程序 appid，当落地页为微信小程序时必填 字段长度最小 1 字节，长度最大 100 字节
	MiniprogramUsername string `json:"miniprogram_username"` // 小程序原始ID，登录小程序管理后台-设置-基本设置-帐号信息中，gh_xx，当落地页为微信小程序时必填字段长度最小 1 字节，长度最大 100 字节
	UrlMiniprogramQq    string `json:"url_miniprogram_qq"`   // qq小程序落地页 url，当落地页为QQ小程序时必填 字段长度最小 1 字节，长度最大 100 字节
	MiniprogramAppidQq  string `json:"miniprogram_appid_qq"` // qq小程序 appid，当落地页为QQ小程序时必填 字段长度最小 1 字节，长度最大 100 字节
	UrlH5               string `json:"url_h5"`               // h5落地页url 字段长度最小 1 字节，长度最大 100 字节
}

type SkuInfoStoreInfo struct {
	ExternalStoreId string  `json:"external_store_id"` // 店铺id，填写`store_name`时必填，字段长度最小 1 字节，长度最大 100 字节
	StoreName       string  `json:"store_name"`        // 店铺名称，填写`external_store_id`时必填，字段长度最小 1 字节，长度最大 100 字节
	StoreGrade      float64 `json:"store_grade"`       // 按满分5分换算的店铺评分，2位小数
}

type SkuInfoCouponInfo struct {
	CouponId      string  `json:"coupon_id"`       // 券批次id 字段长度最小 1 字节，长度最大 100 字节
	AmountCoupon  float64 `json:"amount_coupon"`   // 券面额，单位元，字段长度最小 1 字节，长度最大 100 字节
	UrlList       string  `json:"url_list"`        // 券地址
	AmountMinimum float64 `json:"amount_minimum"`  // 使用卡券的最低消费金额，单位元
	ShowStartTime string  `json:"show_start_time"` // 展示开始时间/券领取开始时间，unix毫秒级时间戳，为空表示永久
	ShowEndTime   string  `json:"show_end_time"`   // 展示结束时间/券领取结束时间，unix毫秒级时间戳，为空表示永久
	StartTime     string  `json:"start_time"`      // 券有效使用开始时间，unix毫秒级时间戳，为空表示永久
	EndTime       string  `json:"end_time"`        // 券有效使用结束时间，unix毫秒级时间戳，为空表示永久
	IsBest        int     `json:"is_best"`         // 是否最优券，1：是；0：否
}

type CommentInfo struct {
	CommentNum            int     `json:"comment_num"`             // 商品评论数
	PositiveCommentRating float64 `json:"positive_comment_rating"` // 商品好评率，2位小数，不带百分号，如 99.99% 填 99.99
}

type ThirdPromotionInfo struct {
	CommissionRate float64 `json:"commission_rate"` // 佣金比例，2位小数，不带百分号，如 10% 填10.00
	PromoteStatus  int     `json:"promote_status"`  // 推广状态；1：推广中；0：停止推广
}

type MediaInfoPrimaryImgs struct {
	ImgUrl string `json:"img_url"` // 商品主图地址；图片300*300以上，1张，正方形图片可正常打开 字段长度最小 1 字节，长度最大 2048 字节
}

type MediaInfoImgs struct {
	ImgUrl string `json:"img_url"` // 商品图片地址；图片300*300以上，10张，正方形图片可正常打开 字段长度最小 1 字节，长度最大 2048 字节
}

type MediaInfoDetailImgs struct {
	ImgUrl string `json:"img_url"` //商品详情图片地址，50张，图片可正常打开，字段长度最小 1 字节，长度最大 2048 字节
}

type MediaInfoVideos struct {
	VideoUrl string `json:"video_url"` //视频地址；字段长度最小 1 字节，长度最大 2048 字节
	ImgUrl   string `json:"img_url"`   //视频封面图片地址，1张，字段长度最小 1 字节，长度最大 2048 字节
}

type CustomTag struct {
	TagName  string `json:"tag_name"`  //标签名称，长度最大 2048 字节
	TagValue string `json:"tag_value"` //标签值，长度最大 2048 字节
}
