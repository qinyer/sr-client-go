package api

import (
	"encoding/json"
	"github.com/qinyer/sr-client-go/context"
	"github.com/qinyer/sr-client-go/models"
	"github.com/qinyer/sr-client-go/utils"
)

type Store struct {
	*context.Context
}

// NewStore 获取门店仓库上报实例
func NewStore(ctx *context.Context) *Store {
	return &Store{ctx}
}

// 添加/更新门店仓库
func (store *Store) Add(req *models.StoreReq) (res utils.ReportCommonRes, err error) {
	response, err := utils.NewHttpClient(store.Context).HttpPostJson("/store/add", req)
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(response), &res)
	if err != nil {
		return
	}
	return
}
