package test

import (
	"fmt"
	sr "github.com/qinyer/sr-client-go"
	"github.com/qinyer/sr-client-go/models"
	"testing"
)

func TestProductCategory(t *testing.T)  {
	req := &models.ProductCategoriesReq{
		DataSourceId: "11263",
		Categories:   &[]models.Category{
			models.Category{
				ExternalCategoryId:       "C02",
				CategoryName:             "服饰",
				CategoryType:             2,
				CategoryLevel:            2,
				ExternalParentCategoryId: "C01",
				IsRoot:                   false,
			},
		},
	}
	res, err := sr.NewYouShuClient(cfg).ProductCategory().Add(req)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(res)
}
