package api

import (
	"encoding/json"
	"github.com/qinyer/sr-client-go/context"
	"github.com/qinyer/sr-client-go/models"
	"github.com/qinyer/sr-client-go/utils"
)

type ProductCategory struct {
	*context.Context
}

// NewProductCategory 获取商品类目上报实例
func NewProductCategory(ctx *context.Context) *ProductCategory {
	return &ProductCategory{ctx}
}

// Add 添加/更新商品类目
func (pc *ProductCategory) Add(req *models.ProductCategoriesReq) (res utils.ReportCommonRes, err error) {
	response, err := utils.NewHttpClient(pc.Context).HttpPostJson("/product_categories/add", req)
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(response), &res)
	if err != nil {
		return
	}
	return
}
