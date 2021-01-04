package models

// ProductCategoriesReq 添加/更新商品类目请求
type ProductCategoriesReq struct {
	DataSourceId string      `json:"dataSourceId"` // 数据源id（create方法返回的data.dataSource.id）
	Categories   *[]Category `json:"categories"`   // 商品类目列表 数组最大长度 50
}

// Category
type Category struct {
	ExternalCategoryId       string `json:"external_category_id"`        // 您为商品类目分配的唯一 id。一般而言，是您在商品库为该商品类目分配的 id。 字段长度最小 1 字节，长度最大 255 字节
	CategoryName             string `json:"category_name"`               // 类目名称 字段长度最小 1 字节，长度最大 255 字节
	CategoryType             int    `json:"category_type"`               // 类目类型，1：前台类目；2：后台类目
	CategoryLevel            int    `json:"category_level"`              // 类目层级
	ExternalParentCategoryId string `json:"external_parent_category_id"` // 父类目 id，顶级类目填 0 字段长度最小 1 字节，长度最大 255 字节
	IsRoot                   bool   `json:"is_root"`                     // 是否为根节点，true：是；false：不是
}
