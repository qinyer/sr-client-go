package models

// SkuReq 添加/更新商品SKU请求
type SkuReq struct {
	DataSourceId string `json:"dataSourceId"` // 数据源id（create方法返回的data.dataSource.id）
	Skus         *[]Sku `json:"skus"`         // 商品SKU列表 数组最大长度 50
}

// Sku
type Sku struct {
	ExternalSkuId       string          `json:"external_sku_id"`           // 您为商品SKU分配的唯一ID。 一般而言，是您在商品库为该商品SKU分配的id。 字段长度最小 1 字节，长度最大 128 字节
	ExternalSpuId       string          `json:"external_spu_id,omitempty"` // 您为商品SPU分配的唯一ID。 色码款商品必填。 字段长度最小 1 字节，长度最大 128 字节
	SkuBarcode          string          `json:"sku_barcode,omitempty"`     // 商品条形码，超市商品必填。字段长度最小 1 字节，长度最大 64 字节
	ImgUrls             *[]ImgUrl       `json:"img_urls"`                  // 商品图片url列表 数组最大长度 10
	ProductProps        *ProductProps   `json:"product_props,omitempty"`   // 商品属性
	CategoryProps       *[]CategoryProp `json:"category_props,omitempty"`  // 类目属性
	SalesProps          *SalesProps     `json:"sales_props"`               // 商品销售属性
	DescProps           *DescProps      `json:"desc_props"`                // 商品描述属性
	ExternalCreatedTime string          `json:"external_created_time"`     // 客户侧的商品创建时间 字段长度为 13 字节
}

// ImgUrl
type ImgUrl struct {
	PrimaryImgs *[]Img `json:"primary_imgs"`          // 主图图片信息列表 数组最大长度 10
	Imgs        *[]Img `json:"imgs"`                  // 图片信息列表 数组最大长度 10
	DetailImgs  *[]Img `json:"detail_imgs,omitempty"` // 详情页图片信息列表 数组最大长度 10
}

//CategoryProp
type CategoryProp struct {
	ExternalCategoryIdLeaf string `json:"external_category_id_leaf,omitempty"` // 您为商品分配的叶子类目 id。字段长度最小 1 字节，长度最大 255 字节
	CategoryType           int    `json:"category_type,omitempty"`             // 固定传 2，表示后台类目
}

// Img
type Img struct {
	ImgUrl string `json:"img_url"` // 商品主图地址；图片300*300以上，1张，正方形图片可正常打开 字段长度最小 1 字节，长度最大 2048 字节
}

// ProductProps
type ProductProps struct {
	Color *Color `json:"color,omitempty"` // 商品颜色属性，色码款商品必传
	IsO2o bool   `json:"is_o2o"`          // true：线上线下同款；false：线上线下不同款
	Size  string `json:"size,omitempty"`  // 商品尺码，色码款商品必传 字段长度最小 1 字节，长度最大 8 字节
}

// Color
type Color struct {
	ColorRgb  string `json:"color_rgb,omitempty"`  // 颜色RGB编码 字段长度最小 1 字节，长度最大 64 字节
	ColorName string `json:"color_name,omitempty"` // 颜色名称，如黄色 字段长度最小 1 字节，长度最大 64 字节
}

// SalesProps
type SalesProps struct {
	IsAvailable bool `json:"is_available"` // 上架状态；true：上架；false：下架
}

// DescProps
type DescProps struct {
	ProductNameChinese string `json:"product_name_chinese"` // 商品中文名 字段长度最小 1 字节，长度最大 100 字节
}
