package api

import (
	"encoding/json"
	"sr-client-go/context"
	"sr-client-go/models"
	"sr-client-go/utils"
)

type ProductCategory struct {
	*context.Context
}

// NewProductCategory 获取商品类目上报实例
func NewProductCategory(ctx *context.Context) *ProductCategory {
	productCategory := new(ProductCategory)
	productCategory.Context = ctx
	return productCategory
}

// Add 添加/更新商品类目
func (pc *ProductCategory) Add(req *models.ProductCategoriesReq) (res utils.ReportCommonRes, err error) {
	response, err := utils.NewHttpClient(pc.AppId, pc.AppSecret, pc.Prod).HttpPostJson("/product_categories/add", req)
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(response), &res)
	if err != nil {
		return
	}
	return
}
