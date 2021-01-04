package models

// SpuReq 添加/更新商品spu请求
type SpuReq struct {
	DataSourceId string `json:"dataSourceId"` // 数据源id（create方法返回的data.dataSource.id）
	Spus         *[]Spu `json:"spus"`         // 商品SPU列表 数组最大长度 50
}

// Spu
type Spu struct {
	ExternalSpuId string     `json:"external_spu_id"` // 您为商品SPU分配的唯一ID。 一般而言，是您在商品库为该商品SpU分配的id。 字段长度最小 1 字节，长度最大 128 字节
	DescProps     *DescProps `json:"desc_props"`      // 商品描述类型
}
